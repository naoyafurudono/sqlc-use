# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2025-01-20

### Added
- Support for UNION, UNION ALL clauses
- Support for other set operations (EXCEPT, INTERSECT) via SetOprStmt

### Changed
- **BREAKING**: Output format now follows dirty-effects.schema.json specification
  - Output now includes `version` field (1.0)
  - Query effects wrapped in `effects` object
  - Effects use string format: `{ operation[table] | operation[table] }`
- Simplified internal architecture
  - Removed intermediate TableOperation and UsageReport types
  - Analyzers now return effects strings directly
  - Formatter interface simplified to only handle EffectsReport
- Improved performance by eliminating conversion steps

### Technical Details
- Reduced codebase by 91 lines (105 additions, 196 deletions)
- Direct processing flow without intermediate transformations
- All tests updated to match new structure

## [0.1.0] - 2025-01-20

### Added
- Initial implementation of sqlc-use plugin
- MySQL support using TiDB parser
- JSON output format for query-table operations
- Support for SELECT, INSERT, UPDATE, DELETE operations
- Support for complex queries with JOINs
- Comprehensive test coverage
- CI/CD pipeline with GitHub Actions
- golangci-lint integration
- Process plugin architecture
- Fully qualified package name support via plugin options

### Documentation
- README with usage instructions
- Design documentation (DESIGN.md)
- Development log tracking progress
- JSON Schema for output format