FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:85013b1f6d0a2b9a4bfbbcebfe725295db2c947e307520771944874b2febaccf

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
