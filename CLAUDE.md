# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run the server
go run main.go

# Build the binary
go build -o event-booking-api

# Add/tidy dependencies
go mod tidy
```

The server starts on `localhost:8080`. There are no tests yet.

## Architecture

A minimal REST API built with [Gin](https://github.com/gin-gonic/gin) and SQLite (`api.db` in the project root).

**Request flow:** `main.go` (route handlers) → `models/` (SQL logic) → `db/db.go` (shared `*sql.DB`)

- `main.go` — registers routes and contains handler functions directly (no separate handlers package yet)
- `db/db.go` — opens the SQLite connection, sets pool limits, and runs `CREATE TABLE IF NOT EXISTS` on startup via `InitDB()`
- `models/event.go` — `Event` struct with `Save()` (INSERT) and `GetAllEvents()` (SELECT \*) using prepared statements
- `api-test/` — `.http` files for manual endpoint testing (compatible with VS Code REST Client or IntelliJ HTTP Client)

**Known stubs:** `event.ID` and `event.UserID` are hardcoded to `1` in `createEvent`; authentication and user management are not yet implemented.

## Module name

The Go module is `exammple.com/event-booking-api` (note the double-m typo) — use this exact path for internal imports.
