# Data Model: Responsive Terminal Layout

**Feature**: 002-responsive-terminal-layout  
**Source**: [spec.md](./spec.md) (layout/UX requirements); no new domain entities.

---

## Layout Concepts (TUI State)

This feature does not introduce new domain entities (Product, Cart, etc.); those remain defined in [001-ssh-mushroom-tui/data-model.md](../001-ssh-mushroom-tui/data-model.md). The following describe layout/viewport state used by the TUI for centering and responsiveness.

### Viewport (terminal dimensions)

| Concept    | Description |
|-----------|-------------|
| Width     | Terminal width in columns (from `tea.WindowSizeMsg`). |
| Height    | Terminal height in rows (from `tea.WindowSizeMsg`). |

Used to center the UI block and the loading rectangle. Stored on the root model as `Width`, `Height` (already present).

### Content block (main view)

| Concept     | Description |
|------------|-------------|
| Content width  | Max width of the centered block (e.g. 60), to avoid over-stretching on very wide terminals (terminal.shop behavior). |
| Content height | Logical height of header + body + footer; used with viewport height to compute vertical centering. |

No new structs required; the full view string (header + body + footer) is passed to `lipgloss.Place(viewportWidth, viewportHeight, Center, Center, view)` so the block is centered.

### Loading view

| Concept     | Description |
|------------|-------------|
| Loading    | Boolean: when true, show loading screen; when false, show main view. |
| Box size   | Loading rectangle dimensions (e.g. 40×20) for the random-character box, centered in viewport. |
| Inner text | Centered line inside the box (e.g. "Loading mushroom.gyeongho.dev" or "LOADING STORE"). |

Implementation: build a string representing the rectangle (random Base64-style chars, inner line centered), then center that string in the viewport with Lip Gloss. No persistent storage; loading state is in-memory only.

---

## State transitions

- **Start**: `Loading = true` → only loading view is rendered.
- **After delay or “load complete”**: `Loading = false` → main view (header + body + footer) is rendered, centered.
- **On resize**: `tea.WindowSizeMsg` updates `Width`, `Height`; next `View()` uses new dimensions for centering (no explicit state machine beyond Loading flag and dimensions).

---

## Validation rules

- Viewport: `Width` and `Height` must be positive; if not yet received, use defaults (e.g. 80×24).
- Content width: cap at a maximum (e.g. 60) when computing inner layout so that body text does not exceed terminal.shop-style width on large terminals.
