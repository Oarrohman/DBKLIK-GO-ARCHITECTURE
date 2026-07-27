// Package routes is the root-level aggregator: it lists every module
// that should be mounted onto the API, decoupling internal/bootstrap
// from knowing which modules exist.
package routes

import (
	"github.com/dbklik/dbklik-go-architecture/internal/container"
	internalhttp "github.com/dbklik/dbklik-go-architecture/internal/http"
	inventoryprovider "github.com/dbklik/dbklik-go-architecture/modules/inventory/infrastructure/provider"
)

// Modules returns every module's route registrar. Add new modules here
// as they're built — this is the single place that lists them.
func Modules(c *container.Container) []internalhttp.ModuleRouter {
	return []internalhttp.ModuleRouter{
		inventoryprovider.New(c),
	}
}
