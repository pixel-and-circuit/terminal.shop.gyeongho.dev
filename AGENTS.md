# shop.gyeongho.dev — Agent Guidelines

Auto-generated and maintained for AI agents (AMP, Q, Bob, and others). Last updated: 2025-02-11

## Project purpose

SSH-accessible TUI for shop at shop.gyeongho.dev. Sells all products gyeongho provides: produce, embedded devices, robots, and more. Terminal.shop-style UX: Shop (a), About (s), FAQ (d), Cart (c), scroll, product list, checkout. Go + Charm Bubble Tea; API client abstracted for mocking.

## Active technologies

- **Go** 1.21+ (single module at repo root)
- **Charm**: Bubble Tea (TUI), Lip Gloss (styling), Bubbles (components)
- **API**: `internal/apiclient.Client` interface; mock and HTTP implementations
- **Tooling**: Makefile (format, build, test), pre-commit, GitHub Actions CI

## Project structure

```text
cmd/shop/               # TUI entrypoint (main.go)
internal/
  model/               # Domain: Product, Cart, Order, StoreInfo, FAQEntry
  apiclient/            # client.go (interface), mock.go, http.go
  tui/                  # app.go, header.go, footer.go, keys.go, loader.go
  tui/pages/            # landing, shop, about, faq, cart
tests/unit/             # Unit tests (model, TUI)
tests/integration/      # Integration tests (flows with mock client)
specs/003-rebrand-to-shop/   # rebrand spec, plan, tasks
```

## Commands

| Command | Description |
|--------|-------------|
| `make format` | Format Go code (gofmt, goimports). **Run after code changes.** |
| `make build` | Build binary to `bin/shop`. **Must pass before task complete.** |
| `make test` | Run all tests (`go test ./...`) |
| `make pre-commit-install` | Install pre-commit hooks |
| `./bin/shop` | Run TUI (after build) |
| `go run ./cmd/shop` | Run TUI without building |

## Code style

- **Go**: Standard conventions; gofmt; exported names PascalCase, unexported camelCase.
- **Design**: Model-first (pure structs/interfaces in internal/model); API behind `apiclient.Client`; TUI state in `internal/tui` (Bubble Tea).
- **Quality gate**: After any code modification, run `make format` and `make build` (see [.specify/memory/constitution.md](.specify/memory/constitution.md)).

## Governance

Project constitution: [.specify/memory/constitution.md](.specify/memory/constitution.md) — code quality, testing standards, UX/UI consistency, model-first design, quality gates.

## Recent changes

- **001-ssh-shop-tui**: Full TUI implementation — navigation (a/s/d/c), Shop, About, FAQ, Cart, mock API client, HTTP client stub, Makefile, pre-commit, CI, README.

<!-- MANUAL ADDITIONS START -->
<!-- MANUAL ADDITIONS END -->
