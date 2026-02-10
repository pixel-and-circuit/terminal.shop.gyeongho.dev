# Suggested Commands (mushroom.gyeongho.dev)

## Format & quality
- `make format` — format Go code (gofmt, goimports if available). Run after code changes (constitution).
- `make build` — build TUI binary to `bin/mushroom`. Must pass before considering task complete.
- `make test` — run all tests (`go test ./...`).
- `make pre-commit-install` — install pre-commit hooks (format + checks).

## Run application
- `./bin/mushroom` — run TUI (after `make build`).
- `go run ./cmd/mushroom` — run without building binary.

## Go
- `go test ./...` — run tests.
- `go build -o bin/mushroom ./cmd/mushroom` — build.
- `gofmt -s -w .` — format.

## System (Darwin)
- `git`, `ls`, `cd`, `grep`, `find` — standard Unix commands available.
