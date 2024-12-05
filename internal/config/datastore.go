package config

import (
	"database/sql"
	"log"
	"time"

	"github.com/aasumitro/asttax/internal/util/cache"
)

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
			log.Fatalf("SQLITE_PING_ERROR: %s\n", err.Error())
		}
		// Assign the configured DB connection to the config
		cfg.SQLPool = db
		// Start init database table and others item
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
      	telegram_id BIGINT UNIQUE NOT NULL,
        bot_language TEXT NOT NULL CHECK (bot_language IN ('en', 'id'))  DEFAULT 'en',					-- en: english, id: indonesia
        accept_agreement BOOLEAN DEFAULT FALSE,  														-- accept aggrement
        wallet_address TEXT UNIQUE NOT NULL,
        private_key TEXT NOT NULL,
        trade_fees TEXT NOT NULL CHECK (trade_fees IN ('fast', 'turbo'))  DEFAULT 'fast', 				-- Fast: 0.0015 SOL, Turbo: 0.0075 SOL
        confirm_trade_protection BOOLEAN DEFAULT FALSE,													-- Confirm Before Continue
        buy_amount_p1 FLOAT4 NOT NULL DEFAULT 0.25, 													-- default 0.25 SOL
        buy_amount_p2 FLOAT4 NOT NULL DEFAULT 0.5, 														-- default 0.5 SOL
        buy_amount_p3 FLOAT4 NOT NULL DEFAULT 1, 														-- default 1 SOL
        buy_amount_p4 FLOAT4 NOT NULL DEFAULT 2.5, 														-- default 2.5 SOL
        buy_amount_p5 FLOAT4 NOT NULL DEFAULT 5, 														-- default 5 SOL
        buy_amount_p6 FLOAT4 NOT NULL DEFAULT 10, 														-- default 10 SOL
        buy_slippage FLOAT4 NOT NULL DEFAULT 15, 														-- default 15%
        sell_amount_p1 FLOAT4 NOT NULL DEFAULT 25, 														-- default 25%
        sell_amount_p2 FLOAT4 NOT NULL DEFAULT 50, 														-- default 50%
        sell_amount_p3 FLOAT4 NOT NULL DEFAULT 100, 													-- default 100%
        sell_slippage FLOAT4 NOT NULL  DEFAULT 15, 														-- default 15%
        sell_protection BOOLEAN DEFAULT FALSE,
    	created_at BIGINT, updated_at BIGINT); 
	CREATE INDEX IF NOT EXISTS idx_uid ON users (telegram_id);`)
	return err
}

func InMemoryCache() Option {
	return func(cfg *Config) {
		workerDuration := 5
		defaultWorkerInterval := time.Duration(workerDuration) * time.Minute
		cfg.CachePool = cache.New(defaultWorkerInterval, defaultWorkerInterval)
	}
}
