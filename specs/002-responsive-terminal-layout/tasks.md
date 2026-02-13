# Tasks: Responsive Terminal Layout

**Input**: Design documents from `/specs/002-responsive-terminal-layout/`  
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, quickstart.md (contracts: no new API)

**Tests**: Plan requests extending unit/integration tests for View output and resize behavior; test tasks included below.

**Organization**: Tasks grouped by user story so each story can be implemented and tested independently.

## Format: `[ID] [P?] [Story?] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: User story (US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- Single Go module at repo root: `internal/tui/`, `cmd/shop/`, `tests/`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Verify project and dependencies are ready for layout changes

- [x] T001 Verify Go module and Charm Bubble Tea / Lip Gloss dependencies; confirm `internal/tui/` and `cmd/shop/` exist per plan in repo root

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Loading state, loading view, and centering must be in place before any user story can be verified. All stories depend on the app showing a centered loading screen then a centered main view.

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T002 Add `Loading` field to Model and handle load-complete message (e.g. `tea.Tick` then set `Loading = false`) in `internal/tui/app.go`
- [x] T003 [P] Implement terminal.shop-style loading view: Base64-style random character rectangle (e.g. 40×20) with inner centered line (e.g. "Loading shop.gyeongho.dev") in `internal/tui/loader.go`
- [x] T004 In `View()`, when `Loading` is true return output of loader centered with `lipgloss.Place(Width, Height, Center, Center, loadingView)` in `internal/tui/app.go`
- [x] T005 In `View()`, when `Loading` is false build main view (header + body + footer) and return it centered with `lipgloss.Place(Width, Height, Center, Center, mainView)` in `internal/tui/app.go`
- [x] T006 Apply max content width (e.g. 60) for main view body so the centered block does not over-stretch on very wide terminals in `internal/tui/app.go` and/or `internal/tui/pages/` as needed
- [x] T007 [P] Add or extend unit test for `View()` when `Loading` is true: loading view is centered and contains loader content (random-character rectangle, inner text) in `tests/unit/` (e.g. `tui_view_test.go` or equivalent)

**Checkpoint**: Foundation ready—loading screen and main view are centered; content width is capped; loading view is covered by test. User story implementation can proceed.

---

## Phase 3: User Story 1 - View Content in Any Terminal Size (Priority: P1) 🎯 MVP

**Goal**: Content and layout adapt to terminal width and height so the experience is readable and navigable without unnecessary horizontal scrolling at 80×24, 160×40, and other sizes.

**Independent Test**: Open the app in terminals of different dimensions (e.g. 80×24, 120×30, 200×50); confirm content adapts and remains usable without horizontal scroll.

### Implementation for User Story 1

- [x] T008 [US1] Ensure header and body content wrap and fit within capped width so 80×24 and 160×40 acceptance scenarios pass in `internal/tui/header.go` and `internal/tui/pages/*.go`
- [x] T009 [US1] Add or extend unit test for `View()` when `Loading` is false: main view is centered and contains header and footer in `tests/unit/` (e.g. `tui_view_test.go` or equivalent)

**Checkpoint**: User Story 1 is independently testable; content fits at common terminal sizes.

---

## Phase 4: User Story 2 - Layout Updates When Terminal Is Resized (Priority: P2)

**Goal**: When the user resizes the terminal, the layout updates to match the new dimensions in a timely way without refresh or reload.

**Independent Test**: Load the app, resize the terminal; verify layout and content reflow to the new size within a reasonable time (e.g. &lt; 2 seconds per SC-002).

### Implementation for User Story 2

- [x] T010 [US2] Ensure `tea.WindowSizeMsg` in `Update()` stores `Width` and `Height` and triggers re-render (no debounce that delays layout beyond 2s); ensure no layout loops or visual corruption on repeated or rapid resize in `internal/tui/app.go`
- [x] T011 [US2] Add or extend test for resize: send `tea.WindowSizeMsg` and assert `View()` output uses new dimensions (e.g. centered block reflects new size) in `tests/unit/` or `tests/integration/`

**Checkpoint**: User Stories 1 and 2 both work; resize updates layout in a timely manner.

---

## Phase 5: User Story 3 - Consistent Experience Across Terminal Types (Priority: P3)

**Goal**: Same terminal dimensions produce an equivalent optimized experience across different terminal emulators; no terminal-specific layout bugs.

**Independent Test**: Open the app in two or more terminal types at the same dimensions; confirm layout and readability are equivalent.

### Implementation for User Story 3

- [x] T012 [US3] Verify `View()` and layout logic use only model `Width` and `Height` (no hardcoded terminal or emulator checks) so behavior is consistent in `internal/tui/app.go` and `internal/tui/loader.go`

**Checkpoint**: All three user stories are independently functional; layout is dimension-driven only.

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Quality gate and validation

- [x] T013 Run `make format` and `make build`; fix any format or compile errors
- [x] T014 Validate run and resize steps from `specs/002-responsive-terminal-layout/quickstart.md` (manual or automated as appropriate); document or run SC-004 validation (common sizes and terminals matrix) per quickstart.md

---

## Dependencies & Execution Order

### Phase Dependencies

- **Phase 1 (Setup)**: No dependencies—can start immediately
- **Phase 2 (Foundational)**: Depends on Phase 1—blocks all user stories
- **Phase 3 (US1)**: Depends on Phase 2—MVP
- **Phase 4 (US2)**: Depends on Phase 2; can run after or in parallel with Phase 3
- **Phase 5 (US3)**: Depends on Phase 2; can run after Phase 3/4
- **Phase 6 (Polish)**: Depends on completion of desired user stories (T014 includes SC-004 validation via quickstart and common sizes/terminals matrix)

### User Story Dependencies

- **US1 (P1)**: After Foundational only; no dependency on US2/US3
- **US2 (P2)**: After Foundational only; independently testable
- **US3 (P3)**: After Foundational only; verification that layout uses only dimensions

### Within Each User Story

- Implementation tasks before or with test tasks as needed
- Story complete before moving to next priority when proceeding sequentially

### Parallel Opportunities

- T003 (loader implementation) can run in parallel with T002 (loading state) once T001 is done
- T007 (loading view test) can run in parallel with other Phase 2 tasks once T004 is done
- After Phase 2, US1 / US2 / US3 implementation can be split across developers
- T008 and T009 (US1) can be done in either order; T010 and T011 (US2) similarly

---

## Parallel Example: User Story 1

```text
# After Phase 2 complete:
T008: "Ensure header and body content wrap and fit within capped width in internal/tui/header.go and internal/tui/pages/*.go"
T009: "Add or extend unit test for View() when Loading false in tests/unit/"
```

---

## Parallel Example: Foundational

```text
# After T002 (Loading state) is done:
T003: "Implement terminal.shop-style loading view in internal/tui/loader.go"
# Then T004, T005, T006 in sequence (all touch app.go or depend on loader)
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup  
2. Complete Phase 2: Foundational (loading + centering + content width cap)  
3. Complete Phase 3: User Story 1 (wrap/fit, test)  
4. **STOP and VALIDATE**: Run app at 80×24 and 160×40; confirm no horizontal scroll  
5. Run Phase 6 (T013, T014) and ship if ready  

### Incremental Delivery

1. Setup + Foundational → loading screen and centered main view  
2. Add US1 → test at multiple sizes → demo (MVP)  
3. Add US2 → test resize → demo  
4. Add US3 → verify dimension-only layout → demo  
5. Polish → format, build, quickstart validation  

### Parallel Team Strategy

- One developer: Phase 1 → 2 → 3 → 4 → 5 → 6 in order  
- Multiple: After Phase 2, assign US1 / US2 / US3 to different people; merge and run Phase 6  

---

## Notes

- [P] tasks use different files or have no blocking dependency on same-phase tasks
- [Story] label links task to spec.md user stories for traceability
- Each user story is independently completable and testable
- **Quality gate**: Run `make format` and `make build` after code changes (constitution)
- File paths are under repo root: `internal/tui/`, `cmd/shop/`, `tests/`
- **SC-004 (90% test sessions)**: Validated via T014 quickstart validation plus manual or automated testing across a set of common sizes and terminals (see quickstart.md and common sizes/terminals matrix).
