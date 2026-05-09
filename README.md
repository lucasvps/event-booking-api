# event-booking-api

A REST API for event booking built with Go — a study project based on a basic Go course.

## About

This project was developed for learning purposes while following a foundational Go course. It covers REST API design, JWT authentication, SQLite persistence, and a clean package-by-layer architecture.

## Tech Stack

- **Go** (1.22+)
- **Gin** — HTTP web framework
- **SQLite** — embedded database via `go-sqlite3`
- **JWT** — authentication via `golang-jwt/jwt`
- **bcrypt** — password hashing via `golang.org/x/crypto`

## Architecture

The project uses a **package-by-layer** structure, separating each concern into its own package:

```
cmd/api/              → entry point (main.go)
db/                   → database connection and table initialization
internal/
  domain/             → domain structs (Event, User)
  handler/            → HTTP handlers (events, users, registrations, auth)
  middlewares/        → JWT authentication middleware
  repository/         → data access layer (SQL queries)
pkg/utils/            → shared utilities (JWT, bcrypt)
routes/               → Gin route registration
```

## Running

```bash
go run cmd/api/main.go
```

The server starts on `http://localhost:8080`.

## Endpoints

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | `/events` | No | List all events |
| GET | `/events/:id` | No | Get event by ID |
| POST | `/events` | Yes | Create event |
| PUT | `/events/:id` | Yes | Update event |
| DELETE | `/events/:id` | Yes | Delete event |
| POST | `/events/:id/register` | Yes | Register for event |
| GET | `/events/:id/register` | No | List registrations |
| DELETE | `/events/:id/register` | Yes | Cancel registration |
| POST | `/signup` | No | Create user account |
| GET | `/users` | No | List all users |
| POST | `/login` | No | Login and receive JWT |

Authenticated routes require an `Authorization: <token>` header with a valid JWT obtained from `/login`.

## Manual Testing

The `api-test/` directory contains `.http` files compatible with the VS Code REST Client and IntelliJ HTTP Client.
