FROM ghcr.io/heathcliff26/simple-fileserver:v1.1.1@sha256:363751e15743f191f851e5896ee987ac90eeec4ebe84964a4a1ca816061c2fb3

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
