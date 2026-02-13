# Data Model: Rebrand to Shop

**Feature**: 003-rebrand-to-shop  
**Source**: [spec.md](./spec.md) Key Entities; no new domain entities.

---

## Scope of change

This feature does **not** introduce new domain entities or change the structure of existing ones. The domain model (Product, Cart, Order, StoreInfo, FAQEntry) remains as defined in [001-ssh-shop-tui/data-model.md](../001-ssh-shop-tui/data-model.md).

Changes are limited to:

- **Naming**: Module and package paths module and binary to `shop.gyeongho.dev` and `cmd/shop`.
- **Content**: Store identity and copy (StoreInfo title/body, loader text) updated to "shop" and shop.gyeongho.dev, with store description reflecting that the shop sells all products gyeongho provides (produce, embedded devices, robots, etc.).

---

## Entities (unchanged structure)

| Entity      | Change |
|------------|--------|
| **Product** | No structural change. Comment in code updated from "sellable product" in comment. Product names in mock/fixtures may remain as one category as one category. |
| **CartItem** | No change. |
| **Cart**     | No change. |
| **Order**    | No change. |
| **StoreInfo**| No field change. **Content** (Title, Body) in mock and fixtures updated to shop branding and "all products" description. |
| **FAQEntry** | No structural change. FAQ copy may reference multiple product types where appropriate. |

---

## TUI / presentation

- **Loader**: Constant text changed from previous loader text to `"Loading shop.gyeongho.dev"` in `internal/tui/loader.go`. No new state or entity.
- **Header/Footer**: No structural change; already use "shop", "about", "faq", "cart". No previous store name in header.

---

## Validation rules

Existing validation rules for Product, Cart, Order, StoreInfo, and FAQEntry remain unchanged. No new validation rules added.
