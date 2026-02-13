# Data Model: SSH Access to Terminal Shop

**Feature**: 004-ssh-shop-access  
**Source**: [spec.md](./spec.md); no new domain entities.

---

## Scope

This feature does not introduce new domain entities (Product, Cart, Order, StoreInfo, FAQEntry). Those remain as defined in [001-ssh-shop-tui/data-model.md](../001-ssh-shop-tui/data-model.md). The TUI state (CurrentPage, Products, Cart, About, FAQ, ScrollOffset, Cursor, Width, Height, Loading, Error) is unchanged and is created per SSH session.

---

## Runtime Concepts (Access Channel)

### SSH session

| Concept        | Description |
|----------------|-------------|
| Session        | One SSH connection from a client. Created when the user runs `ssh shop.gyeongho.dev` and accepted by the Wish server. |
| Session scope  | One `tea.Program` (one `tui.Model` instance) per session. State is not shared between sessions. |
| Session end    | When the client disconnects or the user quits the TUI (e.g. q), the program exits and the session ends. No persistent server-side session store for this feature. |

### Server lifecycle (no persisted data)

| Concept       | Description |
|---------------|-------------|
| Wish server   | Single process listening on the configured address (e.g. :22 or :2222). Accepts connections and spawns the bubbletea middleware handler per connection. |
| Host key      | Used by the server to identify itself; path configured via Wish options. Not a domain entity; operational configuration. |

---

## State Transitions

- **Client connects** → Wish accepts connection → middleware runs → `teaHandler(ssh.Session)` builds `tui.Model` and returns it → `tea.Program` runs until quit or disconnect.
- **Client disconnects or TUI quit** → Program exits → session closed; no server-side state retained for that session.
- **Connection failure (before TUI)** → SSH handshake or PTY allocation fails → user sees SSH client error (or server message); no TUI state created.

---

## Validation Rules

- One TUI model per SSH session; model fields follow existing rules (see 001-ssh-shop-tui data-model).
- No new validation on domain entities; API client (mock or HTTP) is unchanged.
