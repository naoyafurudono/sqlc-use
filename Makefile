.PHONY: build test clean install example

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

# Install to GOPATH/bin
install: build
	go install ./cmd/sqlc-use

# Run example
example: build
	cd examples && sqlc generate

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run ./...