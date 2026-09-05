FROM docker.io/library/rust:1.98.0

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Builder and CI image for rust ui projects. Can be used to cross-compile." \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/rust-qt-builder" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="rust-qt-builder"

# renovate: datasource=github-releases depName=goreleaser/goreleaser extractVersion=^(?<version>.*)$
ENV GORELEASER_VERSION=v2.18.1

COPY install-deps.sh .
RUN ./install-deps.sh

WORKDIR /app
