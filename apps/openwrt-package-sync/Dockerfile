FROM docker.io/library/alpine:3.23.3@sha256:eb37f58646a901dc7727cf448cae36daaefaba79de33b5058dab79aa4c04aefb

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Container for scripting openwrt repository updates" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/openwrt-package-sync" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="openwrt-package-sync"

RUN echo "@testing https://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && \
        apk add --no-cache opkg-utils@testing curl jq cosign@testing signify@testing

USER 1001
