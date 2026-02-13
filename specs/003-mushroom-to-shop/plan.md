# Implementation Plan: Rebrand from Mushroom to Shop

**Branch**: `003-mushroom-to-shop` | **Date**: 2025-02-13 | **Spec**: [spec.md](./spec.md)  
**Input**: Feature specification from `specs/003-mushroom-to-shop/spec.md`

**User input (planning)**: GitHub repository path will change later; Go package name must change from mushroom to shop. Search "mushroom" across all files and apply appropriate changes. README must describe that gyeongho sells all products. Comprehensively investigate UI and test impact and apply changes correctly.

## Summary

Rebrand the TUI from "mushroom" to "shop": (1) change Go module from `mushroom.gyeongho.dev` to `shop.gyeongho.dev` and update all imports and the `cmd/` directory name to `cmd/shop` with binary `bin/shop`; (2) replace all user-facing and internal references to "mushroom" (store name, domain, loader text, about/FAQ copy) with "shop" and shop.gyeongho.dev; (3) update README and AGENTS.md to describe the shop as selling all products gyeongho provides (mushrooms, embedded devices, robots, etc.); (4) update tests and fixtures (import paths, loader assertions, about/store content) so UI and tests remain correct.

## Technical Context

**Language/Version**: Go 1.21+ (go.mod: 1.25.7)  
**Primary Dependencies**: Charm Bubble Tea, Lip Gloss, Bubbles  
**Storage**: N/A (API client; mock and HTTP)  
**Testing**: `go test ./...` (unit: `tests/unit/`, integration: `tests/integration/`, e2e: `tests/e2e/`)  
**Target Platform**: SSH-accessible terminal (TUI)  
**Project Type**: single (one Go module at repo root)  
**Performance Goals**: Same as existing TUI (responsive, no regressions)  
**Constraints**: Quality gate `make format` and `make build` must pass; tests must pass  
**Scale/Scope**: Single binary; all source under `cmd/mushroom` → `cmd/shop`, `internal/`, `tests/`

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

Verify alignment with `.specify/memory/constitution.md`:

- **Code Quality**: Rename and string changes preserve formatting; `make format` and style consistency applied after edits.
- **Testing Standards**: All tests updated (import paths, assertions on loader text and store content); deterministic; coverage preserved.
- **UX/UI Consistency**: User-facing strings (loader, about, FAQ, header/footer) use "shop" / shop.gyeongho.dev; terminal UX unchanged.
- **Model-First**: No change to domain model interfaces/structs; only naming and copy (StoreInfo content, loader constant).
- **Quality Gates**: Plan assumes `make format` and `make build` run after all code changes.

## Project Structure

### Documentation (this feature)

```text
specs/003-mushroom-to-shop/
├── plan.md              # This file
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
├── contracts/           # Phase 1 output (no API change; README only)
└── tasks.md             # Phase 2 output (/speckit.tasks – not created by /speckit.plan)
```

### Source Code (repository root)

```text
cmd/shop/                 # TUI entrypoint (renamed from cmd/mushroom)
internal/
  model/                 # Domain (Product, Cart, Order, StoreInfo, FAQ) – unchanged structure
  apiclient/             # client, mock, http – import paths and copy (about, default URL) updated
  tui/                   # app, header, footer, loader, pages – loader text and imports updated
tests/
  unit/                  # Import paths + assertions (loader text, product/store content)
  integration/           # Import paths + assertions (loader, cart product names, about)
  e2e/                   # Import paths + fixtures (about.json, product/cart/order data as needed)
bin/shop                 # Build output (Makefile: bin/shop from cmd/shop)
```

**Structure Decision**: Single Go module; directory rename `cmd/mushroom` → `cmd/shop` and module rename `mushroom.gyeongho.dev` → `shop.gyeongho.dev`. No new packages or directories; only renames and string/copy updates.

## Complexity Tracking

No constitution violations. Rename and copy changes only; quality gate and test strategy unchanged.
