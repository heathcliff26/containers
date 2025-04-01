FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.3@sha256:c9aaf4a3f244f96ec478fc4d1c423f2c6bf480a70cdbaa8ee6dc1bd6404739e1

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
