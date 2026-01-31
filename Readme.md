# GO HTTP 🚀

**Lightweight, modular HTTP server written in Go** — a course project demonstrating clean architecture, middleware, authentication, PostgreSQL integration, and testing practices.

---

## ✨ Features

- **JWT-based authentication** with token handling and middleware
- Clean separation of concerns: handlers, services, repositories
- PostgreSQL persistence with migrations and local data directory (`postgres-data/`)
- Docker Compose support for running the database and the app
- Request validation, consistent response structure, and centralized error handling
- Unit and integration tests for core packages
- Event bus and pluggable components for extensibility

---

## 🧭 Project structure (high level)

- `cmd/` — application entrypoint(s)
- `configs/` — configuration loader
- `internal/` — core application logic and HTTP handlers (e.g. `auth`, `link`, `stat`)
- `pkg/` — reusable packages (db, jwt, middleware, req/res helpers)
- `migrations/` — database migration helpers
- `postgres-data/` — local database files (used with Docker Compose)

> Tip: Inspect handlers in `internal/`, `link/`, `stat/`, and `user/` to see the available routes and payloads.

---

## 🛠️ Tech stack

- Language: **Go** (module-enabled, see `go.mod`)
- Database: **PostgreSQL**
- Containerization: **Docker & Docker Compose**
- Testing: Go's built-in `testing` package

---

## ⚙️ Prerequisites

- Go (recommended 1.20+)
- Docker & Docker Compose (for running PostgreSQL locally)
- Git

---

## ▶️ Run locally (development)

1. Start PostgreSQL with Docker Compose:

```powershell
docker-compose up -d
```

2. Copy or set environment variables used by `configs/config.go` (if any), or edit config values directly for local runs.

3. Run the app from project root:

```powershell
# development run
go run ./cmd

# or build and run
go build -o bin/server ./cmd && .\bin\server
```

4. Check logs and open the API at the configured address (default: check `configs/config.go` for host/port).

---

## ✅ Testing

- Run all tests:

```powershell
go test ./... -v
```

- Run a specific package tests (example):

```powershell
go test ./internal/auth -v
```

---

## 📦 Docker

- `docker-compose.yml` contains the service to spin up PostgreSQL used in development and tests.
- The repository includes `postgres-data/` for local DB data (do not commit secrets to VCS).

---

## 🔒 Configuration & Secrets

- Keep sensitive values (JWT secret, DB credentials) out of source control. Use environment variables or a secrets manager for production.
- Check `configs/config.go` for keys and default values used by the app.

---

## 📚 How to extend

- Add endpoints by creating new handlers in `internal/` or other feature folders and wiring them into the router in `cmd` (or the central HTTP setup code).
- Add service logic in the `service.go` files and repositories in `repository.go`.

---
