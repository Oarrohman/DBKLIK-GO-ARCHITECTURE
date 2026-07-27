// Package bootstrap composes config, logger, database, container and
// router into a ready-to-serve gin.Engine. cmd/* entrypoints call this
// instead of wiring the framework themselves.
package bootstrap

import (
	"github.com/dbklik/dbklik-go-architecture/internal/config"
	"github.com/dbklik/dbklik-go-architecture/internal/container"
	"github.com/dbklik/dbklik-go-architecture/internal/database"
	"github.com/dbklik/dbklik-go-architecture/internal/logger"
	"github.com/dbklik/dbklik-go-architecture/internal/router"
	"github.com/dbklik/dbklik-go-architecture/routes"
	"github.com/gin-gonic/gin"
)

func Boot() (*gin.Engine, *container.Container) {
	cfg := config.Load()
	log := logger.New(cfg)
	db := database.Connect(cfg)
	c := container.New(cfg, log, db)

	engine := router.New(c, routes.Modules(c)...)

	return engine, c
}
