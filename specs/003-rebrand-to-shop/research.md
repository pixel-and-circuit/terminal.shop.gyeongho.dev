# Research: Rebrand to Shop

**Feature**: 003-rebrand-to-shop  
**Purpose**: Resolve approach for Go module rename, UI strings, and test/fixture impact.

---

## 1. Go module and package rename (previous module → shop)

**Decision**: Rename module in `go.mod` from `shop.gyeongho.dev` to `shop.gyeongho.dev`. Rename directory `cmd/shop` to `cmd/shop`. Update Makefile to build `bin/shop` from `./cmd/shop`. Update every Go import path from `shop.gyeongho.dev/...` to `shop.gyeongho.dev/...`.

**Rationale**: User stated the GitHub repository path will change later and the Go package name must change to shop. Standard Go practice is one module per repo; the module path should reflect the future repo (e.g. shop.gyeongho.dev). Renaming the cmd directory and binary to `shop` keeps CLI and binary name aligned with the product name.

**Alternatives considered**:
- Keep module as `shop.gyeongho.dev` and only change user-facing strings: rejected because user explicitly required package name change for future repo path.
- Use a separate module path (e.g. `github.com/GyeongHoKim/shop`): deferred to when the repo is moved; for now `shop.gyeongho.dev` is consistent with the target domain.

**Files to touch**:
- `go.mod`: module line
- All `.go` files: import paths (cmd/shop, internal/*, tests/*)
- Makefile: `build` and `run` targets (cmd/shop, bin/shop)
- README, AGENTS.md: commands and paths (cmd/shop, bin/shop)

---

## 2. User-facing and internal strings (UI and copy)

**Decision**: (1) Loader: change `internal/tui/loader.go` constant from `"Loading shop.gyeongho.dev"` to `"Loading shop.gyeongho.dev"`. (2) About/store info: in `internal/apiclient/mock.go` (and any fixtures), replace previous store title and the about body with shop branding and copy that states gyeongho sells all products (produce, embedded devices, robots, etc.). (3) HTTP client default base URL in `internal/apiclient/http.go`: change default from `https://shop.gyeongho.dev/api` to `https://shop.gyeongho.dev/api`. (4) Model comment in `internal/model/product.go`: change "sellable product" to "sellable product" (catalog is no longer single-category only).

**Rationale**: Spec FR-002 and FR-003 require all user-facing store name and domain to be "shop" / shop.gyeongho.dev and store content to reflect multiple product types. Loader and about are the main UI surfaces; default API URL is an internal identifier (FR-004).

**Alternatives considered**:
- Keep loader as "Loading shop.gyeongho.dev" for nostalgia: rejected; spec requires no previous store name as store name in shipped experience.
- Configurable loader text via env: not required for this feature; constant change is sufficient.

**Impact on UI**: Header/footer in `internal/tui/header.go` use "terminal", "a shop", "s about", "d faq", "cart" – no literal previous store name; no change. Only loader text and about/FAQ content (from API/mock) change.

---

## 3. Tests and fixtures: comprehensive impact

**Decision**: (1) **Imports**: Every test file under `tests/` that imports `shop.gyeongho.dev/...` must be updated to `shop.gyeongho.dev/...`. (2) **Loader assertion**: In `tests/unit/tui_view_test.go`, change assertion from `"Loading shop.gyeongho.dev"` to `"Loading shop.gyeongho.dev"`. (3) **Product names in tests**: Tests that assert on product names (e.g. a catalog product name) are asserting on mock product data; product names can remain as-is (they are catalog items). (4) **Store/about assertions**: Where tests assert on about title or body (e.g. previous store title, "shop.gyeongho.dev"), update expected strings to the new shop title and shop.gyeongho.dev. (5) **Fixtures**: `tests/e2e/fixtures/about.json` – update title and body to shop branding and "gyeongho sells all products" style copy. `tests/e2e/fixtures/products.json` (and cart/order fixtures that reference product names) – product names like "a catalog product name" are valid catalog items; only change if we also change mock product data (optional; spec allows keeping mushroom as one category). For consistency with mock default about, fixture about.json should match shop branding.

**Rationale**: Tests must pass after rename; import paths must match the new module. Assertions that check user-visible strings (loader, about) must expect the new copy. Product-name assertions are about mock data; keeping existing catalog product names is acceptable as one product category.

**Alternatives considered**:
- Leave fixture product names as single-category only: acceptable; spec says "product data may already exist or be added separately". Mock and fixtures can keep single-category products; about/store description must reflect "all products".
- Change all product names in fixtures to other categories: not required for rebrand; can be done later.

**Files to touch**:
- `tests/unit/*.go`: imports + tui_view_test.go loader assertion; shop_test.go and others if they assert on about/store strings.
- `tests/integration/*.go`: imports + any assertions on loader or about content.
- `tests/e2e/e2e_test.go`: imports; product names in fixtures can stay.
- `tests/e2e/fixtures/about.json`: title and body to shop + "all products" copy.
- Other fixtures: only about.json content change required for this feature.

---

## 4. README and AGENTS.md

**Decision**: README: Change title to "shop.gyeongho.dev" (or similar). Replace tagline with description that gyeongho sells all products (produce, embedded devices, robots, etc.). Update commands to `./bin/shop` and `go run ./cmd/shop`; SSH example to `shop.gyeongho.dev`. AGENTS.md: Update project purpose to "SSH-accessible TUI for shop at shop.gyeongho.dev" and "all products gyeongho provides"; update paths to `cmd/shop`, `bin/shop`, and link to specs/003-rebrand-to-shop where relevant.

**Rationale**: User requirement: "README 설명 또한, gyeongho가 제공하는 모든 상품을 판매한다라는 설명으로 바꿔야 한다". FR-004 requires documentation to use new naming and domain.

**Alternatives considered**: None; direct requirement.
