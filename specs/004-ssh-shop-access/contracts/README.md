# Contracts: SSH Access to Terminal Shop

**Feature**: 004-ssh-shop-access

## Access contract

This feature does not introduce HTTP/REST or GraphQL APIs. The only “contract” with the user is **SSH access**:

- **Endpoint**: `ssh shop.gyeongho.dev` (hostname resolves; SSH server listens on the configured port, typically 22).
- **Protocol**: SSH (e.g. SSH-2). Client must support PTY (pseudo-terminal) for the TUI.
- **Behavior**: On successful connection, the server runs the terminal shop TUI in the session. User interacts with the same flows as the local TUI (Shop, About, FAQ, Cart, keys a/s/d/c, scroll, quit).
- **Failure**: If the connection cannot be established, the user sees SSH client or server errors (e.g. connection refused, timeout); no application-level API response.

No OpenAPI or GraphQL schema is defined for this feature. Backend API contracts (if any) for products, cart, orders remain as defined elsewhere (e.g. 001-ssh-shop-tui or backend specs).
