# Implementation Plan: SSH-Based Mushroom Sales TUI

**Branch**: `001-ssh-mushroom-tui` | **Date**: 2025-02-11 | **Spec**: [spec.md](./spec.md)  
**Input**: Feature specification from `/specs/001-ssh-mushroom-tui/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command.

## Summary

Deliver an SSH-accessible TUI for mushroom sales at mushroom.gyeongho.dev that replicates the UI/UX of the terminal.shop clone ([IsaiahPapa/terminal.shop](https://github.com/IsaiahPapa/terminal.shop)): main navigation (Shop a, About s, FAQ d), scroll, product list, cart, and checkout. Implement in Go using the Charm (Bubble Tea) library. API base is mushroom.gyeongho.dev with paths under `/api`; the server is not yet built, so the HTTP client is abstracted behind interfaces for easy mocking in tests. Development follows TDD: user scenarios (menu navigation, scroll, add-to-cart) drive test cases first, then implementation. Project tooling: Makefile for format/build/test, pre-commit for format and commitlint, and GitHub Actions CI. README is updated to reflect the final project layout.

## Technical Context

**Language/Version**: Go 1.21+ (or latest stable)  
**Primary Dependencies**: Charm Bubble Tea (TUI), Lip Gloss (styling), Bubbles (components). HTTP client abstracted via interface (e.g. `internal/apiclient.Client`).  
**Storage**: N/A for initial TUI; API client calls backend at mushroom.gyeongho.dev/api (mock in tests).  
**Testing**: `go test`; unit tests for models and TUI Update/View logic; integration/contract tests with mocked HTTP client.  
**Target Platform**: SSH session (Linux/macOS terminal); ANSI-compatible terminal.  
**Project Type**: Single (one Go module at repo root).  
**Performance Goals**: Responsive key handling and scroll; no strict latency targets for initial scope.  
**Constraints**: UI/UX must match terminal.shop clone flows (header, footer, pages, keys a/s/d, up/down scroll, cart).  
**Scale/Scope**: Single TUI binary served over SSH; API surface under /api (products, about, faq, cart, order).

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

Verify alignment with `.specify/memory/constitution.md`:

- **Code Quality**: Go formatting (gofmt/goimports) and Makefile targets (`make format`, `make build`); pre-commit enforces format and commitlint.
- **Testing Standards**: TDD; tests for menu navigation, scroll, add-to-cart; HTTP client mocked in tests.
- **UX/UI Consistency**: UI flows and key bindings follow terminal.shop clone (header, footer, Shop/About/FAQ, scroll).
- **Model-First**: Domain models as Go structs (Product, Order, Cart, StoreInfo, FAQ); API client and TUI state behind interfaces.
- **Quality Gates**: Makefile provides `make format` and `make build`; agents run both after code changes.

## Project Structure

### Documentation (this feature)

```text
specs/001-ssh-mushroom-tui/
├── plan.md              # This file
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
├── contracts/           # Phase 1 API contracts
└── tasks.md             # Phase 2 output (/speckit.tasks)
```

### Source Code (repository root)

```text
cmd/
└── mushroom/            # Main entry (SSH TUI binary)

internal/
├── model/               # Domain: Product, Order, Cart, StoreInfo, FAQ
├── apiclient/           # HTTP client interface + real/mock implementations
├── tui/
│   ├── app.go           # Bubble Tea Program, root model
│   ├── header.go        # Header (logo, a shop, s about, d faq, cart)
│   ├── footer.go       # Footer (+/- qty, c cart, q quit)
│   ├── pages/
│   │   ├── landing.go
│   │   ├── shop.go     # Product list, scroll, selection, add to cart
│   │   ├── about.go
│   │   ├── faq.go      # Scrollable FAQ
│   │   └── cart.go     # Cart view, checkout
│   └── keys.go         # Key binding constants
└── ...

tests/
├── unit/                # Model and TUI Update/View logic
├── integration/        # Full TUI flows with mock API client
└── contract/           # Optional: API contract tests when server exists

Makefile                # format, build, test, pre-commit install
.pre-commit-config.yaml # format + commitlint
.github/workflows/       # CI (format, build, test)
go.mod
go.sum
README.md
```

**Structure Decision**: Single Go module. Domain lives in `internal/model`; API access behind `internal/apiclient` interface for testability. TUI in `internal/tui` mirrors terminal.shop clone layout (header, footer, pages). Binary is `cmd/mushroom` for SSH server to invoke.

## Complexity Tracking

> No constitution violations. This section left empty.
