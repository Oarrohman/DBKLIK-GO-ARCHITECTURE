// Package actions holds single-purpose use cases that don't fit the
// CQRS command/query split cleanly (typically a single domain operation
// wrapped with its persistence step).
package actions

import (
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/repositories"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/specifications"
	apperrors "github.com/dbklik/dbklik-go-architecture/pkg/errors"
)

type AdjustStock struct {
	repo repositories.ProductRepository
}

func NewAdjustStock(repo repositories.ProductRepository) *AdjustStock {
	return &AdjustStock{repo: repo}
}

func (a *AdjustStock) Execute(productID string, delta int) error {
	product, err := a.repo.FindByID(productID)
	if err != nil {
		return apperrors.Internal(err)
	}
	if product == nil {
		return apperrors.NotFound("product not found")
	}
	if !specifications.CanReceiveGoods(product) && delta > 0 {
		return apperrors.Conflict("archived products cannot receive goods")
	}

	if err := product.AdjustStock(delta); err != nil {
		return err
	}
	if err := a.repo.Save(product); err != nil {
		return apperrors.Internal(err)
	}
	return nil
}

type ArchiveProduct struct {
	repo repositories.ProductRepository
}

func NewArchiveProduct(repo repositories.ProductRepository) *ArchiveProduct {
	return &ArchiveProduct{repo: repo}
}

func (a *ArchiveProduct) Execute(productID string) error {
	product, err := a.repo.FindByID(productID)
	if err != nil {
		return apperrors.Internal(err)
	}
	if product == nil {
		return apperrors.NotFound("product not found")
	}
	if !specifications.CanArchiveProduct(product) {
		return apperrors.Conflict("cannot archive a product that still has stock")
	}

	if err := product.Archive(); err != nil {
		return err
	}
	if err := a.repo.Save(product); err != nil {
		return apperrors.Internal(err)
	}
	return nil
}
