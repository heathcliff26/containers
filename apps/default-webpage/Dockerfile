FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:879ca23b7b6a57ce8705e93fa625da5d4574b374a06fea27669900de441f2fc0

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
