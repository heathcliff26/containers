FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:a7f928d8e4abe99a6c723bad22dd8848233996c614b62bcd39ded9add217fbc1

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
