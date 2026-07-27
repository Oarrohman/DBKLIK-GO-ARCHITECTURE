// Package config defines the application-wide configuration struct,
// loaded once at bootstrap and shared through the container.
package config

import "github.com/dbklik/dbklik-go-architecture/pkg/config"

type Config struct {
	AppEnv    string
	AppPort   string
	LogJSON   bool
	LogDebug  bool
	DBDriver  string
	DBDSN     string
	StoragePath string
}

// Load reads configuration from the environment (optionally seeded from a
// .env file at the project root) into a Config struct.
func Load() *Config {
	config.LoadDotEnv(".env")

	return &Config{
		AppEnv:      config.GetString("APP_ENV", "local"),
		AppPort:     config.GetString("APP_PORT", "8080"),
		LogJSON:     config.GetBool("LOG_JSON", false),
		LogDebug:    config.GetBool("LOG_DEBUG", true),
		DBDriver:    config.GetString("DB_DRIVER", ""),
		DBDSN:       config.GetString("DB_DSN", ""),
		StoragePath: config.GetString("STORAGE_PATH", "storage"),
	}
}
