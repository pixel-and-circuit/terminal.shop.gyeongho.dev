# Data Model: SSH-Based Mushroom Sales TUI

**Feature**: 001-ssh-mushroom-tui  
**Source**: [spec.md](./spec.md) Key Entities + terminal.shop clone

---

## Domain Entities

### Product (Mushroom)

Represents a sellable mushroom product.

| Field         | Type     | Description |
|---------------|----------|-------------|
| ID            | string   | Unique identifier (e.g. from API). |
| Name          | string   | Display name. |
| Attributes    | []string | Tags (e.g. "whole", "dried"). |
| Price         | decimal  | Unit price (e.g. for display "₩X" or "$X"). |
| Description   | string   | Short description. |
| Quantity      | int      | Stock or "bags left"; 0 = sold out. |

**Validation**: Name non-empty; Price ≥ 0; Quantity ≥ 0.

**Source**: API `GET /api/products` (or static/mock until server exists).

---

### CartItem

One line in the shopping cart: product + quantity.

| Field       | Type   | Description |
|-------------|--------|-------------|
| ProductID   | string | References Product.ID. |
| Name        | string | Snapshot of product name for display. |
| UnitPrice   | decimal| Snapshot of price. |
| Quantity    | int    | Number of units (≥ 1). |

**Derived**: LineTotal = UnitPrice × Quantity.

---

### Cart

Current shopping cart state.

| Field     | Type        | Description |
|-----------|-------------|-------------|
| Items     | []CartItem  | Ordered list of items. |
| UpdatedAt | timestamp   | Optional; for display/cache. |

**Derived**: Total = sum(LineTotal); ItemCount = sum(Quantity).  
**Validation**: No duplicate ProductID in Items (merge quantities instead).

**Source**: Local TUI state; optionally sync with API `GET/POST /api/cart` when backend exists.

---

### Order

A submitted purchase (checkout result).

| Field       | Type        | Description |
|-------------|-------------|-------------|
| ID          | string      | Order identifier from server. |
| Items       | []CartItem  | Snapshot at checkout. |
| Total       | decimal     | Order total. |
| Status      | string      | e.g. "confirmed", "pending". |
| CreatedAt   | timestamp   | Optional. |

**Source**: API `POST /api/orders` (or mock response).

---

### StoreInfo (About)

Static content for the About page.

| Field   | Type   | Description |
|---------|--------|-------------|
| Title   | string | e.g. "About the store". |
| Body    | string | Markdown or plain text (farm story, SSH store info, contact). |

**Source**: API `GET /api/about` or static config/file; mock in tests.

---

### FAQEntry

One Q&A pair for the FAQ page.

| Field    | Type   | Description |
|----------|--------|-------------|
| Question | string | Question text. |
| Answer   | string | Answer text (may be multi-line). |

**Source**: API `GET /api/faq` or static config; mock in tests.

---

## TUI State (Application Model)

Not persisted; drives Bubble Tea Model/View.

| Field          | Type        | Description |
|----------------|-------------|-------------|
| CurrentPage    | enum        | Landing \| Shop \| About \| FAQ \| Cart. |
| Products       | []Product   | Cached product list. |
| Cart           | Cart        | Current cart. |
| About          | StoreInfo   | About content. |
| FAQ            | []FAQEntry  | FAQ content. |
| ScrollOffset   | int         | Vertical scroll for current page (e.g. FAQ, Shop). |
| Cursor         | int         | Selection index (e.g. product list or menu). |
| Width, Height  | int         | Terminal dimensions (from WindowSizeMsg). |
| Loading        | bool        | Show loader until first data loaded. |
| Error          | string      | Optional error message to display. |

**State transitions**:

- Key **a** → CurrentPage = Shop, ScrollOffset = 0.
- Key **s** → CurrentPage = About, ScrollOffset = 0.
- Key **d** → CurrentPage = FAQ, ScrollOffset = 0.
- Key **c** (on Shop or from footer) → CurrentPage = Cart.
- Key **Up/Down** → adjust ScrollOffset or Cursor within bounds.
- Add to cart → Cart.Items updated; stay on Shop or switch to Cart per UX.
- Submit order → Cart cleared; show confirmation; optional Order stored for display.

---

## Relationships

- **Product** is standalone; referenced by **CartItem** (ProductID, name/price snapshot).
- **Cart** aggregates **CartItem**; may be synced with API later.
- **Order** is a snapshot of **Cart** at checkout.
- **StoreInfo** and **FAQEntry** are read-only content for About and FAQ pages.
