FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:fa0ce59cb433b5cdb7b40351385e3cb9d37590f57969fd692dde876ff7e0fc00

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
