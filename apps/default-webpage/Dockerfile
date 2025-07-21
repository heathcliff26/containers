FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.7@sha256:749bfe80b34d488a07a9f87e925cdedcc0e1f79755074486a4623e60e3c02bed

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
