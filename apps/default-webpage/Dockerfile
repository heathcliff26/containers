FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:31b303526140f2f52c3c0741fdf9a9c5f86b48f68f5eca6fdaaedaa5f44cd762

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
