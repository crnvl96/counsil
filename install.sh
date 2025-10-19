#!/bin/bash

set -e

# Detect OS and arch
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then
    ARCH="amd64"
elif [ "$ARCH" = "aarch64" ]; then
    ARCH="arm64"
fi

# Download from releases
URL="https://github.com/crnvl96/counsil/releases/latest/download/counsil-${OS}-${ARCH}"
echo "Downloading from $URL"
curl -L -o /tmp/counsil "$URL"
chmod +x /tmp/counsil
sudo mv /tmp/counsil /usr/local/bin/counsil
echo "Counsil installed successfully!"