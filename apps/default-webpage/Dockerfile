FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:1938aa3a4791189bb94c1675a6107269ac67f712e6e8df29568ff78fbb3f0745

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
