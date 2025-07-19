.PHONY: build test clean install example ci ci-quick

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

# Run example
example: build
	cd examples && sqlc generate

# Run tests with coverage
test-coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run ./...

# Run all CI checks locally
ci: clean
	@go mod tidy
	@if [ -n "$$(git status --porcelain go.mod go.sum)" ]; then \
		echo "Error: go.mod or go.sum is not tidy"; \
		exit 1; \
	fi
	@if [ -n "$$(gofmt -l .)" ]; then \
		echo "Error: Code is not formatted"; \
		gofmt -l .; \
		exit 1; \
	fi
	@golangci-lint run ./...
	@go test -race ./...
	@go build ./cmd/sqlc-use
	@if command -v sqlc >/dev/null 2>&1; then \
		cd examples && sqlc generate && test -f gen/query_usage.json; \
	fi

# Quick CI check (no integration test)
ci-quick: clean
	@go mod tidy
	@if [ -n "$$(git status --porcelain go.mod go.sum)" ]; then \
		echo "Error: go.mod or go.sum is not tidy"; \
		exit 1; \
	fi
	@if [ -n "$$(gofmt -l .)" ]; then \
		echo "Error: Code is not formatted"; \
		exit 1; \
	fi
	@golangci-lint run ./...
	@go test -race ./...
	@go build ./cmd/sqlc-use
