# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a Go library for working with gRPC in Senzing Garage projects. It provides helpers for parsing gRPC URLs and configuring gRPC client/server options.

## Common Commands

### Build

```bash
make build              # Build for current platform (output in target/)
```

### Test

```bash
make test               # Run tests with formatted output
go test ./...           # Run tests directly
go test -v ./... -run TestName  # Run a specific test
```

### Lint

```bash
make lint               # Run all linters (golangci-lint, govulncheck, cspell)
make golangci-lint      # Run golangci-lint only
make fix                # Auto-fix linting issues
```

### Coverage

```bash
make coverage           # Generate coverage report and open in browser
make check-coverage     # Check coverage thresholds (80% file/package/total)
```

### Other

```bash
make run                # Run the program directly via go run
make clean              # Clean build artifacts and caches
make dependencies       # Update Go dependencies
make dependencies-for-development  # Install development tools
```

## Architecture

### Project Structure

- `main.go` - Entry point with example usage
- `grpcurl/` - Package for parsing gRPC URLs
  - `main.go` - Package declaration
  - `grpcurl.go` - URL parsing implementation (`Parse` function)
- `clientoptions/` - Package for gRPC client dial options
  - `main.go` - Interface definition (`ClientOptions`)
  - `clientoptions.go` - Interface implementation
- `serveroptions/` - Package for gRPC server options
  - `doc.go` - Package documentation

### Key Patterns

**Interface Pattern**: Define interfaces in `main.go` of a package, implement in separate files.

**URL Parsing**: The `grpcurl.Parse()` function parses gRPC URLs in the format `grpc://host:port` and returns the target and dial options.

### Makefile System

The Makefile uses OS detection with platform-specific includes:

- `makefiles/osdetect.mk` - Detects OS type and architecture
- `makefiles/{darwin,linux,windows}.mk` - OS-specific target implementations

## Linting Configuration

Golangci-lint config: `.github/linters/.golangci.yaml`

- Line length: 120 characters
- Coverage thresholds: 80% (configurable in `.github/coverage/testcoverage.yaml`)
- Uses extensive linter set including exhaustruct, wrapcheck, err113
