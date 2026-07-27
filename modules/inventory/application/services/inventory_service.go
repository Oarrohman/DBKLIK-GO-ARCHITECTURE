// Package services exposes the module's use cases as a single facade
// consumed by the presentation layer, so controllers depend on one
// cohesive API instead of wiring commands/queries/actions themselves.
package services

import (
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/application/actions"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/application/commands"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/application/dto"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/application/handlers"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/application/mapper"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/application/queries"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/repositories"
)

type InventoryService struct {
	createHandler *handlers.CreateProductHandler
	getHandler    *handlers.GetProductHandler
	listHandler   *handlers.ListProductsHandler
	adjustStock   *actions.AdjustStock
	archive       *actions.ArchiveProduct
}

func NewInventoryService(repo repositories.ProductRepository) *InventoryService {
	return &InventoryService{
		createHandler: handlers.NewCreateProductHandler(repo),
		getHandler:    handlers.NewGetProductHandler(repo),
		listHandler:   handlers.NewListProductsHandler(repo),
		adjustStock:   actions.NewAdjustStock(repo),
		archive:       actions.NewArchiveProduct(repo),
	}
}

func (s *InventoryService) CreateProduct(input dto.CreateProductInput) (dto.ProductOutput, error) {
	product, err := s.createHandler.Handle(commands.CreateProductCommand{Input: input})
	if err != nil {
		return dto.ProductOutput{}, err
	}
	return mapper.ToOutput(product), nil
}

func (s *InventoryService) GetProduct(id string) (dto.ProductOutput, error) {
	product, err := s.getHandler.Handle(queries.GetProductQuery{ProductID: id})
	if err != nil {
		return dto.ProductOutput{}, err
	}
	return mapper.ToOutput(product), nil
}

func (s *InventoryService) ListProducts(page, pageSize int) (dto.ListOutput, error) {
	products, total, err := s.listHandler.Handle(queries.ListProductsQuery{Page: page, PageSize: pageSize})
	if err != nil {
		return dto.ListOutput{}, err
	}
	return dto.ListOutput{Items: mapper.ToOutputList(products), Total: total, Page: page}, nil
}

func (s *InventoryService) AdjustStock(id string, delta int) error {
	return s.adjustStock.Execute(id, delta)
}

func (s *InventoryService) ArchiveProduct(id string) error {
	return s.archive.Execute(id)
}
