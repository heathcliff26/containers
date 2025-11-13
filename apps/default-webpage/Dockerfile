FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.14@sha256:06309a0628c54de6e21b4bc9057b795495011ad08a1c05922cc4fcd1014f5eaf

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Placeholder webpage served with simple-fileserver" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/default-webpage" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="default-webpage"

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
