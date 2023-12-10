FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:1a4193015a897bf29e04da7c5e4efa0412a7edff55f3ad3e52586e72fa8655c9

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
