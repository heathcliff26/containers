FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:fb99da3cd2facbd3f27839553410d5c1ebd672c4d59ffd86dedcfdd22d3b3fd1

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
