FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:10c746894c142d2e073d26e62ab7a0b3aabfa6f082363957811230c442a4eb3c

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
