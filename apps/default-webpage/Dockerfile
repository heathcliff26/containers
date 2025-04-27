FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.4@sha256:27e211881470d7cc222ad8c30a3d26fa356ba5d31c09bdcf4d393dfbc220ca05

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
