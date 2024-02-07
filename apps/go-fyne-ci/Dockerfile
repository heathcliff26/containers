FROM docker.io/library/golang:1.21.7@sha256:8144f2d44d2262fa930b437200fc4ada624d8a0b9c83d688e2a6f545d097c45b

# renovate: datasource=github-releases depName=golangci/golangci-lint
ARG GOLANGCI_LINT_VERSION=v1.55.2

RUN apt-get update && apt-get install -y gcc libgl1-mesa-dev xorg-dev && apt-get clean

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}

WORKDIR /app
