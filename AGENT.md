# AGENT.md - Semver Go Package

## Build/Lint/Test Commands
- **Test all**: `go test`
- **Test with verbose**: `go test -v`
- **Test specific function**: `go test -run TestString`
- **Test with coverage**: `go test -cover`
- **Build**: `go build`
- **Format**: `go fmt`
- **Vet (lint)**: `go vet`

## Architecture
Simple Go package implementing semantic versioning (SemVer). Single module with:
- `semver.go`: Core Version struct with String(), Equal(), Less() methods
- `commit_hash.go`: Build info utility (has bug on line 11 - missing semicolon)
- `semver_test.go`: Table-driven tests for all methods
- No external dependencies, uses only Go standard library

## Code Style
- Package comments with copyright header
- Struct fields: PascalCase (Major, Minor, Patch, PreRelease, Build)
- Method receivers: single letter (v Version)
- Table-driven tests with testCases slice
- Error messages: descriptive with expected/actual format
- Import grouping: standard library only
- Function comments follow Go conventions (// FunctionName description)
