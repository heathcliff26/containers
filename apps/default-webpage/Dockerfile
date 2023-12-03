FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:9d166683d08332de966911e62ea423a79a6eb92804d10201daa08e0207949b79

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
