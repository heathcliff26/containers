FROM docker.io/library/alpine:3.24.1@sha256:bec4ccd3817e7c824eb0388971a0b83fab111d586285511ba0266b77e8dc65a9

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Running keepalived in a lightweight alpine container" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/keepalived" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="keepalived"

RUN apk add --no-cache keepalived && \
	rm -rf /etc/keepalived && \
    adduser -D keepalived_script

ENTRYPOINT [ "/usr/sbin/keepalived", "--dont-fork", "--log-console" ]
