SHELL := bash

REPOSITORY ?= localhost
CONTAINER_NAME ?= cloudflare-dyndns
TAG ?= latest

GO_BUILD_FLAGS ?= -ldflags="-w -s"

default: build

build:
	podman build -t $(REPOSITORY)/$(CONTAINER_NAME):$(TAG) .

go-build:
	go build $(GO_BUILD_FLAGS) -o bin/cloudflare-dyndns ./cmd/

go-test:
	go test -v ./...

.PHONY: \
	default \
	build \
	go-build \
	go-test \
	$(NULL)
