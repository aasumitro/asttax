package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aasumitro/asttax/internal/common"
	"github.com/spf13/viper"
)

type (
	Option func(cfg *Config)

	Config struct {
		sync.RWMutex
		ctx context.Context
		// SERVER CONFIGURATION
		ServerName    string `mapstructure:"SERVER_NAME"`
		ServerVersion string `mapstructure:"SERVER_VERSION"`
		// TELEGRAM CONFIGURATION
		TelegramBotToken string `mapstructure:"TELEGRAM_BOT_TOKEN"`
		// DATASTORE URL
		DatastoreDriver string `mapstructure:"DATASTORE_DRIVER"`
		SQLiteDsnURL    string `mapstructure:"SQLITE_DSN_URL"`
		// APP DEPS
		SQLPool *sql.DB
	}
)

var (
	configOnce sync.Once
	instance   *Config
)

func LoadWith(
	ctx context.Context,
	options ...Option,
) *Config {
	configOnce.Do(func() {
		// error handling for a specific case
		if err := viper.ReadInConfig(); err != nil {
			var configFileNotFoundError viper.ConfigFileNotFoundError
			if errors.As(err, &configFileNotFoundError) {
				// Config file not found; ignore error if desired
				log.Fatal(common.ErrEnvMsg)
			}
			log.Fatalf("ENV_ERROR: %s", err.Error())
		}
		// set context & extract config to struct
		instance = &Config{ctx: ctx}
		if err := viper.Unmarshal(instance); err != nil {
			log.Fatalf("ENV_ERROR: %s", err.Error())
		}
		// set options
		for _, opt := range options {
			opt(instance)
		}
	})
	return instance
}

func SQLiteDBConnection() Option {
	return func(cfg *Config) {
		db, err := sql.Open(cfg.DatastoreDriver, cfg.SQLiteDsnURL)
		if err != nil {
			log.Fatalf("SQLITE_OPEN_ERROR: %v", err)
		}
		// Set SQLite PRAGMA configurations for optimized RW
		pragmaStatements := []string{
			"PRAGMA synchronous = OFF;",
			"PRAGMA journal_mode = WAL;",
			"PRAGMA temp_store = MEMORY;",
		}
		for _, pragma := range pragmaStatements {
			if _, err := db.Exec(pragma); err != nil {
				log.Fatalf("SQLITE_PRAGMA_ERROR: %v", err)
			}
		}
		// Configure connection pooling
		const dbMaxOpenConnection, dbMaxIdleConnection = 100, 10
		db.SetMaxIdleConns(dbMaxIdleConnection)
		db.SetMaxOpenConns(dbMaxOpenConnection)
		db.SetConnMaxLifetime(time.Hour)
		// Validate the database connection
		if err := db.Ping(); err != nil {
			log.Fatalf(fmt.Sprintf("SQLITE_PING_ERROR: %s", err.Error()))
		}
		// Assign the configured DB connection to the config
		cfg.SQLPool = db
		//
		if err := cfg.initSQLiteDB(); err != nil {
			log.Fatalf("SQLITE_INIT_ERROR: %s", err.Error())
		}
	}
}

func (c *Config) initSQLiteDB() error {
	c.Lock()
	defer c.Unlock()
	//goland:noinspection ALL
	_, err := c.SQLPool.Exec(`CREATE TABLE IF NOT EXISTS users (
      	id INTEGER PRIMARY KEY,
      	telegram_id TEXT UNIQUE NOT NULL,
        wallet_address TEXT UNIQUE NOT NULL,
        private_key TEXT NOT NULL,
        trade_fees FLOAT4 NOT NULL DEFAULT 0.0015, -- Fast: 0.0015 SOL, Turbo: 0.0075 SOL, Custom: by user
        accept_aggrement BOOLEAN DEFAULT FALSE,  -- accept aggrement
        confirm_trade_protection BOOLEAN DEFAULT FALSE,  -- Confirm Before Continue
        mev_buy_protection BOOLEAN DEFAULT FALSE,  -- Maximal Extractable Value protection
        mev_sell_protection BOOLEAN DEFAULT FALSE,  -- Maximal Extractable Value protection
        buy_slippage INTEGER NOT NULL DEFAULT 15, -- default 15%
        sell_slippage INTEGER NOT NULL  DEFAULT 15);  -- default 15%
	CREATE INDEX IF NOT EXISTS idx_uid ON users (telegram_id);`)
	return err
}
