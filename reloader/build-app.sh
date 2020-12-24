#!/usr/local/bin/bash
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
APP_DIR=${SCRIPT_DIR}/..

echo "Rebuilding app"
cd ${APP_DIR} && go build -o bin/server cmd/server/server.go && chmod +x bin/server || echo "Build failed"
