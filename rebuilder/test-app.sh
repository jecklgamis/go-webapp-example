#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
APP_DIR=${SCRIPT_DIR}/..

echo "Running tests"
cd ${APP_DIR} && go test ./... || echo "Test failed"
