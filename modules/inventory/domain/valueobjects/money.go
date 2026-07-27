package valueobjects

import "github.com/dbklik/dbklik-go-architecture/pkg/errors"

// Money represents a monetary amount in the smallest currency unit (cents)
// to avoid floating-point rounding errors in business calculations.
type Money struct {
	amountCents int64
}

func NewMoney(amountCents int64) (Money, error) {
	if amountCents < 0 {
		return Money{}, errors.InvalidInput("money amount must not be negative")
	}
	return Money{amountCents: amountCents}, nil
}

func (m Money) Cents() int64 { return m.amountCents }

func (m Money) Add(other Money) Money {
	return Money{amountCents: m.amountCents + other.amountCents}
}

func (m Money) Multiply(qty int) Money {
	return Money{amountCents: m.amountCents * int64(qty)}
}
