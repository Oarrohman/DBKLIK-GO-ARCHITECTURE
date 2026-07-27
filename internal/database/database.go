// Package database owns the shared database connection. Modules never
// open their own connections — they receive *gorm.DB through the
// container and build their infrastructure/persistence repositories on it.
package database

import (
	"log/slog"

	"github.com/dbklik/dbklik-go-architecture/internal/config"
)

// DB is a placeholder for the shared connection handle. Swap this for
// *gorm.DB (or *sqlx.DB) once a driver is added to go.mod; every module's
// infrastructure/persistence layer depends on this type, not on gorm directly.
type DB struct {
	Driver string
	DSN    string
}

// Connect opens the shared connection described by cfg. When no driver is
// configured, it returns nil so the app can still run modules that only
// need an in-memory repository (as the sample inventory module does).
func Connect(cfg *config.Config) *DB {
	if cfg.DBDriver == "" {
		slog.Info("database: no driver configured, skipping connection")
		return nil
	}

	slog.Info("database: connecting", "driver", cfg.DBDriver)
	return &DB{Driver: cfg.DBDriver, DSN: cfg.DBDSN}
}
