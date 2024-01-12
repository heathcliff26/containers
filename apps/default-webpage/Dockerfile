FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:4f1bfd65b02e2cceb0c69f5e887d1ad852404b3cad44ba98b79e65889ab543db

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
