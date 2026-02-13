# Quickstart: Rebrand from Mushroom to Shop

**Feature**: 003-mushroom-to-shop  
**Spec**: [spec.md](./spec.md) | **Plan**: [plan.md](./plan.md)

## Prerequisites

Same as the main project: Go 1.21+, Make. See [001-ssh-mushroom-tui/quickstart.md](../001-ssh-mushroom-tui/quickstart.md) for pre-commit and CI.

## Repository layout (after implementation)

```text
cmd/shop/              # TUI entrypoint (renamed from cmd/mushroom)
internal/model/        # Domain structs (unchanged)
internal/apiclient/    # Client interface + mock + HTTP (imports and copy updated)
internal/tui/          # Bubble Tea app, loader, pages (loader text updated)
tests/                 # unit, integration, e2e (imports and assertions updated)
Makefile               # build → bin/shop, run → cmd/shop
```

## Commands (Makefile)

| Target            | Description                                      |
|-------------------|--------------------------------------------------|
| `make format`     | Format Go code. Run after code changes.          |
| `make build`      | Build TUI binary to `bin/shop`.                  |
| `make test`       | Run all tests.                                  |
| `make pre-commit-install` | Install pre-commit hooks.                |

## Run TUI locally

```bash
make build
./bin/shop
```

Or:

```bash
go run ./cmd/shop
```

**Keys**: **a** Shop, **s** About, **d** FAQ, **c** Cart; **Up/Down** scroll; **q** or **Ctrl+C** quit.

Loader shows **"Loading shop.gyeongho.dev"**; about and store copy reflect the shop (shop.gyeongho.dev, all products gyeongho provides).

## Run tests

```bash
make test
# or
go test ./...
```

After the rebrand, all tests use import path `shop.gyeongho.dev/...` and assert on shop branding (loader text, about content) where relevant.

## SSH (when configured)

Once the server is configured for the new domain:

```bash
ssh -a -i /dev/null shop.gyeongho.dev
```

## Quality gate

After any code change: `make format` and `make build`. Build must pass before considering the change complete (see [.specify/memory/constitution.md](../../.specify/memory/constitution.md)).
