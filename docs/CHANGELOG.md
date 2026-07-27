# Changelog

## [0.1.0] - Initial scaffold

- Root framework structure per ARCHITECTURE.md (cmd, internal, pkg, routes, database, resources, configs, tests).
- Reference `inventory` module (products + stock) with full domain/application/infrastructure/presentation layers.
- Gin-based REST API entrypoint with graceful shutdown, health check, and shared middleware.
