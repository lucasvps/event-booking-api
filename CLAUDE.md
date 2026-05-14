# CLAUDE.md

## Commands
```bash
go run main.go        # server on localhost:8080
go build -o event-booking-api
go mod tidy
```
No tests yet.

## Architecture
Gin + SQLite (`api.db`). Flow: `main.go` → `models/` → `db/db.go`.

- `main.go` — routes + handlers (no separate handlers pkg)
- `db/db.go` — SQLite connection, pool limits, `InitDB()` runs CREATE TABLE IF NOT EXISTS
- `models/event.go` — `Event` struct, `Save()` (INSERT), `GetAllEvents()` (SELECT *)
- `api-test/` — `.http` files for VS Code REST Client / IntelliJ

**Stubs:** `event.ID` and `event.UserID` hardcoded to `1`; auth/users not implemented.

## Module
`exammple.com/event-booking-api` (double-m typo) — use exact for internal imports.
