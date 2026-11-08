FROM docker.io/library/alpine:3.24.4@sha256:f59abe656b692d901e184be3445e084786602108599d2a51d229c0f5a48c71bf

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
