IMAGE_NAME:=jecklgamis/go-api-server-template
IMAGE_TAG:=$(shell git rev-parse HEAD)
BUILD_BRANCH:=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_VERSION:=$(shell git rev-parse HEAD)

default:
	cat ./Makefile

dist: clean test server ssl-certs
up: dist image run

image:
	docker build -t $(IMAGE_NAME)/$(IMAGE_TAG) .
run:
	docker run -p 8080:8080 -p 8443:8443 -i -t $(IMAGE_NAME)/$(IMAGE_TAG)
run-bash:
	@docker run -i -t $(IMAGE_NAME)/$(IMAGE_TAG) /bin/bash
login:
	@docker exec -it `docker ps | grep $(IMAGE_NAME) | awk '{print $$1}'` /bin/bash

install-deps:
	go get -u github.com/gorilla/mux
LD_FLAGS:="-X github.com/jecklgamis/go-api-server-template/pkg/version.BuildVersion=$(BUILD_VERSION) \
		  -X github.com/jecklgamis/go-api-server-template/pkg/version.BuildBranch=$(BUILD_BRANCH)"
server: server-linux-amd64
	@go build -ldflags $(LD_FLAGS) -o bin/server cmd/server/server.go
	@chmod +x bin/server
server-linux-amd64:
	@GOOS=linux GOARCH=amd64 go build -ldflags $(LD_FLAGS) -o bin/server-linux-amd64 cmd/server/server.go
	@chmod +x bin/server-linux-amd64
clean:
	@echo "Cleaning up artifacts"
	@rm -f bin/*
ssl-certs:
	@./generate-ssl-certs.sh
test:
	@echo Running tests
	@go test ./...
run-rebuilder:
	@./rebuilder/rebuilder.sh
