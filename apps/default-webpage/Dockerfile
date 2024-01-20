FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:58847d99ced7e44d11f9426aabeb88481a3ccc70ee72f3ce3a172a21d928f9de

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
