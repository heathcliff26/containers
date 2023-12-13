FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:77e66d5daf147cc86757af52373f4ba79c5a36ca9859f936b1ac70188d96b4ac

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
