#!/bin/bash
set -e
cd /app && bin/server --port 8080 --keyFile server.key --certFile server.crt --httpsPort 8443
