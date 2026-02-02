# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build Commands

```bash
make build       # Lint, test, build server binary (bin/server), and generate SSL certs
make test        # Run short tests (excludes integration tests)
make test-all    # Run all tests including integration tests
make lint        # Run go vet on all sources
make up          # Build, create Docker image, and run container
make run         # Run existing Docker image
make rebuilder   # Auto-rebuild on file changes (requires fswatch)
```

Run a single test:
```bash
go test -run TestName ./pkg/server/handler/...
```

## Architecture

This is a skeleton HTTP API server using gorilla/mux router with Prometheus metrics instrumentation.

**Entry point flow:**
- `cmd/server/server.go` → `pkg/server/server.go:Start()` → loads config, sets up routes, starts HTTP/HTTPS listeners

**Key packages:**
- `pkg/server/` - Server configuration and startup logic
- `pkg/server/handler/` - HTTP handlers (api, buildInfo, probes, root)
- `pkg/server/middleware/` - Access logging middleware
- `pkg/version/` - Build version info (injected via ldflags)
- `pkg/integration_test/` - Integration tests (skipped with `-short` flag)

**Configuration:**
- Environment-specific YAML configs in `configs/config-<env>.yaml`
- Selected via `APP_ENV` environment variable (defaults to `dev`)
- Config struct defined in `pkg/server/config.go`

**HTTP endpoints:**
- `/` - Root handler
- `/api` - API handler (GET, POST)
- `/buildInfo` - Build version information
- `/probe/ready`, `/probe/live` - Health probes
- `/metrics` - Prometheus metrics
