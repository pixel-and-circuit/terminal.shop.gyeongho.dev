# Tasks: SSH-Based Mushroom Sales TUI

**Input**: Design documents from `/specs/001-ssh-mushroom-tui/`  
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/

**Tests**: TDD requested in feature—tests are written first per user story, then implementation.

**Organization**: Tasks are grouped by user story for independent implementation and testing.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story (US1, US2, US3, US4)
- Include exact file paths in descriptions

## Path Conventions

- **Single project** (plan.md): `cmd/`, `internal/`, `tests/` at repository root
- Paths use: `cmd/mushroom/`, `internal/model/`, `internal/apiclient/`, `internal/tui/`, `internal/tui/pages/`, `tests/unit/`, `tests/integration/`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 Create project structure per plan.md: cmd/mushroom/, internal/model/, internal/apiclient/, internal/tui/, internal/tui/pages/, tests/unit/, tests/integration/
- [x] T002 Initialize Go module and add Charm dependencies (github.com/charmbracelet/bubbletea, lipgloss, bubbles) in go.mod at repo root
- [x] T003 [P] Add Makefile with targets format (gofmt/goimports), build (go build -o bin/mushroom ./cmd/mushroom), test (go test ./...) at repo root
- [x] T004 [P] Add .pre-commit-config.yaml for format and commitlint hooks at repo root
- [x] T005 [P] Add GitHub Actions workflow for CI (format check, make build, make test) in .github/workflows/ci.yml

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T006 [P] Create Product, CartItem, Cart, Order, StoreInfo, FAQEntry structs in internal/model/ (product.go, cart.go, order.go, storeinfo.go, faq.go per data-model.md)
- [x] T007 Define API client interface (GetProducts, GetAbout, GetFAQ, GetCart, AddToCart, SubmitOrder) in internal/apiclient/client.go
- [x] T008 Implement mock API client satisfying interface with static data in internal/apiclient/mock.go
- [x] T009 Create root Bubble Tea model (CurrentPage enum, width/height, scroll, products/cart/about/faq state) and Init/Update/View stubs in internal/tui/app.go
- [x] T010 Create main entry that runs tea.NewProgram with root model and injects API client in cmd/mushroom/main.go
- [x] T011 Create key binding constants (keys a, s, d, c, up, down, q, ctrl+c) in internal/tui/keys.go

**Checkpoint**: Foundation ready—binary runs with minimal View; API client injectable; user story implementation can begin

---

## Phase 3: User Story 1 - Access TUI via SSH and Navigate (Priority: P1) 🎯 MVP

**Goal**: User sees TUI with header/footer and can switch pages with a, s, d (Shop, About, FAQ).

**Independent Test**: Run binary (or SSH), confirm header shows "a shop", "s about", "d faq"; press a/s/d and confirm content area switches; press q to quit.

### Tests for User Story 1 (TDD: write first, then implement)

- [x] T012 [P] [US1] Write unit tests for Update: key "a" sets CurrentPage to Shop, "s" to About, "d" to FAQ in tests/unit/tui_navigation_test.go
- [x] T013 [P] [US1] Write unit tests for View: header contains shop/about/faq labels and shortcuts in tests/unit/tui_view_test.go
- [x] T014 [US1] Write integration test: start program, send key "a", assert current page is Shop in tests/integration/navigation_test.go

### Implementation for User Story 1

- [x] T015 [US1] Implement header component (logo, a shop, s about, d faq, cart placeholder) in internal/tui/header.go
- [x] T016 [US1] Implement footer component (+/- qty, c cart, q quit) in internal/tui/footer.go
- [x] T017 [US1] Implement landing page content (welcome, navigate with a/s/d) in internal/tui/pages/landing.go
- [x] T018 [US1] Wire page switching in Update (a/s/d) and render header/footer/current page in View in internal/tui/app.go
- [x] T019 [US1] Handle WindowSizeMsg and resize state in internal/tui/app.go for terminal resize

**Checkpoint**: User Story 1 complete—navigation (a/s/d) and landing visible; independently testable

---

## Phase 4: User Story 2 - Shop: Browse and Order Mushrooms (Priority: P2)

**Goal**: User opens Shop (a), sees product list with scroll, selects product, adds to cart; can open cart (c) and complete checkout.

**Independent Test**: Press a, see product list; use Up/Down to scroll; select and add to cart; press c, see cart; submit order and see confirmation.

### Tests for User Story 2 (TDD)

- [x] T020 [P] [US2] Write unit tests for shop page: View shows product list; Update Up/Down changes scroll offset in tests/unit/shop_test.go
- [x] T021 [P] [US2] Write unit tests for cart: AddToCart updates Cart items; Total and ItemCount correct in tests/unit/cart_test.go
- [x] T022 [US2] Write integration test: open shop (a), add product to cart, open cart (c), assert cart content in tests/integration/shop_test.go

### Implementation for User Story 2

- [x] T023 [US2] Implement shop page (list products, scroll Up/Down, cursor selection) in internal/tui/pages/shop.go
- [x] T024 [US2] Implement add-to-cart action (Enter or key) and cart state update in internal/tui/app.go and internal/tui/pages/shop.go
- [x] T025 [US2] Implement cart page (list items, total, checkout action) in internal/tui/pages/cart.go
- [x] T026 [US2] Wire GetProducts from API client in Init or on Shop focus in internal/tui/app.go
- [x] T027 [US2] Implement checkout flow (submit order via client or local state) and show confirmation in internal/tui/pages/cart.go
- [x] T028 [US2] Handle empty cart checkout: show message and option to return to shop in internal/tui/pages/cart.go

**Checkpoint**: User Story 2 complete—shop, cart, checkout flow work independently

---

## Phase 5: User Story 3 - About: View Store Information (Priority: P3)

**Goal**: User presses s, sees About content (store/SSH store info); can return via a/s/d or back behavior.

**Independent Test**: Press s, verify about content visible; press a to return to Shop or landing.

### Tests for User Story 3 (TDD)

- [x] T029 [P] [US3] Write unit tests for about page: View shows StoreInfo title and body in tests/unit/about_test.go
- [x] T030 [US3] Write integration test: press "s", assert about content in View in tests/integration/about_test.go

### Implementation for User Story 3

- [x] T031 [US3] Implement about page (render StoreInfo title and body) in internal/tui/pages/about.go
- [x] T032 [US3] Wire GetAbout from API client when entering About page in internal/tui/app.go
- [x] T033 [US3] Ensure a/s/d from About switches to Shop/About/FAQ per key in internal/tui/app.go

**Checkpoint**: User Story 3 complete—About page reachable and readable

---

## Phase 6: User Story 4 - FAQ: View Frequently Asked Questions (Priority: P4)

**Goal**: User presses d, sees FAQ list; Up/Down scroll; can return via a/s/d.

**Independent Test**: Press d, see FAQ Q&A; scroll; press a to return.

### Tests for User Story 4 (TDD)

- [x] T034 [P] [US4] Write unit tests for FAQ page: View shows questions and answers; Update Up/Down changes scroll in tests/unit/faq_test.go
- [x] T035 [US4] Write integration test: press "d", scroll FAQ, assert content in tests/integration/faq_test.go

### Implementation for User Story 4

- [x] T036 [US4] Implement FAQ page (list FAQEntry, scroll Up/Down) in internal/tui/pages/faq.go
- [x] T037 [US4] Wire GetFAQ from API client when entering FAQ page in internal/tui/app.go
- [x] T038 [US4] Reset scroll offset when switching to FAQ and bound scroll in internal/tui/app.go or internal/tui/pages/faq.go

**Checkpoint**: User Story 4 complete—FAQ page reachable and scrollable

---

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [x] T039 [P] Add loading/splash screen (terminal.shop-style) before first content in internal/tui/loader.go
- [x] T040 [P] Add real HTTP client implementation calling mushroom.gyeongho.dev/api per contracts/openapi.yaml in internal/apiclient/http.go
- [x] T041 Handle network/API errors: show user-friendly message and retry or back in internal/tui/app.go
- [x] T042 Run quickstart.md validation: make format, make build, make test from repo root
- [x] T043 Update README.md with project layout, Makefile targets, quickstart link (specs/001-ssh-mushroom-tui/quickstart.md), and SSH run instructions

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies—can start immediately
- **Foundational (Phase 2)**: Depends on Setup—BLOCKS all user stories
- **User Stories (Phase 3–6)**: Depend on Foundational; US1 first (MVP), then US2, US3, US4 in order or in parallel if staffed
- **Polish (Phase 7)**: Depends on all desired user stories complete

### User Story Dependencies

- **US1 (P1)**: After Foundational only—no other story required
- **US2 (P2)**: After Foundational; uses same app.go and API client; independently testable
- **US3 (P3)**: After Foundational; independently testable
- **US4 (P4)**: After Foundational; independently testable

### Within Each User Story

- Tests MUST be written and FAIL before implementation (TDD)
- Then implement to pass tests; refactor with tests green
- Core implementation before integration tests

### Parallel Opportunities

- Phase 1: T003, T004, T005 [P]; T006 [P] in Phase 2
- Phase 3: T012, T013 [P]; then T014 after
- Phase 4: T020, T021 [P]; Phase 5: T029 [P]; Phase 6: T034 [P]
- Phase 7: T039, T040 [P]
- Different user stories (US2, US3, US4) can be done in parallel after US1

---

## Parallel Example: User Story 1

```bash
# Run unit tests for US1 in parallel (after writing them):
go test ./tests/unit/ -run Navigation
go test ./tests/unit/ -run View
# Then integration:
go test ./tests/integration/ -run Navigation
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup  
2. Complete Phase 2: Foundational  
3. Complete Phase 3: User Story 1 (tests first, then header, footer, landing, page switch)  
4. **STOP and VALIDATE**: Run binary, test a/s/d and q  
5. Deploy/demo TUI entry point if ready  

### Incremental Delivery

1. Setup + Foundational → foundation ready  
2. Add US1 → test navigation independently → MVP  
3. Add US2 → test shop/cart/checkout → full shop  
4. Add US3 → About; Add US4 → FAQ  
5. Polish (loader, real API client, README)  

### Parallel Team Strategy

- Complete Setup + Foundational together  
- Then: Dev A = US1, Dev B = US2, Dev C = US3/US4 (or sequence US1→US2→US3→US4)  
- Each story independently testable  

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to user story for traceability
- TDD: write failing tests first, then implementation
- **Quality gate**: Run `make format` and `make build` after code changes (constitution)
- Commit after each task or logical group
- Avoid: vague tasks, same-file conflicts, cross-story blocking dependencies
