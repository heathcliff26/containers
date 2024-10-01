FROM ghcr.io/heathcliff26/simple-fileserver:v1.0.2@sha256:c7c0358ad05e9a5eab9c32988c7d26a2f9e053bb214b42dede41294c669c774a

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
