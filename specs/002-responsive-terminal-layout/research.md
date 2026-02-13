# Research: Responsive Terminal Layout

**Feature**: 002-responsive-terminal-layout  
**Purpose**: Resolve centering strategy, loading screen layout (terminal.shop reference), and Base64-style random rectangle.

---

## 1. Centering TUI Content in Bubble Tea (Lip Gloss)

**Decision**: Use `lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, fullView)` to center the entire UI block (header + body + footer) in the terminal viewport.

**Rationale**: Spec and user input require menu and body to be in the center of the terminal; FR-001/FR-002 require layout to adapt to terminal size. Bubble Tea provides `tea.WindowSizeMsg` with Width and Height; we already store these in the model. Lip Gloss’s `Place` takes viewport dimensions and alignment and returns a string with the content centered, so the final `View()` output is already padded and centered.

**Findings**:

- **API**: `lipgloss.Place(width, height, horizontalAlign, verticalAlign, content)`. Use `lipgloss.Center` for both to center the content block in the terminal.
- **Dimensions**: Use `m.Width` and `m.Height` from `tea.WindowSizeMsg` (already handled in `app.go`). On startup, Bubble Tea may send an initial size; if not, default 80×24 is acceptable until first resize.
- **Content width**: terminal.shop uses a fixed content width (60) inside a centered block; we can keep a max content width (e.g. 60) so that on very wide terminals the block does not stretch arbitrarily, matching terminal.shop’s `handler.rs` (width 60, start_x/start_y centered).

**Alternatives considered**: Manual padding with spaces and newlines is error-prone and duplicates logic; Lip Gloss `Place` is the standard and keeps layout declarative.

---

## 2. Terminal.shop Loading Screen and Main Layout (Reference)

**Decision**: Replicate the layout semantics of [IsaiahPapa/terminal.shop](https://github.com/IsaiahPapa/terminal.shop): (1) loading screen = centered box of random characters with an inner “LOADING STORE” (or equivalent) line; (2) main UI = centered content block (header, body, footer) with start_x/start_y derived from terminal size.

**Rationale**: User requirement: “UI는 terminal.shop과 동일해야 하며”, “로딩화면, 메인화면 등의 UI 배치를 확인하여 해당 배치를 그대로 따라해야 한다.” GitHub inspection of the repo is the source of truth.

**Findings (from terminal.shop)**:

- **Loader** (`src/ui/loader.rs`):
  - Box size: 40 cols × 20 rows, centered: `start_x = (term_width - 40) / 2`, `start_y = (term_height - 20) / 2`.
  - Inner box: 20×3, with “LOADING STORE” centered on the middle line.
  - Outer area: random alphanumeric characters (A–Z, a–z, 0–9), regenerated every 250ms for 8 iterations (2 seconds total). Inner box: middle line shows “LOADING STORE” (bold white); other inner lines are spaces. Styling: outer chars DarkGrey, inner text bold white.
- **Main layout** (`src/ui/handler.rs`):
  - Content block: width 60, height 45. `start_x = (term_width - 60) / 2`, `start_y = (term_height - 45) / 2`. Header drawn at `(start_x, start_y)`, footer at bottom; body uses same `start_x` and wraps to `width`. On resize, `update_dimensions` recomputes `start_x`, `start_y`, and `width = min(term_width, 60)`.

**Implication for shop**: (1) Loading view: one centered “box” (e.g. 40×20 or similar), random-character rectangle with inner line “Loading shop.gyeongho.dev” or “LOADING STORE” style. (2) Main view: full view (header + body + footer) treated as one block and centered with `lipgloss.Place(m.Width, m.Height, Center, Center, …)`; internal content width capped (e.g. 60) so the block does not over-stretch on large terminals.

**Alternatives considered**: Deviating from terminal.shop layout was rejected per user requirement.

---

## 3. Base64-Encoded Random String Rectangle

**Decision**: Implement the “Base64로 인코딩된 무작위 문자열 사각형 문양” as a centered rectangle of random characters drawn from a Base64-like charset (A–Z, a–z, 0–9, optionally +, /) so it is visually and semantically close to “Base64 random string” while matching terminal.shop’s loader style (random chars in a box).

**Rationale**: terminal.shop uses plain alphanumeric random chars; user asked for “Base64 인코딩된 무작위 문자열”. Using the standard Base64 character set (or strict Base64 encoding of random bytes) satisfies “Base64” and “random string”; the rectangle is the same idea as terminal.shop’s loader box.

**Findings**:

- **Charset**: Standard Base64 = `A–Z`, `a–z`, `0–9`, `+`, `/`. For a “random string square”, either (a) generate N random bytes and Base64-encode to get a string, then display in a rectangular grid, or (b) pick random characters from the Base64 set and fill the grid. Both yield a Base64-looking rectangle; (b) is simpler and avoids line-break handling in Base64.
- **Placement**: Rectangle centered in terminal: same as terminal.shop loader—compute box width/height (e.g. 40×20), then `start_x = (term_width - boxWidth) / 2`, `start_y = (term_height - boxHeight) / 2`. In Bubble Tea we have no direct cursor positioning; we build a string that, when rendered, shows the rectangle centered. So build a string of (boxWidth * boxHeight) chars (or lines of boxWidth), then center that block with `lipgloss.Place(termWidth, termHeight, Center, Center, rectangleString)`.
- **Inner text**: Like terminal.shop, reserve an inner region (e.g. 20×3) and show “Loading shop.gyeongho.dev” or “LOADING STORE” centered on the middle line; rest of inner area can be spaces or subtle styling.

**Alternatives considered**: Plain alphanumeric (no +, /) was considered; including full Base64 set is a small change and better matches “Base64 인코딩된” wording.

---

## 4. Loading Phase in Bubble Tea

**Decision**: Model has a `Loading` boolean (or equivalent). When `Loading` is true, `View()` returns only the loading view (centered Base64-style random rectangle). When `Loading` is false, `View()` returns the main view (header + body + footer) centered. Transition from loading to main via `tea.Cmd`: e.g. `tea.Tick` after ~2 seconds or after initial data load, sending a message that sets `Loading = false`.

**Rationale**: terminal.shop shows the loader first, then clears and draws the main UI. In Bubble Tea we cannot clear the screen imperatively; we switch the rendered view. So “first screen” = first phase of the program where only the loading view is rendered; then one message switches to the main view.

**Findings**:

- **Init()**: Return a command that either (1) waits for a tick (e.g. 2s) then sends “load complete”, or (2) starts async load and sends “load complete” when done. For parity with terminal.shop (fixed 2s loader), a simple delay is enough; if we later want “load until API ready”, we keep the same message and set `Loading = false`.
- **Update**: On “load complete” message, set `m.Loading = false` and return the model; no need to re-fetch.
- **View**: If `m.Loading` then return centered loading rectangle (with optional refresh of random chars via `tea.Tick` for animation); else return centered main view. This keeps a single “first screen” as the loading view.

**Alternatives considered**: Showing main view with a loading overlay is possible but terminal.shop shows a full-screen loader first; we match that for consistency.
