SHELL := bash

REPOSITORY ?= localhost
CONTAINER_NAME ?= simple-fileserver
TAG ?= latest

GO_BUILD_FLAGS ?= -ldflags="-w -s"

default: build

build:
	podman build -t $(REPOSITORY)/$(CONTAINER_NAME):$(TAG) .

go-build:
	go build $(GO_BUILD_FLAGS) -o bin/simple-fileserver ./cmd/

go-test:
	go test -v ./...

.PHONY: \
	default \
	build \
	go-build \
	go-test \
	$(NULL)
