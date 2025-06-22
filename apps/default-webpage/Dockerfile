FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.6@sha256:1bd15021f28250da712e3a423aa0185cfe2d2f8aa13bf5d3c49feec0f245f960

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
