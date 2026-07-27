// Package contracts declares the module's extension points: interfaces
// that infrastructure implementations satisfy so behavior (e.g. stock
// valuation) can be swapped without touching the domain or application layer.
package contracts

import "github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/entities"

// StockValuationMethod computes the monetary value (in cents) of a
// product's current stock, e.g. FIFO, LIFO, or average cost.
type StockValuationMethod interface {
	Valuate(product *entities.Product) int64
}

// ReorderPolicy decides whether a product needs restocking.
type ReorderPolicy interface {
	ShouldReorder(product *entities.Product) bool
}
