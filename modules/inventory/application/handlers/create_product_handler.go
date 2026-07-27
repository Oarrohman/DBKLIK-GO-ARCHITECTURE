// Package handlers implements the CQRS handlers: one Handle method per
// command/query, orchestrating domain entities and repositories.
package handlers

import (
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/application/commands"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/entities"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/repositories"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/valueobjects"
	apperrors "github.com/dbklik/dbklik-go-architecture/pkg/errors"
	"github.com/dbklik/dbklik-go-architecture/pkg/utils"
)

type CreateProductHandler struct {
	repo repositories.ProductRepository
}

func NewCreateProductHandler(repo repositories.ProductRepository) *CreateProductHandler {
	return &CreateProductHandler{repo: repo}
}

func (h *CreateProductHandler) Handle(cmd commands.CreateProductCommand) (*entities.Product, error) {
	sku, err := valueobjects.NewSKU(cmd.Input.SKU)
	if err != nil {
		return nil, err
	}

	if existing, _ := h.repo.FindBySKU(sku.String()); existing != nil {
		return nil, apperrors.Conflict("a product with this SKU already exists")
	}

	price, err := valueobjects.NewMoney(cmd.Input.PriceCents)
	if err != nil {
		return nil, err
	}

	product, err := entities.NewProduct(utils.NewID(), cmd.Input.Name, sku, price, cmd.Input.Quantity)
	if err != nil {
		return nil, err
	}

	if err := h.repo.Save(product); err != nil {
		return nil, apperrors.Internal(err)
	}

	return product, nil
}
