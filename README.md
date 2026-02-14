# shop.gyeongho.dev

> A terminal-based shop for all products gyeongho provides: games, embedded devices, devtools, and more.

## Quick start

- **Run TUI locally**: `make build && ./bin/shop` or `go run ./cmd/shop`
- **Keys**: `a` Shop, `s` About, `d` FAQ, `c` Cart, Up/Down scroll, `q` quit
- **Access via SSH**: `ssh shop.gyeongho.dev` (when the host is deployed and the SSH service is running). No extra login step; you get the same shop TUI in your terminal.

**Prerequisites**: Go 1.21+, Make.

## Running the SSH server locally

To run the shop as an SSH server on your machine (e.g. for development):

1. Build: `make build`
2. Start the server: `./bin/shop --ssh` (listens on port 2222 by default; host key at `.ssh/id_ed25519`, created if missing)
3. From another terminal: `ssh -p 2222 -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null localhost`

You can also set `SHOP_SSH=1` and run `./bin/shop` to start in SSH mode without the `--ssh` flag.

## Develop

- **Build and run SSH server locally**: `make build && ./bin/shop --ssh` (listens on port 2222; host key at `.ssh/id_ed25519`, created if missing).
- **Connect**: `ssh -p 2222 -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null localhost`
- **Run tests** (including SSH testsession tests): `make test`. SSH-related tests are in `tests/ssh/` (connect, full flow view-product-and-add-to-cart, failure).

## Commands (Makefile)

| Target                    | Description                |
| ------------------------- | -------------------------- |
| `make format`             | Format Go code (gofmt)     |
| `make build`              | Build binary to `bin/shop` |
| `make test`               | Run tests                  |
| `make pre-commit-install` | Install pre-commit hooks   |

## Project layout

```
cmd/shop/            # TUI + SSH server entrypoint (--ssh for server mode)
internal/model/      # Domain (Product, Cart, Order, StoreInfo, FAQ)
internal/apiclient/  # API client interface + mock + HTTP
internal/sshsrv/     # Wish SSH server + shop TUI handler
internal/tui/        # Bubble Tea app, header, footer, pages
tests/unit/          # Unit tests
tests/integration/   # Integration tests
tests/ssh/           # Wish testsession tests (connect, failure)
```

## Inspiration

Inspired by [terminal.shop](https://terminal.shop), this project brings the charm of retro terminal interfaces to gyeongho's shop. Buy games, devices, devtools, and more from the terminal.
