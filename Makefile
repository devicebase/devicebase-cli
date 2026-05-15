#!/usr/bin/env bash
set -euo pipefail

APP_NAME="devicebase"
VERSION="2026.5.15"
BUILD_DIR="build"
MAIN_PKG="./cmd/devicebase"

LDFLAGS="-s -w -X github.com/uusense/devicebase-cli/internal/commands.Version=${VERSION}"

PLATFORMS=(
    "darwin/amd64"
    "darwin/arm64"
    "linux/amd64"
    "linux/arm64"
    "windows/amd64"
    "windows/arm64"
)

rm -rf "${BUILD_DIR}"
mkdir -p "${BUILD_DIR}"

for PLATFORM in "${PLATFORMS[@]}"; do
    IFS="/" read -r GOOS GOARCH <<< "${PLATFORM}"
    OUTPUT="${BUILD_DIR}/${APP_NAME}-${GOOS}-${GOARCH}"
    if [ "${GOOS}" = "windows" ]; then
        OUTPUT="${OUTPUT}.exe"
    fi

    echo "Building ${PLATFORM}..."
    CGO_ENABLED=0 GOOS="${GOOS}" GOARCH="${GOARCH}" \
        go build -ldflags "${LDFLAGS}" -o "${OUTPUT}" "${MAIN_PKG}"
done

echo ""
echo "Build complete. Binaries in ${BUILD_DIR}/:"
ls -lh "${BUILD_DIR}/"
