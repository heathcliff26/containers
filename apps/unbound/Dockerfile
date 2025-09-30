FROM docker.io/library/alpine:3.22.1@sha256:4bcff63911fcb4448bd4fdacec207030997caf25e9bea4045fa6c8c44de311d1

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Running unbound dns server in a lightweight alpine container" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/unbound" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="unbound"

RUN apk add --no-cache unbound && \
	rm -f /etc/unbound/unbound.conf

COPY entrypoint.sh /usr/local/bin

WORKDIR /etc/unbound

ENTRYPOINT [ "entrypoint.sh" ]
