# Tasks: SSH Access to Terminal Shop

**Input**: Design documents from `/specs/004-ssh-shop-access/`  
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/, quickstart.md

**Tests**: Plan and constitution require SSH server behavior tested with Wish testsession; test tasks are included for US1 and US2.

**Organization**: Tasks are grouped by user story so each story can be implemented and tested independently.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: User story (US1, US2) for story-phase tasks only
- Include exact file paths in descriptions

## Path Conventions

- **Single project** (per plan): `cmd/shop/`, `internal/tui/`, `internal/apiclient/`, `internal/model/`, optional `internal/ssh/`, `tests/unit/`, `tests/integration/`, optional `tests/ssh/`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Verify environment and add Wish/testsession dependencies.

- [ ] T001 Verify feature branch `004-ssh-shop-access` and that `specs/004-ssh-shop-access/` contains plan.md, spec.md, research.md, data-model.md, quickstart.md, contracts/
- [ ] T002 Add Charm Wish and middlewares to go.mod: `github.com/charmbracelet/wish`, `github.com/charmbracelet/wish/bubbletea`, `github.com/charmbracelet/wish/activeterm`, `github.com/charmbracelet/wish/logging`; add `golang.org/x/crypto/ssh` for testsession client in tests

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Minimal Wish SSH server that accepts connections and runs a stub TUI. Required before US1 (real TUI) and US2 (failure/timeout behavior).

**⚠️ CRITICAL**: No user story work can begin until this phase is complete.

- [ ] T003 [P] Implement stub tea handler returning a minimal tea.Model (e.g. single "Connected" message) for SSH session in `cmd/shop/ssh_handler.go` or `internal/ssh/handler.go`
- [ ] T004 Create Wish server in `cmd/shop/main.go`: `wish.NewServer` with `WithAddress`, `WithHostKeyPath`, `WithMiddleware(bubbletea.Middleware(stub), activeterm.Middleware(), logging.Middleware())`
- [ ] T005 Add mode switch in `cmd/shop/main.go`: when `--ssh` or env `SHOP_SSH=1`, run Wish server (`ListenAndServe` + signal handling); otherwise run existing local TUI (`tea.NewProgram(m).Run()`)
- [ ] T006 Configure `WithIdleTimeout` and `WithMaxTimeout` on Wish server (e.g. 30s) in `cmd/shop/main.go` so connections do not hang (spec SC-002)

**Checkpoint**: Running `./bin/shop --ssh` (or equivalent) starts an SSH server; connecting yields stub TUI. Foundation ready for user stories.

---

## Phase 3: User Story 1 - Reach the Shop via SSH (Priority: P1) 🎯 MVP

**Goal**: User runs `ssh shop.gyeongho.dev` (or local equivalent); connection establishes and the full terminal shop TUI is shown with navigation (Shop, About, FAQ, Cart).

**Independent Test**: Run `ssh -p <port> localhost` (or `ssh shop.gyeongho.dev` when deployed); shop interface appears; keys a/s/d/c and scroll work.

### Implementation for User Story 1

- [ ] T007 [US1] Replace stub tea handler with real handler in `cmd/shop/ssh_handler.go` (or `internal/ssh/handler.go`): from `ssh.Session` get PTY, use `bubbletea.MakeRenderer(s)`, build `tui.Model` with Width/Height from `pty.Window` and inject `apiclient.Client`, return model and `[]tea.ProgramOption{tea.WithAltScreen()}`
- [ ] T008 [US1] Pre-load products, about, FAQ in handler (or ensure `Model.Init()` runs) so TUI shows content after connection in `cmd/shop/ssh_handler.go` or `internal/ssh/handler.go`
- [ ] T009 [P] [US1] Add testsession test in `tests/ssh/connect_test.go` (or `tests/integration/ssh_connect_test.go`): start Wish server with `testsession.Listen` or `testsession.New`, obtain client session, assert output contains expected TUI strings (e.g. header or "Shop"), and assert that at least one full flow (e.g. view product and add to cart) can be completed in the same session (spec SC-003)

**Checkpoint**: User Story 1 complete. `ssh` to server shows full shop TUI; navigation and flows work. Can demo as MVP.

---

## Phase 4: User Story 2 - Clear Feedback When Connection Fails (Priority: P2)

**Goal**: When the user cannot connect (service down, host unreachable), they get a clear failure message; connection attempts do not hang indefinitely.

**Independent Test**: Stop the server or use unreachable host; `ssh` attempt fails with clear message and does not hang (e.g. timeout within 30s).

### Implementation for User Story 2

- [ ] T010 [US2] Add graceful shutdown in `cmd/shop/main.go`: on SIGINT/SIGTERM call `s.Shutdown(ctx)` with 30s timeout so server stops cleanly and in-flight connections end
- [ ] T011 [P] [US2] Add test in `tests/ssh/failure_test.go` that connection attempt to stopped or unreachable server fails with clear error or bounded timeout (e.g. testsession or dial with short timeout)

**Checkpoint**: User Story 2 complete. Connection failures and timeouts behave per spec SC-002.

---

## Phase 5: Polish & Cross-Cutting Concerns

**Purpose**: Documentation, quality gate, and success-criteria validation.

- [ ] T012 [P] Document `ssh shop.gyeongho.dev` access in `README.md` so new users can discover and use it (spec SC-004)
- [ ] T013 Run `make format` and `make build`; validate steps in `specs/004-ssh-shop-access/quickstart.md`; validate SC-001 (TUI reachable within 15 seconds) manually per quickstart (e.g. run `ssh` and confirm interface appears within 15s)

---

## Dependencies & Execution Order

### Phase Dependencies

- **Phase 1 (Setup)**: No dependencies — start immediately.
- **Phase 2 (Foundational)**: Depends on Phase 1 — BLOCKS all user stories.
- **Phase 3 (US1)**: Depends on Phase 2 — implement real TUI handler and testsession test.
- **Phase 4 (US2)**: Depends on Phase 2; can run after or in parallel with Phase 3 (graceful shutdown and failure tests).
- **Phase 5 (Polish)**: Depends on Phase 3 and 4 being complete.

### User Story Dependencies

- **User Story 1 (P1)**: Starts after Phase 2. No dependency on US2.
- **User Story 2 (P2)**: Starts after Phase 2. Timeout config (T006) is in Phase 2; US2 adds shutdown and failure tests.

### Within Each User Story

- US1: Handler (T007) then pre-load (T008); test (T009) can be written in parallel or after.
- US2: Shutdown (T010) and failure test (T011) can be done in parallel.

### Parallel Opportunities

- Phase 1: T002 can run after T001.
- Phase 2: T003 (stub handler) is [P] with T004/T005/T006 (different files); T004–T006 are sequential in main.go.
- Phase 3: T009 [P] (test) can run in parallel with T007–T008 or after.
- Phase 4: T010 and T011 [P] can run in parallel.
- Phase 5: T012 [P] (docs) and T013 (format/build) can overlap.

---

## Parallel Example: User Story 1

```bash
# After T007 and T008 are done, or in parallel if test is written to fail first:
# Task T009: Add testsession test in tests/ssh/connect_test.go
# (Implement T007/T008 in cmd/shop/ssh_handler.go or internal/ssh/handler.go)
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup (T001–T002).
2. Complete Phase 2: Foundational (T003–T006).
3. Complete Phase 3: User Story 1 (T007–T009).
4. **STOP and VALIDATE**: Run `./bin/shop --ssh`, connect with `ssh -p 2222 localhost`, confirm TUI and navigation.
5. Demo/deploy if ready.

### Incremental Delivery

1. Setup + Foundational → server runs with stub TUI.
2. Add US1 → full TUI over SSH → MVP.
3. Add US2 → shutdown and clear failure behavior → spec complete.
4. Polish → README and quickstart validation.

### Parallel Team Strategy

- After Phase 2: One developer can do US1 (T007–T009), another US2 (T010–T011), then merge and run Phase 5.

---

## Notes

- [P] = different files, no dependencies; safe to run in parallel.
- [US1]/[US2] map tasks to user stories for traceability.
- Each user story is independently testable per spec.
- **Quality gate**: Run `make format` and `make build` after code changes (constitution).
- Host key: Wish can create at `WithHostKeyPath` if missing; document in README/quickstart for deploy.
