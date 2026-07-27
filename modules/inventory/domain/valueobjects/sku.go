package valueobjects

import (
	"strings"

	"github.com/dbklik/dbklik-go-architecture/pkg/errors"
)

// SKU is an immutable, validated stock-keeping unit identifier.
type SKU struct {
	value string
}

func NewSKU(raw string) (SKU, error) {
	v := strings.ToUpper(strings.TrimSpace(raw))
	if v == "" {
		return SKU{}, errors.InvalidInput("sku must not be empty")
	}
	return SKU{value: v}, nil
}

func (s SKU) String() string { return s.value }

func (s SKU) Equals(other SKU) bool { return s.value == other.value }
