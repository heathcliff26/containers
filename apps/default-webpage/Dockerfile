FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:84d520f520edad0d8aacae26cfe8b9ea008e2c9e3b69cee86841f63f122f02e5

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
