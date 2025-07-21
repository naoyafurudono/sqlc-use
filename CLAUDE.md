# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is an sqlc plugin written in Go that analyzes SQL queries and generates JSON output describing which tables each query operates on and what type of operation (SELECT, INSERT, UPDATE, DELETE) is performed.

## Essential Commands

### Build and Development

```bash
make build          # Build the plugin binary (sqlc-use)
make install        # Install to GOPATH/bin
make clean          # Remove build artifacts
```

### Testing

```bash
make test           # Run all tests with verbose output
make test-coverage  # Run tests with coverage report
go test -v ./internal/analyzer  # Run specific package tests
```

### Code Quality

```bash
make fmt            # Format code with go fmt
make lint           # Run golangci-lint
make ci             # Run all CI checks locally (build, fmt, lint, test, clean)
```

### Integration Testing

```bash
make example        # Run the example with sqlc to test the plugin
```

## Architecture

The codebase follows a clean architecture pattern with clear separation of concerns:

- **cmd/sqlc-use/**: Entry point that initializes the plugin and handles stdio communication with sqlc
- **internal/plugin/**: Implements the sqlc plugin interface, orchestrating the analysis and formatting
- **internal/analyzer/**: Contains the SQL parsing logic using TiDB parser for MySQL syntax
- **internal/models/**: Defines core data structures (TableOperation, QueryAnalysis)
- **internal/formatter/**: Handles JSON output formatting

Key flow:

1. sqlc sends a GenerateRequest via stdin
2. Plugin parses each SQL query using TiDB parser
3. Analyzer extracts table operations from the AST
4. Formatter creates JSON output mapping query names to their operations
5. Plugin returns the JSON file in GenerateResponse

## Important Development Notes

- Currently only supports MySQL syntax via TiDB parser
- The plugin communicates with sqlc through stdin/stdout using protobuf
- Output file is always named `query-table-operations.json`
- The plugin must be executable and in PATH or specified with full path in sqlc.yaml
- Integration tests require sqlc to be installed

## Testing Strategy

- Unit tests for individual components (analyzer, formatter)
- Integration test in CI that runs the actual sqlc generation
- Test coverage is tracked via Codecov
- All tests run with race detection enabled

## QA

After your development is complete, run the `make ci` commands to ensure everything works as expected.
