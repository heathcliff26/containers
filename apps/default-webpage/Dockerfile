FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:7895a7cfbfe5838679499600266c370e50045ad52ec2c00d6bd69eede1a6941a

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
