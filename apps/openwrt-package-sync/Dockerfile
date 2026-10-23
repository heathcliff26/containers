FROM docker.io/library/alpine:3.24.3@sha256:be2656061b536c1883c4034790e21d0a09c5aaaf51968b89368a3c0371ea2c9a

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Container for scripting openwrt repository updates" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/openwrt-package-sync" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="openwrt-package-sync"

RUN apk add --no-cache curl jq cosign abuild bash

USER nobody:nobody
