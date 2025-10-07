#!/usr/bin/env bash
SCRIPT_DIR=$(realpath $(dirname "$0"))
cd "$SCRIPT_DIR"
LD_LIBRARY_PATH=$SCRIPT_DIR/../build/lib go test -v ./...