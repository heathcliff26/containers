FROM docker.io/library/rust:1.99.1

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Builder and CI image for rust ui projects. Can be used to cross-compile." \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/rust-qt-builder" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="rust-qt-builder"

# Installing packages and cross-compile toolchains
RUN set -e; \
    dpkg --add-architecture arm64; \
    apt-get update; \
    apt-get install -y --no-install-recommends --no-install-suggests \
        appstream \
        qt6-base-dev \
        qt6-wayland-dev \
        libfontconfig-dev \
        libxkbcommon-dev; \
    apt-get install -y --no-install-recommends --no-install-suggests \
        gcc-aarch64-linux-gnu \
        g++-aarch64-linux-gnu \
        qt6-base-dev:arm64 \
        qt6-wayland-dev:arm64 \
        libfontconfig-dev:arm64 \
        libxkbcommon-dev:arm64; \
    apt-get -qy autoremove; \
    apt-get clean; \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/*;

# Installing rust targets
RUN set -e; \
    rustup target add x86_64-unknown-linux-gnu; \
    rustup target add aarch64-unknown-linux-gnu;

# Installing rust components
RUN set -e; \
    rustup component add clippy; \
    rustup component add rustfmt;

# renovate: datasource=github-releases depName=goreleaser/goreleaser extractVersion=^(?<version>.*)$
ENV GORELEASER_VERSION=v2.18.3

# Installing goreleaser
RUN set -e; \
    curl -fsSL -o goreleaser.tar.gz "https://github.com/goreleaser/goreleaser/releases/download/${GORELEASER_VERSION}/goreleaser_Linux_x86_64.tar.gz"; \
    tar -xzf goreleaser.tar.gz -C "/usr/local/bin" goreleaser; \
    rm goreleaser.tar.gz; \
    goreleaser --version;

WORKDIR /app
