// Package dto holds the application layer's input/output data shapes,
// decoupling the domain entity's internal representation from what
// callers (HTTP requests, CLI commands, ...) send and receive.
package dto

type CreateProductInput struct {
	Name        string `json:"name" validate:"required"`
	SKU         string `json:"sku" validate:"required"`
	PriceCents  int64  `json:"price_cents" validate:"gte=0"`
	Quantity    int    `json:"quantity" validate:"gte=0"`
}

type AdjustStockInput struct {
	Delta int `json:"delta" validate:"required"`
}

type ProductOutput struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	SKU        string `json:"sku"`
	PriceCents int64  `json:"price_cents"`
	Quantity   int    `json:"quantity"`
	Archived   bool   `json:"archived"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type ListOutput struct {
	Items []ProductOutput `json:"items"`
	Total int             `json:"total"`
	Page  int             `json:"page"`
}
