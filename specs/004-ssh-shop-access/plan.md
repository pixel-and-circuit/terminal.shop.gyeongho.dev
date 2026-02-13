# Implementation Plan: SSH Access to Terminal Shop

**Branch**: `004-ssh-shop-access` | **Date**: 2025-02-14 | **Spec**: [spec.md](./spec.md)  
**Input**: Feature specification from `/specs/004-ssh-shop-access/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command.

## Summary

Enable users to reach the terminal shop by running `ssh shop.gyeongho.dev`. The existing Bubble Tea TUI (Shop, About, FAQ, Cart) will be served over SSH using Charm Bracelet’s **Wish** library. Each SSH session gets its own TUI instance. Server development uses Wish; tests use Wish’s **testsession** package for in-process SSH client/server tests. No new domain entities; the same TUI and API client are reused behind an SSH server.

## Technical Context

**Language/Version**: Go 1.21+ (matches existing module)  
**Primary Dependencies**: Charm Bubble Tea (TUI, existing), Charm Wish (SSH server), Wish middlewares (bubbletea, activeterm, logging). API client remains `internal/apiclient.Client` (mock/HTTP).  
**Storage**: N/A (TUI state in-memory per session; API client calls backend).  
**Testing**: `go test`; unit tests for models and TUI; integration tests with mock API client; **Wish testsession** for SSH server tests (in-process client session against Wish server).  
**Target Platform**: Linux/macOS server for SSH; clients use any standard SSH client (OpenSSH).  
**Project Type**: Single (one Go module; one binary that can run as SSH server or local TUI per existing behavior).  
**Performance Goals**: Connection establishes and TUI appears within 15 seconds (spec SC-001); connection failure or timeout within 30 seconds (spec SC-002).  
**Constraints**: Wish for server, testsession for SSH-related tests; UI/UX must match existing terminal shop (header, footer, pages, keys).  
**Scale/Scope**: Single SSH service on host `shop.gyeongho.dev`; one TUI instance per SSH session.

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

Verify alignment with `.specify/memory/constitution.md`:

- **Code Quality**: Go formatting (gofmt/goimports) and Makefile (`make format`, `make build`); pre-commit and CI unchanged.
- **Testing Standards**: SSH server behavior tested with Wish testsession; existing TUI and model tests retained; deterministic tests.
- **UX/UI Consistency**: Same TUI over SSH as local; terminal shop flows (a/s/d/c, scroll, cart) unchanged.
- **Model-First**: No new domain entities; existing `internal/model` and `internal/tui.Model` reused; Wish and session handling sit at the entrypoint only.
- **Quality Gates**: `make format` and `make build` run after any code change; build must pass.

## Project Structure

### Documentation (this feature)

```text
specs/004-ssh-shop-access/
├── plan.md              # This file
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
├── contracts/           # Phase 1 (SSH access contract)
└── tasks.md             # Phase 2 output (/speckit.tasks)
```

### Source Code (repository root)

```text
cmd/
└── shop/                # Entrypoint: run as SSH server (Wish) or local TUI (existing)

internal/
├── model/               # Domain: Product, Order, Cart, StoreInfo, FAQ (unchanged)
├── apiclient/           # Client interface + mock/HTTP (unchanged)
├── tui/
│   ├── app.go           # Bubble Tea Model; reused by Wish tea handler
│   ├── header.go, footer.go, keys.go, loader.go
│   └── pages/           # landing, shop, about, faq, cart (unchanged)
└── (optional) ssh/      # If needed: Wish server wiring, host key path config

tests/
├── unit/                # Model and TUI (unchanged)
├── integration/        # Flows with mock client (unchanged)
└── (optional) ssh/     # Wish testsession tests: connect, get session, drive TUI
```

**Structure Decision**: Single Go module. SSH server is added at the entrypoint (`cmd/shop` or a dedicated main that builds the Wish server and uses `bubbletea.Middleware(teaHandler)` where `teaHandler` returns the existing `tui.Model`). Existing `internal/tui` and `internal/model` are unchanged. Tests use `github.com/charmbracelet/wish/testsession` to start the server and obtain a client session for assertions.

## Complexity Tracking

> No constitution violations. This section left empty.
