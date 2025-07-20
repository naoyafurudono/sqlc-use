# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.0.1](https://github.com/naoyafurudono/sqlc-use/commits/v0.0.1) - 2025-07-20
- Bump codecov/codecov-action from 3 to 5 by @dependabot[bot] in https://github.com/naoyafurudono/sqlc-use/pull/4
- Bump softprops/action-gh-release from 1 to 2 by @dependabot[bot] in https://github.com/naoyafurudono/sqlc-use/pull/3
- Bump actions/cache from 3 to 4 by @dependabot[bot] in https://github.com/naoyafurudono/sqlc-use/pull/2

## [Unreleased]

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
- tagpr integration for automated releases

### Documentation
- README with usage instructions
- Design documentation (DESIGN.md)
- Development log tracking progress
- JSON Schema for output format
