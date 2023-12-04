FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:8ea6b5440e9607cbdc762afe0947d60215e229f070d8cafebd37cb6019aa764c

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
