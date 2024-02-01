FROM docker.io/library/golang:1.21.6@sha256:4d1942cb703999c3acd03b8efe5cc588cb5e39bace931d876de170e16d44e2cd

# renovate: datasource=github-releases depName=golangci/golangci-lint
ARG GOLANGCI_LINT_VERSION=v1.55.2

RUN apt-get update && apt-get install -y gcc libgl1-mesa-dev xorg-dev && apt-get clean

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}

WORKDIR /app
