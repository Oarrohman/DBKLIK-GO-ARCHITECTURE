// Package container holds the framework-level dependencies that are
// shared across every module (config, logger, database), so modules
// receive them instead of constructing their own.
package container

import (
	"log/slog"

	"github.com/dbklik/dbklik-go-architecture/internal/config"
	"github.com/dbklik/dbklik-go-architecture/internal/database"
)

type Container struct {
	Config *config.Config
	Logger *slog.Logger
	DB     *database.DB
}

func New(cfg *config.Config, logger *slog.Logger, db *database.DB) *Container {
	return &Container{Config: cfg, Logger: logger, DB: db}
}
