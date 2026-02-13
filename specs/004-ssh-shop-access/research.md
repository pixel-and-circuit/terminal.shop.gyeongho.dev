# Research: SSH Access to Terminal Shop

**Feature**: 004-ssh-shop-access  
**Purpose**: Align on Charm Wish (SSH server) and Wish testsession (testing); no NEEDS CLARIFICATION—user specified both.

---

## 1. Charm Bracelet Wish — SSH Server

**Decision**: Use [Charm Bracelet Wish](https://github.com/charmbracelet/wish) for the SSH server. Server development MUST use Wish (per user input).

**Rationale**: Wish is the standard Charm library for custom SSH applications; it provides an SSH server with sensible defaults and a middleware system. The `bubbletea` middleware runs a Bubble Tea program per session with PTY and window resizing handled.

**Findings**:

- **Server**: `wish.NewServer(wish.WithAddress(...), wish.WithHostKeyPath(...), wish.WithMiddleware(...))` builds the server. Options include `WithAddress`, `WithHostKeyPath`, `WithIdleTimeout`, `WithMaxTimeout`, and middleware composition.
- **Bubble Tea middleware**: `github.com/charmbracelet/wish/bubbletea`. A handler has signature `func(ssh.Session) (tea.Model, []tea.ProgramOption)`. The middleware creates a `tea.Program` per session, connects SSH PTY I/O to the program, and handles window resize. Use `bubbletea.MakeRenderer(s)` to get a Lip Gloss renderer bound to the session for styling.
- **Middleware order**: Compose `bubbletea.Middleware(teaHandler)`, `activeterm.Middleware()` (reject non-PTY sessions), and `logging.Middleware()` as needed. Middleware runs in reverse order of registration (last added runs first).
- **Lifecycle**: `s.ListenAndServe()` in a goroutine; graceful shutdown with `s.Shutdown(ctx)` (e.g. 30s timeout). Handle `os.Interrupt`, `syscall.SIGTERM` to trigger shutdown.
- **Host key**: Use `WithHostKeyPath(".ssh/id_ed25519")` or equivalent; Wish can create a key if missing (document for deploy). For production, hostname `shop.gyeongho.dev` must resolve and the SSH service listen on port 22 (or configured port).

**Alternatives considered**: Custom SSH server or other Go SSH libraries were not chosen; user explicitly required Wish.

---

## 2. Wish testsession — Testing SSH

**Decision**: Use [Wish testsession](https://pkg.go.dev/github.com/charmbracelet/wish/testsession) for tests that need an SSH client connected to the Wish server. Test code MUST use testsession (per user input).

**Rationale**: testsession provides an in-process way to start the Wish server and obtain a client session, so integration tests can drive the TUI over SSH without a real network or external SSH binary.

**Findings**:

- **Listen**: `testsession.Listen(tb, srv)` starts the server and returns the address string. Use with `testing.TB` (e.g. `t` in `*testing.T`).
- **NewClientSession**: `testsession.NewClientSession(tb, addr, config)` returns a `*gossh.Session` and error. Requires `golang.org/x/crypto/ssh` client config (e.g. `&ssh.ClientConfig{ User: "test", Auth: []ssh.AuthMethod{ ssh.Password("") } }` or host key callback for tests).
- **New**: `testsession.New(tb, srv, cfg)` starts the server and returns a client session already connected; handles cleanup. Simplest for “connect and drive” tests.
- **Usage**: In a test, create the Wish server (same config as production but with test host key or in-memory key), pass it to `testsession.New` or `Listen` + `NewClientSession`, then send input (e.g. PTY writes) and read output to assert TUI content or behavior. Cleanup is automatic when the test ends.

**Alternatives considered**: Running a real `ssh` binary or a separate process would complicate CI and determinism; testsession is the intended Wish testing utility.

---

## 3. Integrating Existing TUI with Wish

**Decision**: The existing `internal/tui` Bubble Tea model is reused. A Wish `teaHandler` receives `ssh.Session`, builds `tui.Model` (with session PTY dimensions and the same API client as today), and returns it plus program options (e.g. `tea.WithAltScreen()`).

**Rationale**: Spec FR-002 and FR-005 require the same interface and flows over SSH; no duplicate TUI implementation.

**Findings**:

- **Model construction**: In the handler, get PTY from `s.Pty()`; create renderer with `bubbletea.MakeRenderer(s)`. Build `tui.Model` with `Width`/`Height` from `pty.Window`, and inject the same `apiclient.Client` (real or mock) used today. Pre-load products/about/FAQ in the handler or let `Model.Init()` trigger loading as today.
- **Program options**: Return `[]tea.ProgramOption{tea.WithAltScreen()}` (and any other options the current `main` uses) so behavior matches local run.
- **Binary layout**: Either (1) one binary that detects “run as SSH server” (e.g. env or flag) and calls Wish’s `ListenAndServe`, or (2) a second entrypoint (e.g. `cmd/shop-ssh`) that only runs the Wish server. Plan assumes a single entrypoint that can run in both modes unless tasks decide otherwise.

**Alternatives considered**: Building a separate TUI for SSH was rejected; one codebase for both local and SSH keeps UX and maintenance aligned.

---

## 4. Host Key and Deployment (Operational)

**Decision**: Document host key path (e.g. `.ssh/id_ed25519`) and that `shop.gyeongho.dev` must resolve and listen on the chosen port (22 or other). Key generation and DNS/deploy are operational concerns, not implementation blockers.

**Rationale**: Spec FR-003 and SC-001/SC-002 assume the hostname is configured; the plan does not mandate a specific deployment tool.

**Findings**:

- Wish can create an ed25519 key at the given path if missing (document in quickstart).
- For production, use a proper key and restrict permissions; optionally use `WithAuthorizedKeys` or auth middleware if the product adds auth later.
- Timeouts: `WithIdleTimeout` and `WithMaxTimeout` help meet SC-002 (bounded failure time); 30s shutdown in code matches spec.

**Alternatives considered**: No change to decision; deployment details can be in runbooks or separate docs.
