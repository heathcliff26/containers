FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:82901ba8fad5d4d9e3645de5635eaa40ab2c78cdae02dfa89724d05eb3a3d65b

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
