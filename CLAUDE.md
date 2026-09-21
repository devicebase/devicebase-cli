# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Devicebase CLI (v1.0.0) — a cross-platform Go CLI that wraps the Devicebase HTTP API for remote device control. Supports tap, swipe, input text, launch apps, screenshots, and more.

## Build & Run

```bash
# Build locally
go build -o devicebase ./cmd/devicebase

# Run
go run ./cmd/devicebase -s <serial> tap 100,200

# Cross-platform build (macOS/Linux/Windows, amd64/arm64)
./build.sh
# or: make build

# Run tests
go test ./...
go test -v ./internal/api/...

# Run tests with coverage
go test -cover ./...
```

## Environment Variables

- `DEVICEBASE_BASE_URL` — Base URL of the Devicebase API (e.g. `https://localhost:3410`)
- `DEVICEBASE_API_KEY` — JWT Bearer token for API authentication

Both are required at runtime.

## Architecture

```
cmd/devicebase/main.go     # Entry point — calls commands.Execute()
internal/
  api/
    client.go              # HTTP client: auth, request/response handling
    device.go              # All API method wrappers (Tap, Swipe, Back, etc.)
  commands/
    root.go                # Root cobra command, --serial/-s flag, version
    register.go            # Registers all subcommands
    helper.go              # mustCreateClient(), printResult()
    tap.go                 # tap, parsePoint()
    swipe.go               # swipe, parseBounds()
    screenshot.go          # screenshot with --output/-o flag
    ...                    # One file per subcommand
```

## CLI Usage Pattern

All commands require `-s <serial>` and a subcommand:
```
devicebase -s <serial> tap <x>,<y>
devicebase -s <serial> double-tap <x>,<y>
devicebase -s <serial> long-press <x>,<y>
devicebase -s <serial> swipe <x1>,<y1>,<x2>,<y2>
devicebase -s <serial> back
devicebase -s <serial> home
devicebase -s <serial> launch-app <app_name>
devicebase -s <serial> input <text>
devicebase -s <serial> clear-text
devicebase -s <serial> current-app
devicebase -s <serial> dump-hierarchy
devicebase -s <serial> screenshot [-o file.jpg]
devicebase -s <serial> device-info
```

## Key Design Decisions

- Uses `spf13/cobra` for CLI framework with persistent `--serial/-s` flag
- API client injects `Authorization: Bearer <token>` header on every request
- Request/response types (`Point`, `Bounds`, `LaunchAppRequest`, `InputTextRequest`) defined in `internal/api/device.go`
- Cross-platform builds use `CGO_ENABLED=0` for static binaries
- Version is injected via `-ldflags` at build time
