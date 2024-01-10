FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:56f3ef22e0e1c619bf9c895d20a1ecc39353c7e767ead41aeec8fd374d539431

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
