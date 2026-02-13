# Tasks: Rebrand from Mushroom to Shop

**Input**: Design documents from `specs/003-mushroom-to-shop/`  
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/, quickstart.md

**Tests**: This feature updates existing tests (import paths and assertions); no new test tasks. Spec does not request TDD or new tests.

**Organization**: Tasks are grouped by user story so each story can be implemented and verified independently.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: User story (US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

Single Go module at repo root: `cmd/shop/`, `internal/`, `tests/` (see plan.md).

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Prepare repository for module and path rename.

- [x] T001 Verify feature branch and design docs: ensure on branch `003-mushroom-to-shop` and `specs/003-mushroom-to-shop/` contains plan.md, spec.md, research.md
- [x] T002 Rename `cmd/mushroom` to `cmd/shop`: create `cmd/shop/`, move `cmd/mushroom/main.go` to `cmd/shop/main.go`, remove empty `cmd/mushroom/`
- [x] T003 Update `go.mod`: change module line from `mushroom.gyeongho.dev` to `shop.gyeongho.dev`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Module and import path changes so the project builds with the new identity. No user story work can begin until this phase is complete.

- [x] T004 [P] Update import paths in `cmd/shop/main.go` from `mushroom.gyeongho.dev` to `shop.gyeongho.dev`
- [x] T005 [P] Update import paths in `internal/apiclient/client.go`, `internal/apiclient/mock.go`, `internal/apiclient/http.go` from `mushroom.gyeongho.dev` to `shop.gyeongho.dev`
- [x] T006 [P] Update import paths in `internal/tui/app.go`, `internal/tui/loader.go`, and `internal/tui/pages/*.go` from `mushroom.gyeongho.dev` to `shop.gyeongho.dev`
- [x] T007 Update `Makefile`: build target output `bin/shop` from `./cmd/shop`; run target `go run ./cmd/shop`
- [x] T008 Run `make format` and `make build`; fix any remaining import or path errors until build passes

**Checkpoint**: Foundation ready — binary builds as `bin/shop`; user story implementation can begin.

---

## Phase 3: User Story 1 – Access shop via new domain (Priority: P1) – MVP

**Goal**: User sees shop branding when opening the app (loader and domain identity show shop.gyeongho.dev).

**Independent Test**: Run `./bin/shop` and confirm loader shows "Loading shop.gyeongho.dev" and no "mushroom" as store name on first screen.

- [x] T009 [US1] Change loader constant in `internal/tui/loader.go` from `"Loading mushroom.gyeongho.dev"` to `"Loading shop.gyeongho.dev"`
- [x] T010 [US1] Change default base URL in `internal/apiclient/http.go` from `https://mushroom.gyeongho.dev/api` to `https://shop.gyeongho.dev/api`

**Checkpoint**: User Story 1 deliverable — loader and API default use shop.gyeongho.dev.

---

## Phase 4: User Story 2 – Consistent shop naming across the experience (Priority: P2)

**Goal**: No "mushroom" store name in about, copy, or comments; store referred to as shop / shop.gyeongho.dev everywhere.

**Independent Test**: Walk every screen (About, FAQ, Shop, Cart); verify store title and copy use "shop" or shop.gyeongho.dev, not "mushroom".

- [x] T011 [US2] Replace About title and body in `internal/apiclient/mock.go` (GetAbout) with shop branding and shop.gyeongho.dev; body must state that the shop sells all products gyeongho provides (mushrooms, embedded devices, robots, etc.)
- [x] T012 [US2] Update Product comment in `internal/model/product.go` from "sellable mushroom product" to "sellable product"
- [x] T013 [US2] Update comments in `internal/apiclient/mock.go` (e.g. "Mushroom Department Store", "mushroom.gyeongho.dev") to shop/store wording and shop.gyeongho.dev

**Checkpoint**: User Story 2 deliverable — about and internal copy use shop identity.

---

## Phase 5: User Story 3 – Correct product and store context in content (Priority: P3)

**Goal**: About and store-level content describe the shop as selling a variety of products (mushrooms, embedded devices, robots, etc.).

**Independent Test**: Read About and FAQ; confirm copy reflects multiple product types and does not imply mushrooms-only.

- [x] T014 [US3] Update `tests/e2e/fixtures/about.json`: set title and body to shop branding and description that gyeongho sells all products (mushrooms, embedded devices, robots, etc.)
- [x] T015 [US3] Review FAQ entries in `internal/apiclient/mock.go` (GetFAQ); if any copy implies mushrooms-only, add or adjust wording to reference multiple product types where appropriate

**Checkpoint**: User Story 3 deliverable — store and fixture content reflect product diversity.

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Update tests, documentation, and run quality gate so the rebrand is complete and all checks pass.

- [x] T016 [P] Update `tests/unit/tui_view_test.go`: import path to `shop.gyeongho.dev`; change loader assertion from "Loading mushroom.gyeongho.dev" to "Loading shop.gyeongho.dev"
- [x] T017 [P] Update import paths in `tests/unit/shop_test.go`, `tests/unit/faq_test.go`, `tests/unit/cart_test.go`, `tests/unit/cart_tui_test.go`, `tests/unit/about_test.go`, `tests/unit/tui_navigation_test.go` from `mushroom.gyeongho.dev` to `shop.gyeongho.dev`; update any assertions on about title/body to expect shop branding
- [x] T018 [P] Update import paths in `tests/integration/shop_test.go`, `tests/integration/faq_test.go`, `tests/integration/cart_test.go`, `tests/integration/about_test.go`, `tests/integration/navigation_test.go` to `shop.gyeongho.dev`; update any assertions on about content or loader to shop branding
- [x] T019 [P] Update import paths in `tests/e2e/e2e_test.go` from `mushroom.gyeongho.dev` to `shop.gyeongho.dev`
- [x] T020 [P] Update `README.md`: title to shop.gyeongho.dev; tagline/description to state that gyeongho sells all products (mushrooms, embedded devices, robots, etc.); commands `./bin/shop` and `go run ./cmd/shop`; SSH example `shop.gyeongho.dev`
- [x] T021 [P] Update `AGENTS.md`: project purpose to SSH-accessible TUI for shop at shop.gyeongho.dev and all products gyeongho provides; paths `cmd/shop`, `bin/shop`; reference specs/003-mushroom-to-shop where relevant
- [x] T022 Run `make format`, `make build`, and `make test`; fix any remaining failures until all pass

---

## Dependencies & Execution Order

### Phase Dependencies

- **Phase 1 (Setup)**: No dependencies — start immediately.
- **Phase 2 (Foundational)**: Depends on Phase 1 — blocks all user stories.
- **Phase 3 (US1)**: Depends on Phase 2 — can start when build passes.
- **Phase 4 (US2)**: Depends on Phase 2 — can run after or in parallel with Phase 3 (no code conflict; different files).
- **Phase 5 (US3)**: Depends on Phase 2; fixture/about content aligns with US2 — can run after Phase 4 or in parallel with Phase 3/4 where files differ.
- **Phase 6 (Polish)**: Depends on Phases 3–5 for correct assertions and docs — run after user story tasks.

### User Story Dependencies

- **US1 (P1)**: No dependency on US2/US3 — loader and HTTP default only.
- **US2 (P2)**: No dependency on US1 for implementation (different files); test assertions in Polish phase expect US1+US2.
- **US3 (P3)**: About body and FAQ can be done with US2; fixture about.json and FAQ review are US3-specific.

### Within Each Phase

- Phase 2: T004–T006 can run in parallel; T007 then T008.
- Phase 3: T009 and T010 can run in parallel.
- Phase 4: T011–T013 can run in parallel after Phase 2.
- Phase 5: T014 and T015 can run in parallel.
- Phase 6: T016–T021 can run in parallel; T022 last.

### Parallel Opportunities

- Phase 2: T004, T005, T006 [P] in parallel.
- Phase 3: T009, T010 [P] in parallel.
- Phase 4: T011, T012, T013 [P] in parallel.
- Phase 5: T014, T015 [P] in parallel.
- Phase 6: T016–T021 [P] in parallel; T022 after.

---

## Parallel Example: User Story 1

```bash
# After Phase 2 completes, run both US1 tasks:
Task T009: "Change loader constant in internal/tui/loader.go to Loading shop.gyeongho.dev"
Task T010: "Change default base URL in internal/apiclient/http.go to https://shop.gyeongho.dev/api"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup (T001–T003).
2. Complete Phase 2: Foundational (T004–T008) — build must pass.
3. Complete Phase 3: User Story 1 (T009–T010).
4. **Validate**: Run `./bin/shop` and confirm loader shows "Loading shop.gyeongho.dev".
5. Optionally run Phase 6 test/doc updates (T016–T022) so CI and docs match.

### Incremental Delivery

1. Setup + Foundational → build as `bin/shop`.
2. Add US1 → loader and API default use shop.gyeongho.dev (MVP).
3. Add US2 → about and copy use shop identity.
4. Add US3 → about/FAQ and fixtures reflect all products.
5. Polish → tests and README/AGENTS updated; `make test` passes.

### Parallel Team Strategy

- One developer: execute phases in order; use [P] within a phase where possible.
- Multiple developers: after Phase 2, one can do US1 (loader + http), another US2 (mock about + model comment), another US3 (fixtures + FAQ); then one does Phase 6 (tests + docs + gate).

---

## Notes

- [P] tasks use different files and have no ordering dependency within the same phase.
- [USn] maps the task to the user story for traceability.
- Each user story phase is independently testable per the spec’s Independent Test.
- **Quality gate**: Run `make format` and `make build` after code changes (constitution); Phase 6 includes `make test`.
- Product names in mock/fixtures (e.g. "Oyster Mushroom") may stay as catalog data; only store name, domain, and about/FAQ copy change to shop and all-products messaging.
