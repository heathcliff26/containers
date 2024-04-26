FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:937a49c28fa850994c3dae826327a2382aff32d2e7b33639803497b122b397d3

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
