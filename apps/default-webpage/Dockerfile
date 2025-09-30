FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.11@sha256:ea2dc6fba49c5d5a5ccd77f6459efa53e6a01dd973854a6fab329296a64814fb

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
