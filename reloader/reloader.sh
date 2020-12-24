#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

function check_binaries() {
  hash go || echo "go binary not found!" && exit -1
  hash fswatch || echo "fswatch binary not found!" && exit -1
}

function sigint_handler() {
  ${SCRIPT_DIR}/kill-app.sh
}

trap 'sigint_handler' SIGINT

check_binaries
${SCRIPT_DIR}/build-app.sh
${SCRIPT_DIR}/kill-app.sh
${SCRIPT_DIR}/run-app.sh
fswatch -o pkg -o cmd | xargs -n1 -I{} sh -c "${SCRIPT_DIR}/build-app.sh && ${SCRIPT_DIR}/kill-app.sh && ${SCRIPT_DIR}/run-app.sh"
