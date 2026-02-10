# Style and conventions (mushroom.gyeongho.dev)

- **Language**: Go. Follow standard Go conventions (effective go, gofmt).
- **Naming**: PascalCase for exported, camelCase for unexported; short names where clear.
- **Structure**: Pure interfaces/structs for domain (model-first); implementations behind interfaces (apiclient.Client) for testability.
- **TUI**: Bubble Tea Model/Update/View; key bindings in internal/tui/keys.go; terminal.shop-style header/footer.
- **Testing**: Unit tests for model and TUI Update/View; integration tests with mock API client. TDD where specified in tasks.
- **Quality**: Constitution requires make format and make build after code changes; no trailing whitespace; single blank line between sections in docs.
