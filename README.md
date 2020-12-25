## go-api-server-template

This is an HTTP API server template app written in Go.

Features:
* Uses [gorilla mux](https://github.com/gorilla/mux) request router
* Uses [spf13/viper](https://github.com/spf13/viper) for config management
* Creates Docker container image based on Ubuntu
* Serves /buildInfo - Server build info
* Serves /probe/ready  - Readiness probe
* Serves /probe/live  - Liveness probe
* Starts HTTP/HTTPS listeners

## Requirements
* Go toolchain
* Docker
* Make

## Building
Build the app:
```
make dist
```
This will build two binaries in `bin` directory, one for the current OS you're building on and one for Linux 
AMD64 platform for the Docker image. This will also generate self-signed certificates.

Build Docker image:
```
make image
```

## Running
Start the server using the native binary:
```
$ bin/server
```
Start the server using Docker container:
```
make up
```

## Testing EndPoints
```
curl -v -X GET  http://localhost:8080/
curl -v -X GET  http://localhost:8080/buildInfo
curl -v -X GET  http://localhost:8080/probe/ready
curl -v -X GET  http://localhost:8080/probe/live
curl -v -X GET  http://localhost:8080/api
```

## Developing
The `reloader` dir contains helper scripts for auto building the app on file changes. It depends on `fswatch`for 
detecting file system changes in `pkg` and `cmd` directory.

Install `fswatch` (Mac OS X):
````
brew install fswatch
````

Run the reloader:
```
./reloader/reloader.sh
```

## Configuration
The `config/config-<env>.yml` contains the environment specific configuration. The config file is selected  based on 
the `APP_ENVIRONMENT` environment variable and is `dev` by default.
