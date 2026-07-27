# Domain Events

Modules declare domain events in `domain/events/` as plain structs and raise
them from application-layer handlers/actions. Dispatch them onto `pkg/queue`
(in-memory by default) from the module's `infrastructure/provider/`.

## Inventory module

- `StockAdjusted` — raised when a product's quantity changes.
- `ProductArchived` — raised when a product is archived.
