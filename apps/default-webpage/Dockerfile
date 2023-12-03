FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:0602584d44f2e661d5842f4daa1057fafed151f95a5f4166052a5e84562d7987

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
