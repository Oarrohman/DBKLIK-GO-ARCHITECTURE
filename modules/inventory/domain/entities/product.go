package entities

import (
	"time"

	"github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/valueobjects"
	"github.com/dbklik/dbklik-go-architecture/pkg/errors"
)

// Product is the inventory domain's core entity. It carries business
// behavior (stock adjustment rules) — never persistence or HTTP concerns.
type Product struct {
	ID        string
	Name      string
	SKU       valueobjects.SKU
	Price     valueobjects.Money
	Quantity  int
	Archived  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProduct(id, name string, sku valueobjects.SKU, price valueobjects.Money, quantity int) (*Product, error) {
	if name == "" {
		return nil, errors.InvalidInput("product name must not be empty")
	}
	if quantity < 0 {
		return nil, errors.InvalidInput("product quantity must not be negative")
	}

	now := time.Now()
	return &Product{
		ID:        id,
		Name:      name,
		SKU:       sku,
		Price:     price,
		Quantity:  quantity,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// AdjustStock changes quantity by delta (positive to receive goods,
// negative to ship/consume), refusing to let quantity go negative.
func (p *Product) AdjustStock(delta int) error {
	if p.Archived {
		return errors.Conflict("cannot adjust stock of an archived product")
	}
	if p.Quantity+delta < 0 {
		return errors.InvalidInput("insufficient stock for this adjustment")
	}
	p.Quantity += delta
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Product) Archive() error {
	if p.Quantity > 0 {
		return errors.Conflict("cannot archive a product that still has stock")
	}
	p.Archived = true
	p.UpdatedAt = time.Now()
	return nil
}
