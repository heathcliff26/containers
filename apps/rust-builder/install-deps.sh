#!/bin/bash

set -ex

arches="arm64"

for arch in ${arches}; do
    echo "Adding architecture ${arch}"
    dpkg --add-architecture "${arch}"
done

echo "Updating package lists"
apt-get update

echo "Installing native dependencies"
apt-get install -y --no-install-recommends --no-install-suggests \
        musl-tools \
        upx

echo "Adding rust target for architecture x86_64"
rustup target add "x86_64-unknown-linux-gnu" "x86_64-unknown-linux-musl"

for arch in ${arches}; do
    case "${arch}" in
        arm64)
            pkg_arch="aarch64"
            musl_arch="arm_64"
            ;;
        *)
            pkg_arch="${arch}"
            musl_arch="${arch}"
    esac

    echo "Adding rust target for architecture ${arch}"
    rustup target add "${pkg_arch}-unknown-linux-gnu" "${pkg_arch}-unknown-linux-musl"

    echo "Installing dependencies for architecture ${arch}"
    apt-get install -y --no-install-recommends --no-install-suggests \
        "gcc-${pkg_arch}-linux-gnu" \
        "g++-${pkg_arch}-linux-gnu" \
        "libssl-dev:${arch}"

    echo "Installing musl cross compile toolchain for architecture ${arch}"
    curl -SL -o musl-toolchain.tar.xz "https://github.com/dyne/musl/releases/download/${DYNE_MUSL_VERSION}/dyne-gcc-musl-${musl_arch}.tar.xz"
    tar -xJf musl-toolchain.tar.xz -C /opt
    rm musl-toolchain.tar.xz
done

echo "Installing goreleaser"
curl -SL -o goreleaser.tar.gz "https://github.com/goreleaser/goreleaser/releases/download/${GORELEASER_VERSION}/goreleaser_Linux_x86_64.tar.gz"
tar -xzf goreleaser.tar.gz -C "/usr/local/bin" goreleaser
rm goreleaser.tar.gz
goreleaser --version

echo "Cleaning up"
apt-get clean
rm -rf /var/lib/apt/lists/* /var/cache/apt/*
