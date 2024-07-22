#!/bin/bash

# Define the package name
PACKAGE="bruno-init-suite"

# Define the platforms to package
PLATFORMS=("windows/amd64" "windows/386" "darwin/amd64" "darwin/arm64" "linux/amd64" "linux/386")

# Loop over the platforms and package the binaries
for PLATFORM in "${PLATFORMS[@]}"; do
    IFS="/" read -r -a SPLIT <<< "$PLATFORM"
    GOOS="${SPLIT[0]}"
    GOARCH="${SPLIT[1]}"
    FILENAME="$PACKAGE-$GOOS-$GOARCH"

    if [ "$GOOS" = "windows" ]; then
        zip "build/$FILENAME.zip" "build/$FILENAME.exe"
    else
        tar -czvf "build/$FILENAME.tar.gz" -C build "$FILENAME"
    fi
done

echo "Packaging complete. Archives are located in the 'build' directory."