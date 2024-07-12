FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:c77caf60013aef298d6fbc660546084a117a20ab19d43b58d46cb365dcfa9c99

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
