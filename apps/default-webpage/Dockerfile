FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:f48891bbc12b59433bd2d338389d12518518092d5fa581360e3a3b91733dca1f

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
