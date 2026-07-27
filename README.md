# DBKLIK Go Architecture

A modular Go framework following Domain-Driven Design, Clean Architecture and
Hexagonal Architecture, as defined in [ARCHITECTURE.md](ARCHITECTURE.md).

## Stack

- Go 1.26
- [Gin](https://github.com/gin-gonic/gin) for HTTP (presentation layer only — domain/application stay framework-agnostic)
- In-memory repositories by default; swap in a real database per module without touching domain/application code

## Getting started

```bash
cp .env.example .env
go mod tidy
go run ./cmd/api
```

The API listens on `:8080` (see `.env`). Health check: `GET /health`.

## Modules

Business features live under `modules/`, each isolated with its own
domain/application/infrastructure/presentation layers. See
[modules/inventory](modules/inventory/README.md) for a full reference
implementation (products + stock management).

To add a new module, copy the `inventory` module's structure, then register
its provider in [routes/api.go](routes/api.go).

## Entrypoints

| Command             | Purpose            |
|----------------------|--------------------|
| `go run ./cmd/api`    | REST API server    |
| `go run ./cmd/worker` | Background jobs    |
| `go run ./cmd/cli`    | CLI commands       |

## Testing

```bash
go test ./...
```
