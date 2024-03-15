FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:29781026104439731abed3a2ac39b7f11b23d1f3a434b28ca6abcc50dca2fb5c

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
