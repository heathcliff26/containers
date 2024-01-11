FROM docker.io/library/golang:1.21.6@sha256:6fbd2d3398db924f8d708cf6e94bd3a436bb468195daa6a96e80504e0a9615f2

# renovate: datasource=github-releases depName=golangci/golangci-lint
ARG GOLANGCI_LINT_VERSION=v1.55.2

RUN apt-get update && apt-get install -y gcc libgl1-mesa-dev xorg-dev && apt-get clean

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}

WORKDIR /app
