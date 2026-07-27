// Package queries holds read-intent inputs for CQRS handlers.
package queries

type GetProductQuery struct {
	ProductID string
}

type ListProductsQuery struct {
	Page     int
	PageSize int
}
