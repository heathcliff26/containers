FROM docker.io/library/golang:1.21.7@sha256:cdb10c5f73e8ad5fb905e8c4995192e361534a8a288edc37ca161a8eb7cf9a3a

# renovate: datasource=github-releases depName=golangci/golangci-lint
ARG GOLANGCI_LINT_VERSION=v1.55.2

RUN apt-get update && apt-get install -y gcc libgl1-mesa-dev xorg-dev && apt-get clean

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}

WORKDIR /app
