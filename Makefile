IMAGE_NAME:=jecklgamis/go-api-server-template
IMAGE_TAG:=$(shell git rev-parse HEAD)
BUILD_BRANCH:=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_VERSION:=$(shell git rev-parse HEAD)

default:
	@echo "Makefile targets"
	@echo "make dist - build app binaries"
	@echo "make image - build Docker image"
	@echo "make run - run Docker image"
	@echo "make up - build and run Docker image (dist, image run)"
	@echo "make login - attach /bin/bash shell to a runnning Docker container"
	@echo "make rebuilder - build app automatically on file changes"
	@echo "make clean - delete built artifacts"
	@echo "make test - run tests (short tests)"
	@echo "make test-all - run all tests"
	@echo "make lint - run linter"
	@echo "make ssl-certs - generate self-signed certificates"
	@echo "See Makefile for details or to add your own target"
dist: lint clean test server ssl-certs
up: dist image run
image:
	@docker build -t $(IMAGE_NAME)/$(IMAGE_TAG) .
run:
	@docker run -p 8080:8080 -p 8443:8443 -i -t $(IMAGE_NAME)/$(IMAGE_TAG)
run-bash:
	@docker run -i -t $(IMAGE_NAME)/$(IMAGE_TAG) /bin/bash
login:
	@docker exec -it `docker ps | grep $(IMAGE_NAME) | awk '{print $$1}'` /bin/bash
install-deps:
	@@go get -u github.com/gorilla/mux
	@go get -u golang.org/x/lint/golint
LD_FLAGS:="-X github.com/jecklgamis/go-api-server-template/pkg/version.BuildVersion=$(BUILD_VERSION) \
		  -X github.com/jecklgamis/go-api-server-template/pkg/version.BuildBranch=$(BUILD_BRANCH)"
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
rebuilder:
	@$(CURDIR)/scripts/rebuilder/rebuilder.sh
lint:
	@$(CURDIR)/scripts/linter.sh
