FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.8@sha256:1fa9461ab088cb66e7244694387ec948dd2ff946ab48cf83881014e33a24863c

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
