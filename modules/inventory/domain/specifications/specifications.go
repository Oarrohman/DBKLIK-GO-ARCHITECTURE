// Package specifications encodes business rules as predicates that can be
// checked (and reused/tested) independently of any single use case.
package specifications

import "github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/entities"

func CanReceiveGoods(p *entities.Product) bool {
	return !p.Archived
}

func CanArchiveProduct(p *entities.Product) bool {
	return p.Quantity == 0
}
