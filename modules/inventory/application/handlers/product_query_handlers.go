package handlers

import (
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/application/queries"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/entities"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/repositories"
	apperrors "github.com/dbklik/dbklik-go-architecture/pkg/errors"
	"github.com/dbklik/dbklik-go-architecture/pkg/utils"
)

type GetProductHandler struct {
	repo repositories.ProductRepository
}

func NewGetProductHandler(repo repositories.ProductRepository) *GetProductHandler {
	return &GetProductHandler{repo: repo}
}

func (h *GetProductHandler) Handle(q queries.GetProductQuery) (*entities.Product, error) {
	product, err := h.repo.FindByID(q.ProductID)
	if err != nil {
		return nil, apperrors.Internal(err)
	}
	if product == nil {
		return nil, apperrors.NotFound("product not found")
	}
	return product, nil
}

type ListProductsHandler struct {
	repo repositories.ProductRepository
}

func NewListProductsHandler(repo repositories.ProductRepository) *ListProductsHandler {
	return &ListProductsHandler{repo: repo}
}

func (h *ListProductsHandler) Handle(q queries.ListProductsQuery) ([]*entities.Product, int, error) {
	offset, limit := utils.Paginate(q.Page, q.PageSize)
	products, total, err := h.repo.List(offset, limit)
	if err != nil {
		return nil, 0, apperrors.Internal(err)
	}
	return products, total, nil
}
