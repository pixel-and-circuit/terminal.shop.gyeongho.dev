.PHONY: format build test run pre-commit-install

format:
	gofmt -s -w .
	@which goimports >/dev/null 2>&1 && goimports -w . || true

build:
	go build -o bin/shop ./cmd/shop

test:
	go test ./...

run:
	go run ./cmd/shop

pre-commit-install:
	@which pre-commit >/dev/null 2>&1 && pre-commit install || echo "pre-commit not installed; run: pip install pre-commit"
