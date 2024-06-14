FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:ec0b5e21d017b0305894d67bf6be9110c5e37bc40baa61d2a44ec303c2f52c80

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
