# Project overview: shop.gyeongho.dev

## Purpose
SSH-accessible TUI for shop at shop.gyeongho.dev. Replicates terminal.shop-style UX: main navigation (Shop a, About s, FAQ d), scroll, product list, cart, checkout. Terminal-based shopping for gyeongho's products.

## Tech stack
- **Language**: Go 1.21+
- **TUI**: Charm Bubble Tea, Lip Gloss, Bubbles
- **API**: HTTP client abstracted (internal/apiclient); mock for tests; optional real client for shop.gyeongho.dev/api
- **Tooling**: Makefile (format, build, test), pre-commit, GitHub Actions CI

## Design
- Model-first: domain in internal/model (Product, Cart, Order, StoreInfo, FAQ); API behind interface.
- TUI in internal/tui (Bubble Tea Program, header, footer, pages).
- Single Go module at repo root. Governance: .specify/memory/constitution.md.
