#!/bin/bash

# Define the package name
PACKAGE="bruno-init-suite"

# Define the version
VERSION=$(git describe --tags --always)

# Define the platforms to build for
PLATFORMS=("windows/amd64" "windows/386" "darwin/amd64" "darwin/arm64" "linux/amd64" "linux/386")

# Create the build directory
mkdir -p build

# Loop over the platforms and build the binaries
for PLATFORM in "${PLATFORMS[@]}"; do
    IFS="/" read -r -a SPLIT <<< "$PLATFORM"
    GOOS="${SPLIT[0]}"
    GOARCH="${SPLIT[1]}"
    OUTPUT="build/$PACKAGE-$GOOS-$GOARCH"

    if [ "$GOOS" = "windows" ]; then
        OUTPUT+=".exe"
    fi

    echo "Building for $PLATFORM..."
    env GOOS="$GOOS" GOARCH="$GOARCH" go build -ldflags "-X bruno-init-suite/internal/version.Version=$VERSION" -o "$OUTPUT"
done

echo "Build complete. Binaries are located in the 'build' directory."
