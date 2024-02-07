FROM docker.io/library/golang:1.22.0@sha256:094e47ef90125eb49dfbc67d3480b56ee82ea9b05f50b750b5e85fab9606c2de

# renovate: datasource=github-releases depName=golangci/golangci-lint
ARG GOLANGCI_LINT_VERSION=v1.55.2

RUN apt-get update && apt-get install -y gcc libgl1-mesa-dev xorg-dev && apt-get clean

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}

WORKDIR /app
