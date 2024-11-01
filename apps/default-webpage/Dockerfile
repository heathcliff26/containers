FROM ghcr.io/heathcliff26/simple-fileserver:v1.1.0@sha256:84e54fcf466df26de900844fcd311a88eb566cab7de0016bb9b8959b1b4d8fc8

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
