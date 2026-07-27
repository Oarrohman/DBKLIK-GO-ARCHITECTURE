// Package requests holds HTTP-bound structs validated at the controller
// boundary before being mapped into application-layer DTOs.
package requests

type CreateProductRequest struct {
	Name       string `json:"name" validate:"required"`
	SKU        string `json:"sku" validate:"required"`
	PriceCents int64  `json:"price_cents" validate:"gte=0"`
	Quantity   int    `json:"quantity" validate:"gte=0"`
}

type AdjustStockRequest struct {
	Delta int `json:"delta" validate:"required"`
}
