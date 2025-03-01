FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.2@sha256:1057c8813b9ff43c509c762bd018b38f9b19ef5bbe70bfb2acc9a68f1866844a

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
