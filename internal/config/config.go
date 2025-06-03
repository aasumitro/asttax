package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/aasumitro/asttax/internal/common"
	"github.com/aasumitro/asttax/internal/util/cache"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/spf13/viper"
)

const (
	envStaging    = "staging"
	envProduction = "production"
)

type (
	Option func(cfg *Config)

	Config struct {
		sync.RWMutex
		ctx context.Context
		// SERVER CONFIGURATION
		ServerName        string `mapstructure:"SERVER_NAME"`
		ServerVersion     string `mapstructure:"SERVER_VERSION"`
		ServerEnvironment string `mapstructure:"SERVER_ENVIRONMENT"`
		// DATASTORE URL
		DatastoreDriver string `mapstructure:"DATASTORE_DRIVER"`
		SQLiteDsnURL    string `mapstructure:"SQLITE_DSN_URL"`
		// TELEGRAM CONFIGURATION
		TelegramBotToken string `mapstructure:"TELEGRAM_BOT_TOKEN"`
		// ENCRYPTION SECRET KEY
		SecretKey string `mapstructure:"SECRET_KEY"`
		// API URL
		CoingeckoAPIURL string `mapstructure:"COINGECKO_API_URL"`
		// APP DEPS
		SQLPool   *sql.DB
		CachePool *cache.Cache
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

func (c *Config) GetServerIdentity() string {
	return fmt.Sprintf("%s v%s",
		c.ServerName, c.ServerVersion)
}

func (c *Config) GetRPCEndpoint() string {
	rpcEndpoint := rpc.DevnetRPCEndpoint
	if c.ServerEnvironment == envStaging {
		rpcEndpoint = rpc.TestnetRPCEndpoint
	}
	if c.ServerEnvironment == envProduction {
		rpcEndpoint = rpc.MainnetRPCEndpoint
	}
	return rpcEndpoint
}
