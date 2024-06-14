FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:81d60d98d346c91ea215ee761862402fa2833022552ca7240c3ee78db8b594c8

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
