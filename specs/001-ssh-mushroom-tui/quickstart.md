# Quickstart: SSH-Based Mushroom Sales TUI

**Feature**: 001-ssh-mushroom-tui  
**Spec**: [spec.md](./spec.md) | **Plan**: [plan.md](./plan.md)

## Prerequisites

- Go 1.21+
- Make
- Optional: [pre-commit](https://pre-commit.com/) (for format + commitlint hooks)  
  - Install: `pip install pre-commit` or `brew install pre-commit`  
  - Commitlint (conventional commits) may require Node/npx; see `.pre-commit-config.yaml` when added.

## Repository Layout (after implementation)

```text
cmd/mushroom/          # Main binary (TUI entry)
internal/model/        # Domain structs
internal/apiclient/    # HTTP client interface + real/mock
internal/tui/          # Bubble Tea app, header, footer, pages
tests/                 # unit, integration
Makefile
```

## Commands (Makefile)

| Target | Description |
|--------|-------------|
| `make format` | Format Go code (gofmt/goimports). Run after code changes (constitution). |
| `make build` | Build TUI binary (e.g. `bin/mushroom`). |
| `make test` | Run all tests. |
| `make pre-commit-install` | Install pre-commit hooks (format + commitlint). |

## Run TUI Locally

Until SSH is configured, run the binary directly:

```bash
make build
./bin/mushroom
```

Or:

```bash
go run ./cmd/mushroom
```

**Keys**: **a** Shop, **s** About, **d** FAQ; **c** Cart; **Up/Down** scroll; **q** or **Ctrl+C** quit. Footer: +/- qty, c cart, q quit (terminal.shop-style).

## Run Tests

```bash
make test
# or
go test ./...
```

Unit tests cover model and TUI Update/View logic; integration tests use a mock API client.

## Pre-commit and CI

- **Pre-commit**: After adding `.pre-commit-config.yaml`, run `pre-commit install`. Each commit will run format and commitlint.
- **CI**: GitHub Actions runs on push/PR: format check, `make build`, `make test`.

## API Base URL

Default: `https://mushroom.gyeongho.dev/api`. Override via env (e.g. `MUSHROOM_API_BASE`) for local or mock server. When the server is not running, the client uses mock data or an in-memory implementation for development and tests.

## Reference

- UI/UX reference: [IsaiahPapa/terminal.shop](https://github.com/IsaiahPapa/terminal.shop)  
- API contract: [contracts/openapi.yaml](./contracts/openapi.yaml)  
- Data model: [data-model.md](./data-model.md)
