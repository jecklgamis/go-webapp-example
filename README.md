## go-api-server-template

This is a an example API server app using [Gorilla Mux](https://github.com/gorilla/mux).

Features:
* Serves /buildInfo - server build info
* Serves /probe/ready  - readiness probe
* Serves /probe/live  - liveness probe
* Creates Docker container image based on Ubuntu

## Building
Install `mux`:
```
go get -u github.com/gorilla/mux
```

Build the app:
```
make dist
```
This will build two binaries, one for the current OS you're building on and one for Linux AMD64 platform.

Build Docker image:
```
make image
```

## Running
Start the server on port `8080`:
```
$ bin/server -port 8080
```
Run Docker container:
```
make run
```

## Testing EndPoints
```
curl -v -X GET  http://localhost:8080/
curl -v -X GET  http://localhost:8080/buildInfo
curl -v -X GET  http://localhost:8080/probe/ready
curl -v -X GET  http://localhost:8080/probe/live
curl -v -X GET  http://localhost:8080/api
```


