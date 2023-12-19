FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:e96ad7e742032fc936f68931c1681be3c1e01511eb169771a86e7548ac4e29f7

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
