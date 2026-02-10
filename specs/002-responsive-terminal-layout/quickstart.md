# Quickstart: Responsive Terminal Layout

**Feature**: 002-responsive-terminal-layout  
**Spec**: [spec.md](./spec.md) | **Plan**: [plan.md](./plan.md)

## Prerequisites

Same as the main project: Go 1.21+, Make. See [001-ssh-mushroom-tui/quickstart.md](../001-ssh-mushroom-tui/quickstart.md) for pre-commit and CI.

## What This Feature Adds

- **Centered layout**: Menu (header) and body are centered in the terminal (horizontal and vertical). Resizing the terminal recenters the block.
- **Loading screen**: First screen shows a centered “Base64-style” random-character rectangle with an inner line (e.g. “Loading mushroom.gyeongho.dev”), matching terminal.shop’s loader layout. After a short delay (or when load completes), the main view appears, also centered.
- **Responsiveness**: Layout updates on `tea.WindowSizeMsg`; no refresh or reload required.

## Commands (unchanged)

| Target | Description |
|--------|-------------|
| `make format` | Format Go code. Run after changes (constitution). |
| `make build`  | Build TUI binary. |
| `make test`   | Run all tests (including new layout/loading tests). |

## Run TUI and Verify Layout

```bash
make build
./bin/mushroom
```

1. **Loading screen**: You should see a centered rectangle of random characters with a loading message in the middle. It may animate (new random chars) for ~2 seconds.
2. **Main view**: Then the main view appears: header (mushroom, a shop, s about, d faq, cart) and body centered in the terminal.
3. **Resize**: Resize the terminal window; the whole block (header + menu + body + footer) should recenter and remain usable without horizontal overflow.

## Verify Across Sizes

- **Small**: Run in a terminal ~80×24; content should stay centered and readable.
- **Large**: Run in a wide terminal (e.g. 160×40); the content block should stay centered and not stretch to full width (capped content width, terminal.shop-style).
- **Resize**: Shrink and expand; layout should update within a couple of seconds (per spec SC-002).

## Run Tests

```bash
make test
```

New or updated tests should cover:

- View output when `Loading == true` (loading rectangle present, centered).
- View output when `Loading == false` (main view centered; header/footer present).
- Resize: `tea.WindowSizeMsg` updates dimensions and subsequent View uses them for centering.

## Reference

- Layout reference: [IsaiahPapa/terminal.shop](https://github.com/IsaiahPapa/terminal.shop) (loader, handler layout).
- Research: [research.md](./research.md) (Lip Gloss Place, loading screen, Base64 rectangle).
- Data model: [data-model.md](./data-model.md) (viewport, content block, loading state).
