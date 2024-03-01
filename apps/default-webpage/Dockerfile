FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:bfdda60818303454a8a2aa19fae7d17046f6b57f645bb031a53d72176bcd4ae5

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
