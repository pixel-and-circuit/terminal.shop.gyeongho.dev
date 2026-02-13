# Contracts: 003-mushroom-to-shop

This feature does **not** change API contracts. The shop TUI continues to use the same API surface (products, about, FAQ, cart, orders) as defined in [001-ssh-mushroom-tui/contracts](../001-ssh-mushroom-tui/contracts/).

Only naming and content change:

- **Base URL**: Client default and documentation should reference `shop.gyeongho.dev` (e.g. `https://shop.gyeongho.dev/api`) instead of `mushroom.gyeongho.dev`.
- **Response content**: Store info (about) and any store-level copy returned by the API should reflect shop branding and the expanded product range; contract schema (e.g. StoreInfo fields) is unchanged.

For OpenAPI or other contract artifacts, see [001-ssh-mushroom-tui/contracts](../001-ssh-mushroom-tui/contracts/).
