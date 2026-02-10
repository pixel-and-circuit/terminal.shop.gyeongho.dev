# Research: SSH-Based Mushroom Sales TUI

**Feature**: 001-ssh-mushroom-tui  
**Purpose**: Resolve technical context and align with terminal.shop clone UI/UX and tooling.

---

## 1. Terminal.shop Clone UI Flow (IsaiahPapa/terminal.shop)

**Decision**: Use [IsaiahPapa/terminal.shop](https://github.com/IsaiahPapa/terminal.shop) as the single reference for screens, key bindings, and component layout.

**Rationale**: User requirement and spec (FR-002) require UI/UX to follow terminal.shop; this clone is a well-known, concrete implementation to mirror.

**Findings**:

- **Stack**: Rust, crossterm. We implement the same flow in Go with Charm (Bubble Tea).
- **Pages**: `Landing` (welcome), `Store`, `About`, `FAQ`. Navigation keys: **a** = Shop, **s** = About, **d** = FAQ.
- **Layout**: Fixed-width content (e.g. 60 cols), vertically centered. **Header**: logo, "a shop", "s about", "d faq", cart (total + count). **Footer**: "+/- qty   c cart   q quit". **Content**: Page-specific (landing text, product list, about text, FAQ Q&A).
- **Input**: Up/Down for scroll (e.g. FAQ); scroll offset reset on page change. Ctrl+C and **q** quit.
- **Store page**: Products listed with name, attributes, price, description, quantity (e.g. "Bags left"). No in-repo add-to-cart logic in the clone; we add product selection and add-to-cart to satisfy spec (Shop flow).
- **Components**: `header`, `footer`, `loader` (loading screen), `handler` (UIState, Page enum, show_page), `pages/` (landing, store, about, faq).

**Alternatives considered**: Designing a new TUI from scratch was rejected; spec mandates terminal.shop-style UX.

---

## 2. Go TUI: Charm Bubble Tea

**Decision**: Use [Charm Bubble Tea](https://github.com/charmbracelet/bubbletea) (and Lip Gloss / Bubbles as needed) for the TUI.

**Rationale**: Mature, well-documented Go TUI framework; Model-View-Update style fits testable state and key handling. Constitution requires Go and model-first design.

**Findings**:

- **Model**: Struct holding application state (current page, products, cart, scroll offset, dimensions).
- **Init()**: Optional initial command (e.g. fetch products via API).
- **Update(tea.Msg)**: Handle `tea.KeyMsg` (e.g. "a", "s", "d", "up", "down", "q", "ctrl+c"), `tea.WindowSizeMsg`; return updated model and optional Cmd.
- **View()**: Return string (or use Lip Gloss for styling); render header, body, footer.
- **Testing**: Update and View are pure functions; unit tests can call them with KeyMsg and assert model/view changes. No need for a real terminal in unit tests.

**Alternatives considered**: Other Go TUI libs (e.g. tview) were not required; Bubble Tea is the user-specified Charm stack and has strong docs.

---

## 3. HTTP Client Abstraction

**Decision**: Define an API client interface (e.g. `GetProducts`, `GetAbout`, `GetFAQ`, `GetCart`, `AddToCart`, `SubmitOrder`) in `internal/apiclient`. Provide a real implementation that calls `mushroom.gyeongho.dev/api` and a mock implementation for tests.

**Rationale**: Server not yet implemented; abstracting the client keeps TUI and domain logic independent of HTTP and allows tests to run without a backend.

**Findings**:

- Interface in Go: e.g. `type Client interface { GetProducts(ctx) ([]model.Product, error); ... }`.
- Real client: `http.Client` + base URL `https://mushroom.gyeongho.dev/api` (or configurable).
- Mock: In-memory struct implementing the same interface; return fixed data or configurable responses for tests.
- TUI and services depend on the interface, not `*http.Client`; inject real or mock in `main` and tests.

**Alternatives considered**: Direct HTTP in TUI was rejected for testability and constitution’s model-first principle.

---

## 4. Pre-commit and Commitlint

**Decision**: Use [pre-commit](https://pre-commit.com/) with (1) a hook that runs `make format` (or gofmt/goimports) and (2) commitlint (e.g. [commitlint-pre-commit-hook](https://github.com/alessandrojcm/commitlint-pre-commit-hook)) for conventional commits.

**Rationale**: User requirement: "pre commit 훅을 쉽게 설정할 수 있는 유명 라이브러리" and "format과 commitlint을 반드시 실행."

**Findings**:

- pre-commit: `.pre-commit-config.yaml`; `pre-commit install`; language-agnostic.
- Format: Use `local` hook running `make format` or `gofmt -w .` / `goimports -w .`.
- Commitlint: Use a repo that runs commitlint on commit-msg (e.g. `@commitlint/config-conventional`). Requires Node/npx for commitlint unless a pre-commit-native alternative is used; document in quickstart.

**Alternatives considered**: Git hooks only (no pre-commit) were rejected to use a standard, easy-to-setup library.

---

## 5. CI (GitHub Actions)

**Decision**: Add a GitHub Actions workflow that runs on push/PR: install Go, run `make format` (check only), `make build`, `make test` (or equivalent).

**Rationale**: User requirement: "github action으로 ci workflow를 추가."

**Findings**:

- Single job: checkout, set up Go, cache modules, run format check, build, test.
- Fail if format diff or build/test failure. Align with constitution quality gates.

**Alternatives considered**: None; CI is a stated requirement.

---

## 6. TDD and User Scenarios

**Decision**: Implement in TDD order: (1) define user scenarios (menu navigation, scroll, add-to-cart, checkout); (2) write tests that drive the behavior (e.g. Update with KeyMsg "a" switches to Shop; View contains expected strings); (3) implement to pass tests.

**Rationale**: User requirement: "사용자 시나리오를 먼저 구성하여 테스트 코드를 먼저 만들고 해당 테스트 코드를 만족하는 구현을 하는 TDD 방식."

**Findings**:

- Scenarios from spec: SSH → TUI; a/s/d switch pages; Up/Down scroll; product selection and add to cart; cart and checkout.
- Tests: Unit tests on Model’s Update/View with synthetic KeyMsg and WindowSizeMsg; integration tests with mock API client to verify full flows.
- Red-green-refactor: write failing test, then minimal implementation, then refactor while keeping tests green.

**Alternatives considered**: Implementation-first was rejected per user and constitution testing standards.

---

## 7. Makefile Targets

**Decision**: Provide at least `format`, `build`, `test`; optionally `lint`, `pre-commit-install`. `format` should be idempotent (e.g. gofmt -w); CI may run format-check (e.g. exit 1 if diff).

**Rationale**: Constitution requires `make format` and `make build`; pre-commit and CI call the same targets.

**Findings**:

- `format`: e.g. `gofmt -s -w .` and/or `goimports -w .`.
- `build`: `go build -o bin/mushroom ./cmd/mushroom` (or similar).
- `test`: `go test ./...`.
- Optional: `lint` with golangci-lint; document in quickstart.

**Alternatives considered**: Scripts only (no Makefile) were rejected; user asked for Makefile-based dev scripts.
