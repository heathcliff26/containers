FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:3b5e4514ba693e0a7f528343d24e4877e175d17e8ed108a456aa930a0d578ad7

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
