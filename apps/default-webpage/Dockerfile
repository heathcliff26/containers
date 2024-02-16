FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:b8f4ef3bb66a14d984deac5626d81a906d1f422497ed09aeda0c05a770d633e2

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
