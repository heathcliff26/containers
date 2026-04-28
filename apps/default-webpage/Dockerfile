FROM ghcr.io/heathcliff26/simple-fileserver:v1.3.1@sha256:0816efc0bef6064419de339c5558087677a0c8a31f50dddad9db44227ca67b20

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Placeholder webpage served with simple-fileserver" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/default-webpage" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="default-webpage"

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=65534:65534 --chmod=644 index.html /webroot/
COPY --chown=65534:65534 --chmod=644 kundenumfrage /webroot/kundenumfrage
