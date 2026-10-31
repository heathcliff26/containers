FROM docker.io/library/rust:1.99.1

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Builder and CI image for rust projects. Can be used to cross-compile." \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/rust-builder" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="rust-builder"

# Installing packages and cross-compile toolchains
RUN set -e; \
    dpkg --add-architecture arm64; \
    apt-get update; \
    apt-get install -y --no-install-recommends --no-install-suggests \
        musl-tools \
        upx; \
    apt-get install -y --no-install-recommends --no-install-suggests \
        gcc-aarch64-linux-gnu \
        g++-aarch64-linux-gnu \
        libssl-dev:arm64; \
    apt-get -qy autoremove; \
    apt-get clean; \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/*;

# Installing rust targets
RUN set -e; \
    rustup target add x86_64-unknown-linux-gnu x86_64-unknown-linux-musl; \
    rustup target add aarch64-unknown-linux-gnu aarch64-unknown-linux-musl;

# Installing rust components
RUN set -e; \
    rustup component add clippy; \
    rustup component add rustfmt;

# renovate: datasource=github-releases depName=goreleaser/goreleaser extractVersion=^(?<version>.*)$
ENV GORELEASER_VERSION=v2.19.0

# Installing goreleaser
RUN set -e; \
    curl -fsSL -o goreleaser.tar.gz "https://github.com/goreleaser/goreleaser/releases/download/${GORELEASER_VERSION}/goreleaser_Linux_x86_64.tar.gz"; \
    tar -xzf goreleaser.tar.gz -C "/usr/local/bin" goreleaser; \
    rm goreleaser.tar.gz; \
    goreleaser --version;


# renovate: datasource=github-releases depName=dyne/musl extractVersion=^(?<version>.*)$
ENV DYNE_MUSL_VERSION=v2.2.1

# Installing musl cross compile toolchain
RUN set -e; \
    curl -fsSL -o musl-toolchain.tar.xz "https://github.com/dyne/musl/releases/download/${DYNE_MUSL_VERSION}/dyne-gcc-musl-arm_64.tar.xz"; \
    tar -xJf musl-toolchain.tar.xz -C /opt; \
    rm musl-toolchain.tar.xz;

WORKDIR /app
