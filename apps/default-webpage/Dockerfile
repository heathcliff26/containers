FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:a78dfdaf52fc670ff4c5cbf7e0191950fb20d857b1efa75569bd4d41c2177dcd

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
