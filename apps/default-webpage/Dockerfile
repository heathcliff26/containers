FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:ae0fbe2b625090d005aba81f5c231968161e69cff373d6ead8e043497be66af5

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
