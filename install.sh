#!/usr/bin/env bash
set -euo pipefail

# =============================================================================
# Devicebase CLI Installation Script
# Usage: curl -fsSL https://git.uusense.cn/devicebase/devicebase-cli/releases/latest/download/install.sh | bash
# =============================================================================

VERSION="${VERSION:-v2026.5.15}"
BASE_URL="${BASE_URL:-https://github.com/uusense/devicebase-cli/releases}"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
BINARY_NAME="devicebase"
CHECKSUM_URL="${CHECKSUM_URL:-}"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

info() { echo -e "${GREEN}[INFO]${NC} $1"; }
warn() { echo -e "${YELLOW}[WARN]${NC} $1"; }
error() { echo -e "${RED}[ERROR]${NC} $1"; }

# Detect OS
detect_os() {
    case "$(uname -s)" in
        Linux*)     echo "linux" ;;
        Darwin*)    echo "darwin" ;;
        CYGWIN*)    echo "windows" ;;
        MINGW*)     echo "windows" ;;
        MSYS*)      echo "windows" ;;
        *)          error "Unsupported OS: $(uname -s)"; exit 1 ;;
    esac
}

# Detect architecture
detect_arch() {
    case "$(uname -m)" in
        x86_64)          echo "amd64" ;;
        aarch64)         echo "arm64" ;;
        armv7l)          echo "arm" ;;
        i386)            echo "386" ;;
        i686)            echo "386" ;;
        arm64)           echo "arm64" ;;
        *)               error "Unsupported architecture: $(uname -m)"; exit 1 ;;
    esac
}

# Get latest version number
get_latest_version() {
    if [ "$VERSION" != "latest" ]; then
        echo "$VERSION"
        return
    fi

    local release_url="${BASE_URL}/tags/${VERSION}"
    local tag

    if command -v curl >/dev/null 2>&1; then
        tag=$(curl -fsSL "${release_url}" 2>/dev/null | grep -oP '(?<=<a href="/devicebase/devicebase-cli/releases/tag/)[^"]+' | head -1 || true)
    elif command -v wget >/dev/null 2>&1; then
        tag=$(wget -qO- "${release_url}" 2>/dev/null | grep -oP '(?<=<a href="/devicebase/devicebase-cli/releases/tag/)[^"]+' | head -1 || true)
    fi

    if [ -n "$tag" ]; then
        echo "$tag"
    else
        error "Failed to fetch latest version"
        exit 1
    fi
}

# Build download URL
build_url() {
    local os="$1"
    local arch="$2"
    local ext=""
    [ "$os" = "windows" ] && ext=".exe"

    echo "${BASE_URL}/download/${VERSION}/${BINARY_NAME}-${os}-${arch}${ext}"
}

# Download file
download() {
    local url="$1"
    local dest="$2"

    info "Downloading ${url} ..."

    if command -v curl >/dev/null 2>&1; then
        curl -fsSL -o "$dest" "$url"
    elif command -v wget >/dev/null 2>&1; then
        wget -q -O "$dest" "$url"
    else
        error "Neither curl nor wget found. Please install curl or wget."
        exit 1
    fi
}

# Verify checksum
verify_checksum() {
    local file="$1"
    local expected_checksum="$2"

    if [ -z "$expected_checksum" ]; then
        warn "No checksum provided, skipping verification"
        return
    fi

    info "Verifying checksum ..."

    local actual_checksum
    case "${#expected_checksum}" in
        64)   actual_checksum=$(sha256sum "$file" | cut -d' ' -f1) ;;
        128)  actual_checksum=$(sha512sum "$file" | cut -d' ' -f1) ;;
        *)    actual_checksum=$(shasum -a 256 "$file" | cut -d' ' -f1) ;;
    esac

    if [ "$actual_checksum" != "$expected_checksum" ]; then
        error "Checksum mismatch! Expected: $expected_checksum, Got: $actual_checksum"
        exit 1
    fi

    info "Checksum verified."
}

# Check if running as root for system-wide install
check_permissions() {
    if [ -d "$INSTALL_DIR" ] && [ ! -w "$INSTALL_DIR" ]; then
        warn "${INSTALL_DIR} is not writable. Using sudo..."
        SUDO="sudo"
    else
        SUDO=""
    fi
}

# Install binary
install_binary() {
    local src="$1"
    local dest="${INSTALL_DIR}/${BINARY_NAME}"

    check_permissions

    info "Installing to ${dest} ..."

    if [ -n "${SUDO:-}" ]; then
        $SUDO cp "$src" "$dest"
        $SUDO chmod +x "$dest"
    else
        cp "$src" "$dest"
        chmod +x "$dest"
    fi

    info "Installed successfully: ${dest}"
}

# Main
main() {
    info "Devicebase CLI Installer"
    info "Version: ${VERSION}"
    echo ""

    local os=$(detect_os)
    local arch=$(detect_arch)

    info "Detected: ${os}/${arch}"

    # Get version
    VERSION=$(get_latest_version)
    info "Installing version: ${VERSION}"
    echo ""

    # Build download URL
    local url=$(build_url "$os" "$arch")
    local tmp_dir=$(mktemp -d)
    local tmp_file="${tmp_dir}/${BINARY_NAME}"

    trap "rm -rf $tmp_dir" EXIT

    # Download
    download "$url" "$tmp_file"

    # Verify checksum if provided
    if [ -n "$CHECKSUM_URL" ]; then
        local checksum_file="${tmp_dir}/checksum.txt"
        download "$CHECKSUM_URL" "$checksum_file"
        local checksum
        checksum=$(grep -E "${os}-${arch}" "$checksum_file" | cut -d' ' -f1 || true)
        verify_checksum "$tmp_file" "$checksum"
    fi

    # Install
    install_binary "$tmp_file"

    echo ""
    info "Devicebase CLI installed successfully!"
    info "Run 'devicebase --help' to get started."
    info "Don't forget to set DEVICEBASE_API_KEY environment variables."
}

main
