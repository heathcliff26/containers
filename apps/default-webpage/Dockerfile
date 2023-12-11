FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:94e6200617d5e09ebad78bc160d1c9c457355e4a03205e1b288a1dc447415e55

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
