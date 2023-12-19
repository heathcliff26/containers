FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:6a581079c183039598686a0facaf4bee8250120138ed423cbefaa47792bd1d43

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
