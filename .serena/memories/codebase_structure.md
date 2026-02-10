# Codebase structure (mushroom.gyeongho.dev)

```
cmd/mushroom/           # Main entry (TUI binary)
internal/
  model/                # Domain: product.go, cart.go, order.go, storeinfo.go, faq.go
  apiclient/            # client.go (interface), mock.go, http.go
  tui/                  # app.go, header.go, footer.go, keys.go, loader.go
  tui/pages/            # landing.go, shop.go, about.go, faq.go, cart.go
tests/
  unit/                 # *_test.go for model and TUI logic
  integration/          # navigation, shop, about, faq tests
specs/001-ssh-mushroom-tui/  # plan, spec, data-model, contracts, quickstart, tasks
```

Entrypoint: cmd/mushroom/main.go. Domain and API client are in internal/; tests in tests/.
