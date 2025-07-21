.PHONY: build test clean install example ci

# Default target
.DEFAULT_GOAL := help

# Build the plugin binary
build:
	go build -o sqlc-use ./cmd/sqlc-use

# Run all tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -f sqlc-use
	rm -rf examples/gen
	rm -f coverage.out coverage.html

# Install to GOPATH/bin
install: build
	go install ./cmd/sqlc-use

example: build ## Run example
	cd examples && sqlc generate

test-coverage: ## Run tests with coverage
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

fmt: ## Format code
	go fmt ./...

# Lint code
lint:
	golangci-lint run ./...

ci: build fmt lint test clean ## Run all CI checks locally
