FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:d1006d21281b38d46f3fdba8e81b8ec445ce6aea18dde5ba935316145156fe2d

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
