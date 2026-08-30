# AGENTS.md

Go backend for a personal-portfolio site. Gin + PostgreSQL via raw `database/sql` (`lib/pq`). No ORM, no migration tool, no tests.

## Commands

- Build / verify: `go build ./...` (this is the only real check — no Makefile, CI, lint, or test suite)
- Run: `go run ./cmd/main.go`
- Hot reload: `air` (config in `.air.toml`; builds `./cmd/main.go` → `tmp/main`)
- Full stack: `docker compose up` from the repo root (parent dir). Backend `Dockerfile` builds via `go build -o main ./cmd`.

## Env / config gotchas

- `config.LoadConfig` loads `.env` then `.env.access_credentials` with godotenv and calls `os.Exit(0)` if either is missing — unless `RUNNING_IN_DOCKER=true`, which skips both files.
- Both env files are gitignored. Copy from `.env.example` and `.env.access_credentials.example`.
- Required vars: `SERVER_PORT`, `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `PROFILE_DB_PATH`, `ALLOWED_ORIGIN_AUTHORIZED_ROUTES`; auth uses `ACCESS_KEY`/`AUTH_KEY`/`AUTH_TOKEN`; storage uses `IMAGES_STORAGE_PATH`/`DOCUMENTS_STORAGE_PATH`/`ICONS_STORAGE_PATH`.

## Two data stores (critical)

- **Postgres** — blogs, projects, technologies, tech_types, experiences, daily_visit. Schema comes only from `postgres-init/init.sql` (a pg_dump, loaded on first container init). No migration tooling; edit `init.sql` for schema changes.
- **Profile data** (name, summary, bio, skills) is **not** in Postgres — it is a JSON file at `PROFILE_DB_PATH` (`internal/db/profile_data.json`), read/written wholesale by `ProfileRepository`. The Dockerfile copies `internal/db/profile_data.docker.json` → `data/profile_data.json` instead.

## Auth flow gotcha

- `/auth/login` takes an `ACCESS_KEY`; success returns a JWT signed with `AUTH_KEY`.
- `auth.AuthMiddleware` requires the request token to **both** equal `os.Getenv("AUTH_TOKEN")` and pass JWT signature validation. `AUTH_TOKEN` is set in-memory via `os.Setenv` after login/renew — it is never persisted, so it resets on restart (single in-process admin session).

## Layout & conventions

- Layered: `routes` → `handlers` → `services` → `repositories`. Handlers are wired manually in `internal/handlers/handler.go`; new endpoints must be registered in both `internal/routes/routes.go` and `handler.go`.
- API response envelope is `{message, data, errors}` — use `internal/response.Success`/`response.Error`, not raw `c.JSON`.
- Raw SQL with `$1` placeholders in repositories; dynamic updates build SQL strings via `utils.GenerateSingleUpdateQuery`.
- Storage files are served from `storage/{images,documents,icons}`, joined to `os.Getwd()` in `internal/handlers/storage_handler.go`.
- CORS allows a single origin from `ALLOWED_ORIGIN_AUTHORIZED_ROUTES`.

## Module path note

`go.mod` module is `github.com/husni-robani/abstracted_self/backend` while the directory is `abstracted_self-backend/backend`. Import paths use the module path, not the filesystem path.

## Repo-local opencode

`opencode.json` runs an MCP `db-adapter` (Postgres) and `api-adapter` (OpenAPI spec served at `http://localhost:9001`), server port 9001. `.opencode/` also holds custom agents and `git-commit`/`grill-me` skills.
