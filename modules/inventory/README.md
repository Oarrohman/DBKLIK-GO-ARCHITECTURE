# Inventory Module

Reference implementation of the module structure defined in the root
[ARCHITECTURE.md](../../ARCHITECTURE.md). Manages products and stock levels.

## Layers

- `domain/` — `Product` entity, `SKU`/`Money` value objects, repository interface, stock rules.
- `application/` — `InventoryService` facade over CQRS commands/queries/handlers and single-purpose actions.
- `infrastructure/` — in-memory `ProductRepository` implementation and the module's DI provider.
- `presentation/http/` — Gin controllers, request validation, response/resource formatting.

## Endpoints

All routes are mounted under `/api/v1/inventory/products`.

| Method | Path              | Description          |
|--------|-------------------|-----------------------|
| POST   | `/`                | Create a product      |
| GET    | `/`                | List products (paged) |
| GET    | `/:id`             | Get a product         |
| PATCH  | `/:id/stock`       | Adjust stock quantity |
| POST   | `/:id/archive`     | Archive a product     |

## Swapping the persistence driver

`infrastructure/repository/ProductRepositoryMemory` is the default driver.
To back it with a real database, implement `domain/repositories.ProductRepository`
under `infrastructure/persistence/`, then swap the constructor call in
`infrastructure/provider/InventoryProvider.New`.
