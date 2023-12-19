FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:2bc8f54240af4f31e341943c1929e0a13ea179212a5b77ebd6e2dba4ed6b81fe

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
