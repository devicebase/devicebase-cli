# Devicebase CLI

A cross-platform Go CLI tool that wraps the Devicebase HTTP API for remote device control. Supports tap, swipe, input text, launch apps, screenshots, and more.

## Requirements

- Go 1.24+
- Devicebase API server running and accessible
- API credentials (API KEY)

## Environment Variables

Set these before running the CLI:

| Variable | Description |
|----------|-------------|
| `DEVICEBASE_BASE_URL` | Base URL of the Devicebase API (e.g. `https://api.devicebase.cn`) |
| `DEVICEBASE_API_KEY` | API KEY for authentication |

## Quick Start

```bash
# Build locally
go build -o devicebase ./cmd/devicebase

# Run
DEVICEBASE_BASE_URL=https://api.devicebase.cn \
DEVICEBASE_API_KEY=your_api_key \
  ./devicebase -s <serial> tap 100,200
```

## Installation

### Pre-built Binaries

Download from the `build/` directory after running the build script.

### Build from Source

```bash
# Clone and build
go build -o devicebase ./cmd/devicebase

# Cross-platform build (macOS/Linux/Windows, amd64/arm64)
./build.sh
```

The cross-platform build produces static binaries with no CGO dependencies.

## Usage

All commands require the `-s <serial>` flag to specify the target device.

### Touch Interactions

```bash
# Tap at coordinates
devicebase -s <serial> tap 100,200

# Double tap
devicebase -s <serial> double-tap 100,200

# Long press
devicebase -s <serial> long-press 100,200

# Swipe from (x1,y1) to (x2,y2)
devicebase -s <serial> swipe 100,200,300,400
```

### Navigation

```bash
# Press back button
devicebase -s <serial> back

# Press home button
devicebase -s <serial> home
```

### App Management

```bash
# Launch an app by name
devicebase -s <serial> launch-app com.example.app

# Get the current foreground app
devicebase -s <serial> current-app
```

### Text Input

```bash
# Input text
devicebase -s <serial> input "Hello World"

# Clear text field
devicebase -s <serial> clear-text
```

### Device Management

```bash
# List all devices
devicebase list-devices

# Filter by keyword (brand/model/serial/name)
devicebase list-devices --keyword "iPhone"

# Filter by state (busy/free/offline)
devicebase list-devices --state free

# Combine filters
devicebase list-devices --keyword "Samsung" --state busy
```

### Device Information

```bash
# Get device info
devicebase -s <serial> device-info

# Dump UI hierarchy (Android accessibility tree)
devicebase -s <serial> dump-hierarchy

# Take a screenshot
devicebase -s <serial> screenshot

# Save screenshot to a file
devicebase -s <serial> screenshot -o screenshot.jpg
```

### Global Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--serial` | `-s` | Device serial number (required) |
| `--help` | `-h` | Show help |
| `--version` | | Show version |

## Architecture

```
cmd/devicebase/main.go     # Entry point
internal/
  api/
    client.go              # HTTP client: auth, request/response handling
    device.go              # All API method wrappers
  commands/
    root.go                # Root cobra command, --serial flag
    register.go            # Registers all subcommands
    helper.go              # mustCreateClient(), printResult()
    tap.go                 # tap command, parsePoint()
    swipe.go               # swipe command, parseBounds()
    screenshot.go          # screenshot with --output flag
    ...                    # One file per subcommand
```

## Development

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detection
go test -race ./...

# Format code
go fmt ./...

# Lint
go vet ./...
```

## License

MIT
