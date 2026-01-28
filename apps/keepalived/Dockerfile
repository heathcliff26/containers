FROM docker.io/library/alpine:3.23.3@sha256:25109184c71bdad752c8312a8623239686a9a2071e8825f20acb8f2198c3f659

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Running keepalived in a lightweight alpine container" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/keepalived" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="keepalived"

RUN apk add --no-cache keepalived && \
	rm -rf /etc/keepalived && \
    adduser -D keepalived_script

ENTRYPOINT [ "/usr/sbin/keepalived", "--dont-fork", "--log-console" ]
