# backend

Base golang backend project setup.

## Libs

### Config

<https://github.com/spf13/viper> - for env/yaml config reading

### HTTP

<https://github.com/gin-gonic/gin> - simple and popular web framework

<http://github.com/go-chi/chi> - idiomatic and lightweight web framework

<https://github.com/gofiber/fiber> - ergonomic and fast web framework

### Observability

<https://github.com/uber-go/zap> - better logging

<https://github.com/open-telemetry/opentelemetry-go> - for logs, metrics, traces

### Databases and message queues

<https://github.com/jackc/pgx> - for PostgreSQL connection

<https://github.com/redis/rueidis> - for Redis connection

<https://github.com/mongodb/mongo-go-driver> for MongoDB connection

<https://github.com/ClickHouse/clickhouse-go> for CLickHouse connection

<https://github.com/IBM/sarama> - for Kafka connection

<https://github.com/Masterminds/squirrel> - sql builder

<https://github.com/avito-tech/go-transaction-manager> - database transaction manager

### Testing

<https://github.com/stretchr/testify> - unit tests

<https://github.com/testcontainers/testcontainers-go> - integration tests

## Tools

<https://github.com/golangci/golangci-lint> - linter

<https://github.com/oapi-codegen/oapi-codegen> - for generating code from openapi specification

<https://github.com/swaggo/swag> - for generating swagger specification from code

<https://github.com/sqlc-dev/sqlc> - for generating code from sql

<https://github.com/pressly/goose> - for database migration

## Project structure

```plaitext
backend/
├── api/                    # OpenAPI specs (.yaml), generated code goes here
├── bin/                    # local tool binaries (gitignored)
├── cmd/
│   └── api/
│       └── main.go         # entrypoint - wire up deps, start server
├── docs/                   # documentation, erd, c4 and etc
├── infra/                  # infrastructure config, not Go code
│   ├── docker-compose.yaml
│   ├── otelcol/
│   ├── prometheus/
│   └── grafana/
├── internal/               # private app code, not importable from outside
│   ├── config/             # config struct, viper loading
│   ├── handler/            # HTTP handlers (one file per domain)
│   ├── service/            # business logic
│   ├── repository/         # DB queries (pgx), one file per entity
│   ├── middleware/         # HTTP middleware (auth, logging, tracing)
│   └── model/              # domain types / DTOs
├── migrations/             # SQL migration files (goose)
├── Dockerfile
├── Makefile
└── go.mod
```

### Conventions

- `api/` - source-of-truth OpenAPI spec. Run `oapi-codegen` to generate server stubs and types into `internal/`.
- `bin/` - pinned tool binaries installed via `make install-*`. Checked into gitignore so everyone gets the same version via `go install`.
- `cmd/` - one subdirectory per binary. Keeps `main.go` thin: parse config, init deps, call into `internal/`.
- `infra/` - docker compose, OTel collector config, Prometheus, Grafana dashboards. Nothing here is compiled.
- `internal/` - everything the app owns. Go enforces that nothing outside this module can import it.
  - `handler/` calls `service/`, `service/` calls `repository/`. Never skip layers.
  - `repository/` only knows about the DB. No business logic here.
  - `model/` holds plain structs shared across layers. No methods with side effects.
- `migrations/` - goose SQL files named `00001_create_users.sql`, etc. Never edit applied migrations.
