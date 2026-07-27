// Package repositories declares the domain's persistence contracts.
// Only interfaces live here; implementations belong in the
// infrastructure/repository layer.
package repositories

import "github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/entities"

type ProductRepository interface {
	FindByID(id string) (*entities.Product, error)
	FindBySKU(sku string) (*entities.Product, error)
	List(offset, limit int) ([]*entities.Product, int, error)
	Save(product *entities.Product) error
	Delete(id string) error
}
