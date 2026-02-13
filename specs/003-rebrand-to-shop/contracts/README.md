# Contracts: 003-rebrand-to-shop

This feature does **not** change API contracts. The shop TUI continues to use the same API surface (products, about, FAQ, cart, orders) as defined in [001-ssh-shop-tui/contracts](../001-ssh-shop-tui/contracts/).

Only naming and content change:

- **Base URL**: Client default and documentation should reference `shop.gyeongho.dev` (e.g. `https://shop.gyeongho.dev/api`) (see 001-ssh-shop-tui contracts).
- **Response content**: Store info (about) and any store-level copy returned by the API should reflect shop branding and the expanded product range; contract schema (e.g. StoreInfo fields) is unchanged.

For OpenAPI or other contract artifacts, see [001-ssh-shop-tui/contracts](../001-ssh-shop-tui/contracts/).
