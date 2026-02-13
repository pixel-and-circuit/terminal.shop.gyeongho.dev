# Implementation Plan: Responsive Terminal Layout

**Branch**: `002-responsive-terminal-layout` | **Date**: 2026-02-11 | **Spec**: [spec.md](./spec.md)  
**Input**: Feature specification from `/specs/002-responsive-terminal-layout/spec.md`  
**User addition**: Content and menu must be center-aligned to terminal size; loading screen must show a Base64-encoded random string square pattern centered; UI must match terminal.shop (reference: [IsaiahPapa/terminal.shop](https://github.com/IsaiahPapa/terminal.shop)).

## Summary

Make shop.gyeongho.dev TUI responsive and centered: (1) entire UI block (header, menu, body) centered horizontally and vertically in the terminal; (2) first screen a loading view with a centered random-character (Base64-style) rectangle matching terminal.shop’s loader; (3) layout and loading behavior aligned with terminal.shop’s loader and main layout. Implement within the existing Go + Bubble Tea TUI; use `tea.WindowSizeMsg` for dimensions and Lip Gloss placement for centering; add a loading phase before the main view.

## Technical Context

**Language/Version**: Go 1.21+ (existing)  
**Primary Dependencies**: Charm Bubble Tea, Lip Gloss (existing); use `lipgloss.Place` (or equivalent) for centering content in viewport.  
**Storage**: N/A (layout only).  
**Testing**: `go test`; extend unit/integration tests for View output (centered layout, loading rectangle) and resize behavior.  
**Target Platform**: SSH/terminal (Linux/macOS); ANSI-compatible terminal.  
**Project Type**: Single (existing Go module).  
**Performance Goals**: Resize/layout update within 2 seconds (per spec SC-002).  
**Constraints**: UI/UX must match terminal.shop: centered content block, loading screen with centered random-character box, menu and body in the center of the terminal.  
**Scale/Scope**: Same binary as 001-ssh-shop-tui; no new services or APIs.

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

Verify alignment with `.specify/memory/constitution.md`:

- **Code Quality**: Existing `make format` and style; changes follow current TUI and Lip Gloss patterns.
- **Testing Standards**: Tests for centered View output, loading view content, and resize handling; deterministic.
- **UX/UI Consistency**: Centering and loading screen follow terminal.shop layout; consistent with existing header/footer/pages.
- **Model-First**: Reuse existing Model; add or extend fields only for loading state and dimensions; no new domain entities.
- **Quality Gates**: `make format` and `make build` after any code change.

## Project Structure

### Documentation (this feature)

```text
specs/002-responsive-terminal-layout/
├── plan.md              # This file
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output (layout/viewport concepts)
├── quickstart.md        # Phase 1 output
├── contracts/           # Phase 1 (N/A or README: no new API)
└── tasks.md             # Phase 2 output (/speckit.tasks)
```

### Source Code (repository root)

No new top-level directories. Changes confined to:

```text
internal/tui/
├── app.go           # Use Width/Height to center full view; loading state → loading view
├── loader.go        # Replace with terminal.shop-style loader: centered Base64-like random rectangle
├── header.go        # No structural change; remains terminal.shop-style
├── footer.go        # No structural change
└── pages/           # Optional: ensure body content respects max width for centering

cmd/shop/main.go # Optional: trigger loading phase (e.g. delay or Init cmd) then run Program
```

**Structure Decision**: Single Go module unchanged. All work in `internal/tui` (and optionally `cmd/shop`). Layout and loading screen follow terminal.shop’s loader and handler (centered box, start_x/start_y semantics).

## Complexity Tracking

> No constitution violations. This section left empty.
