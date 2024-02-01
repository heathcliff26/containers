FROM docker.io/library/golang:1.21.6@sha256:0c22572a0b01ce93bb9d1f0bea2f198b6b827225a194a1f3a185d0fd4b4513b7

# renovate: datasource=github-releases depName=golangci/golangci-lint
ARG GOLANGCI_LINT_VERSION=v1.55.2

RUN apt-get update && apt-get install -y gcc libgl1-mesa-dev xorg-dev && apt-get clean

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}

WORKDIR /app
