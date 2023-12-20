FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:0055827c76557f28219058aa0d3ca5248c6fdab49ed196a19e8791058a1bd5e4

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
