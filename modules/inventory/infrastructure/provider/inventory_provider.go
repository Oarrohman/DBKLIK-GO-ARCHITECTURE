// Package provider wires the inventory module's dependencies together
// (repository -> service -> controller) and exposes RegisterRoutes so
// internal/router can attach the module without knowing its internals.
package provider

import (
	"github.com/dbklik/dbklik-go-architecture/internal/container"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/application/services"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/infrastructure/repository"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/presentation/http/controllers"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/routes"
	"github.com/gin-gonic/gin"
)

type InventoryProvider struct {
	controller *controllers.ProductController
}

// New builds the module's dependency graph. When the shared container
// exposes a real database connection, swap ProductRepositoryMemory for a
// persistence-backed implementation here — nothing else in the module changes.
func New(c *container.Container) *InventoryProvider {
	repo := repository.NewProductRepositoryMemory()
	service := services.NewInventoryService(repo)
	controller := controllers.NewProductController(service)

	return &InventoryProvider{controller: controller}
}

func (p *InventoryProvider) RegisterRoutes(api *gin.RouterGroup) {
	routes.RegisterAPI(api, p.controller)
}
