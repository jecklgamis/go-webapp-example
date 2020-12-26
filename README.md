## go-api-server-template

This is an HTTP API server template app written in Go 

Features:
* Uses [gorilla mux](https://github.com/gorilla/mux) request router
* Uses [spf13/viper](https://github.com/spf13/viper) for config management
* Creates Docker container image based on Ubuntu
* Contains /buildInfo, /probe/ready, /probe/live handlers  
* Starts HTTP/HTTPS listeners
* Uses environment specific config in YAML format

## Requirements
* [Golang Toolchain](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/get-docker/)
* Make

## Building
```
make dist image
```
This will: 
* Run all tests 
* Build server binaries in `bin`
* Build Docker image

Explore the `Makefile` or simply type `make` in the current directory for commonly used tasks.

## Running
Run the server using the native binary:
```
$ bin/server
```
Run the server using the Docker container:
```
make up
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

Install `fswatch` (Mac OS X):
````
brew install fswatch
````

Run the rebuilder:
```
make run-rebuilder
```


## Testing

```
make test
```
This will run all `*_test.go` files in the current and child directories.

To verify the endpoints against a running server:
```
curl -v -X GET  http://localhost:8080/
curl -v -X GET  http://localhost:8080/buildInfo
curl -v -X GET  http://localhost:8080/probe/ready
curl -v -X GET  http://localhost:8080/probe/live
curl -v -X GET  http://localhost:8080/api
```

## Customizing
* Replace all `go-api-server-template` references
* Replace Go module import names in the Go sources 

Example:
```
import (
	"github.com/<your-user-name>/<your-app-name>/pkg/server"
)
```

## Navigating The Sources
* Program execution starts from `cmd/server/server.go`, it then starts  the server in `pkg/server/server.go`
* `server.go` loads environment-specific configuration under `config`.

## Contributing
Sure, send pull request?
