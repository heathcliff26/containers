FROM ghcr.io/heathcliff26/simple-fileserver:v1.0.0@sha256:7c48b5282798eae467ffeaec2b2a0200cf2d0204446edef5d0836b42aceae565

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
