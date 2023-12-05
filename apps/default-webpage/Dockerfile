FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:f3085bc4921e3eb379202fb05eb2a3547227fee4d9ae37a7fce84d8bfb830c3b

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
