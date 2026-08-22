FROM ghcr.io/heathcliff26/simple-fileserver:v1.3.4@sha256:5afd3bfe2cf8b48e51a1e923e7418ed2a24bc38816d9933d4e8549b752b96c44

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Placeholder webpage served with simple-fileserver" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/default-webpage" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="default-webpage"

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=65534:65534 --chmod=644 index.html /webroot/
COPY --chown=65534:65534 --chmod=644 kundenumfrage /webroot/kundenumfrage
