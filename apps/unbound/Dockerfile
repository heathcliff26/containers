FROM docker.io/library/alpine:3.24.0@sha256:660e0827bd401543d81323d4886abbd08fda0fe3ba84337837d0b11a67251283

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
