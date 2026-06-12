FROM ghcr.io/heathcliff26/simple-fileserver:v1.3.3@sha256:546b80a8c1771051dca3c27d37a5ca5c51b3ae16c061e4451fc93639554d6e47

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Placeholder webpage served with simple-fileserver" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/default-webpage" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="default-webpage"

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=65534:65534 --chmod=644 index.html /webroot/
COPY --chown=65534:65534 --chmod=644 kundenumfrage /webroot/kundenumfrage
