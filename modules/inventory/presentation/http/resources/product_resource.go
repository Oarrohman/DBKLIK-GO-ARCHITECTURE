// Package resources transforms application DTOs into the module's public
// JSON API shape, keeping the wire format independent of internal naming.
package resources

import "github.com/dbklik/dbklik-go-architecture/modules/inventory/application/dto"

type ProductResource struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	SKU      string `json:"sku"`
	Price    int64  `json:"price_cents"`
	Quantity int    `json:"quantity"`
	Archived bool   `json:"archived"`
}

func FromOutput(o dto.ProductOutput) ProductResource {
	return ProductResource{
		ID:       o.ID,
		Name:     o.Name,
		SKU:      o.SKU,
		Price:    o.PriceCents,
		Quantity: o.Quantity,
		Archived: o.Archived,
	}
}

func FromOutputList(list []dto.ProductOutput) []ProductResource {
	items := make([]ProductResource, 0, len(list))
	for _, o := range list {
		items = append(items, FromOutput(o))
	}
	return items
}
