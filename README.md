## go-api-server-template

[![Go Report Card](https://goreportcard.com/badge/github.com/jecklgamis/go-api-server-template)](https://goreportcard.com/report/github.com/jecklgamis/go-api-server-template)

This is an HTTP API server template app written using Go. 

If you have the toolchains installed, run `make up` to build and run the app in Docker on port 8080.

Features:
* Uses [gorilla mux](https://github.com/gorilla/mux) request router
* Uses [spf13/viper](https://github.com/spf13/viper) for config management
* Exposes [Prometheus](prometheus.io) metrics endpoint and instrumented handlers
* Creates Docker container image based on Ubuntu
* Contains /buildInfo, /probe/ready, /probe/live, /metrics handlers  
* Starts HTTP/HTTPS listeners
* Uses environment specific config in YAML format


## Requirements
* [Golang Toolchain](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/get-docker/)
* [GNU Make](https://www.gnu.org/software/make/)

## Building
```
make install-deps
make dist image
```
This will: 
* Run all tests 
* Build server binaries in `bin` (one for the OS you're building on and one for Linux AMD64 platform to be used 
  inside the Docker image)
* Build Docker image `jecklgamis/go-api-server-template`

Explore the `Makefile` or simply type `make` in the current directory for commonly used tasks.

## Running
Run the server using the native binary:
```
$ bin/server
```
Run the server using the Docker container:
```
make run
```

Verify the endpoints:
```
curl http://localhost:8080/
curl http://localhost:8080/buildInfo
curl http://localhost:8080/probe/ready
curl http://localhost:8080/probe/live
curl http://localhost:8080/api
curl http://localhost:8080/metrics
```

## Configuration
The `config/config-<env>.yml` contains the environment specific configuration. The config file is selected  based on 
the `APP_ENVIRONMENT` variable and is `dev` by default.

## Docker Image Namespace
The Docker image uses Ubuntu base image. It has the following directory namespace:
* `/app` is the base app directory
* `/app/bin` contains the server binary 
* `/app/config` contains the environment-specific config files in YAML format

## Developing
The `rebuilder` directory contains helper scripts for automatically building the app on file changes. 
It depends on `fswatch`for detecting file system changes in `pkg` and `cmd` directories.

Install `fswatch` (Mac OS):
````
brew install fswatch
````

Run the rebuilder:
```
make rebuilder
```


## Linting
```
make lint
```
This wil run `golint` on all the Go sources it can find.

## Testing

```
make test
```
This will run all `*_test.go` files it can find.
`

## Navigating The Sources
* Program execution starts from `cmd/server/server.go`, it then starts the server in `pkg/server/server.go`
* `server.go` loads environment-specific configuration under `config`.

## Customizing
* Replace all `go-api-server-template` references
* Replace Go module import names in the Go sources 

Example:
```
import (
	"github.com/<your-user-name>/<your-app-name>/pkg/server"
)
```

## Contributing
Sure, send pull request?
