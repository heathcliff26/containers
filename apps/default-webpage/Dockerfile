FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:ea8cf558011087243678d5e6f290793b0040506c1a217530659e1ad3216f68fd

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
