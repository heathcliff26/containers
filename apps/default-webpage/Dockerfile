FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:f6dbe9c3d34fdb096c0efbdb0d706efec8b6beed25694021e2173c45d3fa27a0

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
