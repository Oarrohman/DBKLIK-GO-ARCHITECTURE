package unit

import (
	"testing"

	"github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/entities"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/domain/valueobjects"
)

func TestProduct_AdjustStock(t *testing.T) {
	sku, _ := valueobjects.NewSKU("SKU-1")
	price, _ := valueobjects.NewMoney(1000)
	product, err := entities.NewProduct("id-1", "Widget", sku, price, 10)
	if err != nil {
		t.Fatalf("unexpected error creating product: %v", err)
	}

	if err := product.AdjustStock(-5); err != nil {
		t.Fatalf("unexpected error adjusting stock: %v", err)
	}
	if product.Quantity != 5 {
		t.Fatalf("expected quantity 5, got %d", product.Quantity)
	}

	if err := product.AdjustStock(-100); err == nil {
		t.Fatal("expected error adjusting stock below zero, got nil")
	}
}

func TestProduct_ArchiveRequiresZeroStock(t *testing.T) {
	sku, _ := valueobjects.NewSKU("SKU-2")
	price, _ := valueobjects.NewMoney(500)
	product, _ := entities.NewProduct("id-2", "Gadget", sku, price, 3)

	if err := product.Archive(); err == nil {
		t.Fatal("expected error archiving product with stock, got nil")
	}

	_ = product.AdjustStock(-3)
	if err := product.Archive(); err != nil {
		t.Fatalf("unexpected error archiving empty-stock product: %v", err)
	}
	if !product.Archived {
		t.Fatal("expected product to be archived")
	}
}
