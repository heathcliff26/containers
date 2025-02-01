FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.1@sha256:5d92e7e5bd68005fe24474185c81802ba2c51d8b811fb91a7d370fe0b22dd969

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
