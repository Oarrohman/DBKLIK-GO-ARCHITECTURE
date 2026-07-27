// Package commands holds write-intent inputs for CQRS handlers.
package commands

import "github.com/dbklik/dbklik-go-architecture/modules/inventory/application/dto"

type CreateProductCommand struct {
	Input dto.CreateProductInput
}

type AdjustStockCommand struct {
	ProductID string
	Delta     int
}

type ArchiveProductCommand struct {
	ProductID string
}
