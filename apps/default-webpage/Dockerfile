FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:a94df648c2c5235d2a655134319c3e0c4c70dbfbebd5940ca6762ed2592edbcc

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
