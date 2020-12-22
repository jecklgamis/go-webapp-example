IMAGE_NAME:=jecklgamis/go-api-server-template
IMAGE_TAG:=$(shell git rev-parse HEAD)

default:
	cat ./Makefile

dist:  server
image:
	docker build -t $(IMAGE_NAME)/$(IMAGE_TAG) .
run:
	docker run -p 8080:8080 -i -t $(IMAGE_NAME)/$(IMAGE_TAG)
run-bash:
	docker run -i -t $(IMAGE_NAME)/$(IMAGE_TAG) /bin/bash
login:
	docker exec -it `docker ps | grep $(IMAGE_NAME) | awk '{print $$1}'` /bin/bash
up: dist image run

install-deps:
	go get -u github.com/gorilla/mux

server: server-linux-amd64
	go build -o bin/server cmd/server/server.go
	@chmod +x bin/server

server-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/server-linux-amd64 cmd/server/server.go
	@chmod +x bin/server-linux-amd64

clean:
	@rm -f bin/*
