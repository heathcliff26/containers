FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.13@sha256:e4713d80f0f723bbb79371f86e74a88d9c9a1a4b3872f8d798c6a78d352deaed

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Placeholder webpage served with simple-fileserver" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/default-webpage" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="default-webpage"

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
