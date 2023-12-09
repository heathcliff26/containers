FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:c0ad90578117580b9a79b194359860d8c11a9ebd1327916cccd92273d2d636ee

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
