FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:a10162bc3585bdffba4abeec84c22537dbf20c43f71fe9723b688957b779b975

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
