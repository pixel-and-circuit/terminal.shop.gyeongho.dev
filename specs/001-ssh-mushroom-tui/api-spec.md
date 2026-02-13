# Shop TUI API Specification

## Overview

This API backs the SSH-accessible TUI at shop.gyeongho.dev. It provides products, store info, FAQ, cart, and order endpoints. The client is configured with a base URL and may send an optional user identity header for scoped data (cart, orders).

## Base URL

- **Configurable**: Set `SHOP_API_BASE` (e.g. `https://shop.gyeongho.dev/api` or `http://localhost:8080` for tests).
- **Default**: `https://shop.gyeongho.dev/api`
- Clients use `NewHTTPClient(baseURL)`; empty base URL causes the client to use the env default so tests can point to a local mock server.

## User authentication (사용자 인증)

- **Identity**: The user is identified by the **SSH public key fingerprint**. Same key implies the same user; no key implies anonymous. This aligns with a terminal.shop-style strategy.
- **Flow**: The SSH gateway authenticates the user (e.g. by public key), computes the key fingerprint, and injects it into the TUI process (e.g. env `SHOP_USER_ID`). The HTTP client sends this value on every request (header `X-User-Id`). The backend uses it to scope cart and orders to that user.
- **Header**: `X-User-Id: <fingerprint>` (recommended). Alternative: `Authorization: Bearer <opaque-id>` if the gateway issues a token instead of the raw fingerprint.
- **Optional**: Requests without `X-User-Id` are treated as anonymous (e.g. read-only or ephemeral cart). The mock server and E2E tests may omit the header or send it to cover both paths.

## Endpoints

| Method | Path      | Description                    | Response (success)     |
|--------|-----------|--------------------------------|------------------------|
| GET    | /products | List products                  | 200, array of Product  |
| GET    | /about    | Store information (About)     | 200, StoreInfo         |
| GET    | /faq      | FAQ entries                    | 200, array of FAQEntry  |
| GET    | /cart     | Current cart                   | 200, Cart               |
| POST   | /cart     | Add item to cart (or create)   | 200/201, Cart           |
| POST   | /orders   | Submit order (checkout)        | 201, Order; 400 if empty cart |

### POST /cart

- **Request body**: `{ "productId": string, "quantity": integer }` (required).
- **Response**: Updated Cart (JSON).

### POST /orders

- **Request body**: Cart (JSON).
- **Response**: 201 with Order (JSON), or 400 for invalid/empty cart.

## Request / response schemas

JSON field names and shapes follow `internal/model` and `specs/001-ssh-mushroom-tui/contracts/openapi.yaml`:

- **Product**: `id`, `name`, `attributes` (array), `price`, `description`, `quantity`
- **StoreInfo**: `id`, `title`, `body`
- **FAQEntry**: `id`, `question`, `answer`
- **Cart**: `items` (array of CartItem), `updatedAt` (optional)
- **CartItem**: `productId`, `name`, `unitPrice`, `quantity`
- **Order**: `id`, `items`, `total`, `status`, `createdAt` (optional)

No duplicate schema definitions here; see OpenAPI or the Go model types for full details.
