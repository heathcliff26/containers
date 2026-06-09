FROM docker.io/library/alpine:3.24.0@sha256:8ddefa941e689fc29abcdeb8dae3b3c6d139cc08ce9a52633931160701770685

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Container for scripting openwrt repository updates" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/openwrt-package-sync" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="openwrt-package-sync"

RUN apk add --no-cache curl jq cosign abuild bash

USER nobody:nobody
