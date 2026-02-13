# shop.gyeongho.dev

> A terminal-based shop for all products gyeongho provides: mushrooms, embedded devices, robots, and more.

## Quick start

- **Run TUI**: `make build && ./bin/shop` or `go run ./cmd/shop`
- **Keys**: `a` Shop, `s` About, `d` FAQ, `c` Cart, Up/Down scroll, `q` quit
- **SSH** (when server is configured): `ssh -a -i /dev/null shop.gyeongho.dev`

**Prerequisites**: Go 1.21+, Make.

## Commands (Makefile)

| Target                    | Description                    |
| ------------------------- | ------------------------------ |
| `make format`             | Format Go code (gofmt)         |
| `make build`              | Build binary to `bin/shop`     |
| `make test`               | Run tests                      |
| `make pre-commit-install` | Install pre-commit hooks       |

## Project layout

```
cmd/shop/            # TUI entrypoint
internal/model/      # Domain (Product, Cart, Order, StoreInfo, FAQ)
internal/apiclient/  # API client interface + mock + HTTP
internal/tui/        # Bubble Tea app, header, footer, pages
tests/unit/          # Unit tests
tests/integration/   # Integration tests
```

## Inspiration

Inspired by [terminal.shop](https://terminal.shop), this project brings the charm of retro terminal interfaces to gyeongho's shop. Buy mushrooms, devices, robots, and more from the terminal.
