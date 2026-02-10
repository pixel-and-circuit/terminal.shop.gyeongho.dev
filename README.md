# mushroom.gyeongho.dev

> A nostalgic terminal-based shopping experience for dad's mushroom farm.

## Quick start

- **Run TUI**: `make build && ./bin/mushroom` or `go run ./cmd/mushroom`
- **Keys**: `a` Shop, `s` About, `d` FAQ, `c` Cart, Up/Down scroll, `q` quit
- **SSH** (when server is configured): `ssh -a -i /dev/null mushroom.gyeongho.dev`

See [specs/001-ssh-mushroom-tui/quickstart.md](specs/001-ssh-mushroom-tui/quickstart.md) for full setup, Makefile targets, and tests.

## Commands (Makefile)

| Target                    | Description                    |
| ------------------------- | ------------------------------ |
| `make format`             | Format Go code (gofmt)         |
| `make build`              | Build binary to `bin/mushroom` |
| `make test`               | Run tests                      |
| `make pre-commit-install` | Install pre-commit hooks       |

## Project layout

```
cmd/mushroom/         # TUI entrypoint
internal/model/       # Domain (Product, Cart, Order, StoreInfo, FAQ)
internal/apiclient/   # API client interface + mock + HTTP
internal/tui/         # Bubble Tea app, header, footer, pages
tests/unit/           # Unit tests
tests/integration/    # Integration tests
```

## Inspiration

Inspired by [terminal.shop](https://terminal.shop), this project brings the charm of retro terminal interfaces to my father's mushroom farm back home. Because buying mushrooms should be more fun than scrolling through a boring web page!
