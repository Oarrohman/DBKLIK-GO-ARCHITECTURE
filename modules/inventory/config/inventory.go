// Package config holds inventory-module-specific configuration,
// separate from the app-wide internal/config.
package config

import "github.com/dbklik/dbklik-go-architecture/pkg/config"

type Config struct {
	DefaultPageSize int
}

func Load() *Config {
	return &Config{
		DefaultPageSize: config.GetInt("INVENTORY_DEFAULT_PAGE_SIZE", 20),
	}
}
