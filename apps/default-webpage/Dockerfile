FROM ghcr.io/heathcliff26/simple-fileserver:v1.3.2@sha256:51096b137f983fed8e03d4c86ac760690d4f9b611346986a11302ebd4cd219d1

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Placeholder webpage served with simple-fileserver" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/default-webpage" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="default-webpage"

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=65534:65534 --chmod=644 index.html /webroot/
COPY --chown=65534:65534 --chmod=644 kundenumfrage /webroot/kundenumfrage
