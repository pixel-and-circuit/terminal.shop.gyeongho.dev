# Quickstart: SSH Access to Terminal Shop

**Feature**: 004-ssh-shop-access  
**Spec**: [spec.md](./spec.md) | **Plan**: [plan.md](./plan.md)

## Prerequisites

- Go 1.21+
- Make
- Optional: pre-commit (see main project README)

## What This Feature Adds

- **SSH server**: The shop TUI is reachable via `ssh shop.gyeongho.dev` (or `ssh -p <port> localhost` for local dev). Built with Charm Wish and the Wish bubbletea middleware.
- **One TUI per session**: Each SSH connection gets its own Bubble Tea program (same `internal/tui` model).
- **Tests**: SSH server and session behavior tested with Wish **testsession** (in-process client).

## Commands

| Target        | Description |
|---------------|-------------|
| `make format` | Format Go code. Run after changes (constitution). |
| `make build`  | Build binary (TUI + SSH server entrypoint). |
| `make test`   | Run all tests (including testsession-based SSH tests). |

## Run SSH Server Locally

1. **Build**  
   `make build`

2. **Start the SSH server** (example; actual flag or subcommand may vary per implementation)  
   - The binary may run as an SSH server by default when a host key path or port is configured, or via a flag (e.g. `./bin/shop --ssh`).  
   - Example: server listens on `localhost:2222` and uses a host key at `.ssh/id_ed25519` (Wish can create the key if missing).

3. **Connect from another terminal**  
   `ssh -p 2222 -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null localhost`  
   You should see the shop TUI (header, a/s/d/c, etc.).

4. **Production**  
   For `ssh shop.gyeongho.dev`, the hostname must resolve and the service listen on port 22 (or the configured port). Deploy and DNS are outside this quickstart.

## Run Tests

```bash
make test
```

- **Unit / integration**: Existing tests for model and TUI remain; no change to how they are run.
- **SSH / testsession**: New tests use `github.com/charmbracelet/wish/testsession` to start the Wish server and obtain a client session, then drive or assert behavior. Run with the rest of the suite via `go test ./...`.

## Verify Success Criteria (Manual)

- **SC-001**: From a machine with SSH and network, run `ssh shop.gyeongho.dev` (or the local equivalent). TUI appears within 15 seconds.
- **SC-002**: Stop the server or use an unreachable host; connection attempt fails with a clear message and does not hang (e.g. timeout within 30 seconds).
- **SC-003**: In a connected session, complete at least one flow (e.g. open Shop, add to cart).
- **SC-004**: Document `ssh shop.gyeongho.dev` (e.g. in README or docs) so new users can discover and use it.

## Reference

- Research: [research.md](./research.md) (Wish, bubbletea middleware, testsession).
- Data model: [data-model.md](./data-model.md) (session scope, no new domain entities).
- Contracts: [contracts/README.md](./contracts/README.md) (SSH access contract only).
