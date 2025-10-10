FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.12@sha256:ab63a38086fb3443bf0d6b9ef29245f7e3a6263cf486476b62f4a7af3fb1e47d

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Placeholder webpage served with simple-fileserver" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/default-webpage" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="default-webpage"

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
