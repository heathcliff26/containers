SHELL := bash

REPOSITORY ?= localhost
CONTAINER_NAME ?= simple-fileserver
TAG ?= latest
CLI_TAG ?= cli

GO_BUILD_FLAGS ?= -ldflags="-w -s"

default: build

build:
	podman build -t $(REPOSITORY)/$(CONTAINER_NAME):$(TAG) .

go-build:
	make -p bin
	go build $(GO_BUILD_FLAGS) -o bin/simple-fileserver

.PHONY: \
	default \
	build \
	go-build \
	$(NULL)
