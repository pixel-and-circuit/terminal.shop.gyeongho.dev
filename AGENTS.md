# shop.gyeongho.dev — Agent Guidelines

Auto-generated and maintained for AI agents (AMP, Q, Bob, and others). Last updated: 2026-02-14

## Project purpose

SSH-accessible TUI for shop at shop.gyeongho.dev. Sells all products gyeongho provides: games, embedded devices, devtools, and more. Terminal.shop-style UX: Shop (a), About (s), FAQ (d), Cart (c), scroll, product list, checkout. Go + Charm Bubble Tea; API client abstracted for mocking. Deployed on AWS Lightsail, accessible via `ssh shop.gyeongho.dev`.

## Active technologies

- **Go** 1.21+ (single module at repo root)
- **Charm**: Bubble Tea (TUI), Lip Gloss (styling), Bubbles (components), Wish (SSH server)
- **API**: `internal/apiclient.Client` interface; mock and HTTP implementations
- **Infrastructure**: Terraform (AWS Lightsail), Cloudflare DNS
- **CI/CD**: GitHub Actions (CI + Terraform Plan); deployment is manual (`make deploy`)
- **Tooling**: Makefile (format, build, test, deploy, infra), pre-commit

## Project structure

```text
cmd/shop/               # TUI + SSH server entrypoint (main.go, --ssh for server mode)
internal/
  model/               # Domain: Product, Cart, Order, StoreInfo, FAQEntry
  apiclient/           # client.go (interface), mock.go, http.go
  sshsrv/              # Wish SSH server (server.go, handler.go)
  testserver/          # Test HTTP server helper (handler.go)
  tui/                 # app.go, header.go, footer.go, keys.go, loader.go
  tui/pages/           # landing, shop, about, faq, cart
tests/
  unit/                # Unit tests (model, TUI views, navigation, cart)
  integration/         # Integration tests (navigation, shop, about, faq, cart)
  ssh/                 # Wish testsession tests (connect, failure)
  e2e/                 # End-to-end tests with fixtures (products, cart, faq, about, order)
infra/                 # Terraform (main.tf, variables.tf, outputs.tf, user_data.sh)
.github/workflows/     # CI (ci.yml), Terraform plan (terraform-plan.yml)
specs/                 # Feature specs (001-ssh-shop-tui, 002-responsive-terminal-layout, 003-rebrand-to-shop, 004-ssh-shop-access)
```

## Commands

| Command | Description |
|--------|-------------|
| `make format` | Format Go code (gofmt, goimports). **Run after code changes.** |
| `make build` | Build binary to `bin/shop`. **Must pass before task complete.** |
| `make test` | Run all tests (`go test ./...`) |
| `make run` | Run TUI without building (`go run ./cmd/shop`) |
| `make build-linux` | Cross-compile for Linux AMD64 |
| `make deploy` | Build + upload + restart (full deploy to Lightsail) |
| `make deploy-status` | Check production service status |
| `make deploy-logs` | View production logs |
| `make deploy-ssh` | SSH into Lightsail instance (port 2222) |
| `make infra-init` | Initialize Terraform |
| `make infra-plan` | Preview infrastructure changes |
| `make infra-apply` | Create/update infrastructure |
| `make infra-destroy` | Destroy all infrastructure |
| `make infra-output` | Show Terraform outputs (static IP) |
| `make pre-commit-install` | Install pre-commit hooks |

## Code style

- **Go**: Standard conventions; gofmt; exported names PascalCase, unexported camelCase.
- **Design**: Model-first (pure structs/interfaces in internal/model); API behind `apiclient.Client`; TUI state in `internal/tui` (Bubble Tea); SSH server in `internal/sshsrv` (Wish).
- **Quality gate**: After any code modification, run `make format` and `make build` (see [.specify/memory/constitution.md](.specify/memory/constitution.md)).

## Governance

Project constitution: [.specify/memory/constitution.md](.specify/memory/constitution.md) — code quality, testing standards, UX/UI consistency, model-first design, quality gates.

## Recent changes

- **001-ssh-shop-tui**: Full TUI implementation — navigation (a/s/d/c), Shop, About, FAQ, Cart, mock API client, HTTP client stub, Makefile, pre-commit, CI, README.
- **002-responsive-terminal-layout**: Responsive terminal layout support.
- **003-rebrand-to-shop**: Rebrand from terminal to shop.
- **004-ssh-shop-access**: SSH server mode via Wish — `--ssh` flag, Wish SSH handler, testsession tests, deploy pipeline (Terraform + Makefile).

<!-- MANUAL ADDITIONS START -->
<!-- MANUAL ADDITIONS END -->
