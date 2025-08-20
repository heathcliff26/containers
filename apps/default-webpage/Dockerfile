FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.10@sha256:daa2f241c866bee6ab34f6dbdc7fe629195880b35702b5ddac0603e4df5fdc35

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
