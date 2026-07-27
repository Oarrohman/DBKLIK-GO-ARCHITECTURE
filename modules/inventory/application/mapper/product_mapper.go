// Package mapper converts between DTOs and domain entities so neither
// side needs to know about the other's shape.
package mapper

import (
	"time"

	"github.com/dbklik/dbklik-go-architecture/modules/inventory/application/dto"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/entities"
)

func ToOutput(p *entities.Product) dto.ProductOutput {
	return dto.ProductOutput{
		ID:         p.ID,
		Name:       p.Name,
		SKU:        p.SKU.String(),
		PriceCents: p.Price.Cents(),
		Quantity:   p.Quantity,
		Archived:   p.Archived,
		CreatedAt:  p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  p.UpdatedAt.Format(time.RFC3339),
	}
}

func ToOutputList(products []*entities.Product) []dto.ProductOutput {
	items := make([]dto.ProductOutput, 0, len(products))
	for _, p := range products {
		items = append(items, ToOutput(p))
	}
	return items
}
