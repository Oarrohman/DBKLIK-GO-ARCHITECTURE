// Package services holds pure domain services: business logic that
// doesn't naturally belong to a single entity.
package services

import "github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/entities"

// StockCalculator computes stock-related figures across a set of products.
type StockCalculator struct{}

func NewStockCalculator() *StockCalculator {
	return &StockCalculator{}
}

func (StockCalculator) TotalQuantity(products []*entities.Product) int {
	total := 0
	for _, p := range products {
		total += p.Quantity
	}
	return total
}

func (StockCalculator) TotalValueCents(products []*entities.Product) int64 {
	var total int64
	for _, p := range products {
		total += p.Price.Multiply(p.Quantity).Cents()
	}
	return total
}
