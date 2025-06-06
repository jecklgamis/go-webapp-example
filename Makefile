IMAGE_NAME:=go-webapp-example
IMAGE_TAG:=$(shell git rev-parse HEAD)
BUILD_BRANCH:=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_VERSION:=$(shell git rev-parse HEAD)

default:
	@echo "Makefile targets"
	@echo "make build - build app binary"
	@echo "make image - build Docker image"
	@echo "make up - build and run Docker image"
	@echo "make run - run Docker image"
	@echo "make rebuilder - build app automatically on file changes"
	@echo "make clean - delete built artifacts"
	@echo "make test - run tests (short tests)"
	@echo "make test-all - run all tests"
	@echo "See Makefile for details or to add your own target"
build: lint clean test server ssl-certs
up: build image run-docker
image:
	@docker build -t $(IMAGE_NAME)/$(IMAGE_TAG) .
run:
	./bin/server
run-docker:
	@docker run -p 8080:8080 -p 8443:8443 -i -t $(IMAGE_NAME)/$(IMAGE_TAG)
run-docker-shell:
	@docker run -i -t $(IMAGE_NAME)/$(IMAGE_TAG) /bin/bash
exec-docker-shell:
	@docker exec -it `docker ps | grep $(IMAGE_NAME) | awk '{print $$1}'` /bin/bash
update-modules:
	go get -u ./...
	go mod tidy
LD_FLAGS:="-X github.com/jecklgamis/go-webapp-example/pkg/version.BuildVersion=$(BUILD_VERSION) \
		  -X github.com/jecklgamis/go-webapp-example/pkg/version.BuildBranch=$(BUILD_BRANCH)"
server: server-linux-amd64
	@echo "Building $@"
	@go build -ldflags $(LD_FLAGS) -o bin/server cmd/server/server.go
	@chmod +x bin/server
server-linux-amd64:
	@echo "Building $@"
	@GOOS=linux GOARCH=amd64 go build -ldflags $(LD_FLAGS) -o bin/server-linux-amd64 cmd/server/server.go
	@chmod +x bin/server-linux-amd64
clean:
	@echo "Cleaning up artifacts"
	@rm -f $(CURDIR)/bin/*
	@go clean --testcache
ssl-certs:
	@$(CURDIR)/scripts/generate-ssl-certs.sh
test:
	@echo Running tests
	@go test -short ./...
test-all:
	@echo Running all tests
	@go test  ./...

.PHONY: rebuilder
rebuilder:
	@$(CURDIR)/scripts/rebuilder/rebuilder.sh
lint:
	@$(CURDIR)/scripts/linter.sh
