FROM docker.io/library/alpine:3.24.2@sha256:3cf95fe0816180395592b8373f3ec60663f076127617bbacb4eacf9667afe2e9

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Container for scripting openwrt repository updates" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/openwrt-package-sync" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="openwrt-package-sync"

RUN apk add --no-cache curl jq cosign abuild bash

USER nobody:nobody
