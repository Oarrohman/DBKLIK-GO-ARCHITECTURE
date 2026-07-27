# Extension Points

Each module exposes its business rules as interfaces under `domain/contracts/`
(or `domain/repositories/` for persistence). Concrete implementations live in
that module's `infrastructure/` layer and are wired up in
`infrastructure/provider/`.

Example — inventory module (`modules/inventory/domain/contracts`):

- `StockValuationMethod` — pluggable stock valuation strategy (FIFO/LIFO/average cost).
- `ReorderPolicy` — pluggable restock-decision strategy.
- `ProductRepository` (in `domain/repositories`) — pluggable persistence driver; ships with an in-memory implementation, swappable for a database-backed one.

To add a new extension point in any module: declare the interface in
`domain/contracts/`, implement it under `infrastructure/`, and inject the
chosen implementation from that module's `infrastructure/provider/`.
