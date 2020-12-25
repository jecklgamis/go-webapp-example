#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
APP_DIR=${SCRIPT_DIR}/..

${APP_DIR}/bin/server &

echo $! >${SCRIPT_DIR}/server.pid
