// Package events declares domain events raised by inventory operations.
// Application-layer handlers may dispatch these onto pkg/queue.
package events

import "time"

type StockAdjusted struct {
	ProductID string
	Delta     int
	OccurredAt time.Time
}

type ProductArchived struct {
	ProductID  string
	OccurredAt time.Time
}
