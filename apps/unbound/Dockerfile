FROM docker.io/library/alpine:3.23.4@sha256:c7989ac7a27b473e1795973c98d714f62b4dd0b134594d36880505ce0bfd716b

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
