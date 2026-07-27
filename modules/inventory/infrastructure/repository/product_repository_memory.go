// Package repository implements the domain's repository interfaces.
// ProductRepositoryMemory is the default driver used when no external
// database is configured; swap in a gorm/sqlx-backed implementation
// under infrastructure/persistence without touching domain or application code.
package repository

import (
	"sort"
	"sync"

	"github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/entities"
)

type ProductRepositoryMemory struct {
	mu       sync.RWMutex
	products map[string]*entities.Product
}

func NewProductRepositoryMemory() *ProductRepositoryMemory {
	return &ProductRepositoryMemory{products: make(map[string]*entities.Product)}
}

func (r *ProductRepositoryMemory) FindByID(id string) (*entities.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.products[id], nil
}

func (r *ProductRepositoryMemory) FindBySKU(sku string) (*entities.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, p := range r.products {
		if p.SKU.String() == sku {
			return p, nil
		}
	}
	return nil, nil
}

func (r *ProductRepositoryMemory) List(offset, limit int) ([]*entities.Product, int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	all := make([]*entities.Product, 0, len(r.products))
	for _, p := range r.products {
		all = append(all, p)
	}
	sort.Slice(all, func(i, j int) bool { return all[i].CreatedAt.Before(all[j].CreatedAt) })

	total := len(all)
	if offset >= total {
		return []*entities.Product{}, total, nil
	}
	end := offset + limit
	if end > total {
		end = total
	}
	return all[offset:end], total, nil
}

func (r *ProductRepositoryMemory) Save(product *entities.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.products[product.ID] = product
	return nil
}

func (r *ProductRepositoryMemory) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.products, id)
	return nil
}
