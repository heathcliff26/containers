FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:e2e755a49c4d1b7ebf124cecaf7acb3a398707258412622898b04cf86e7d295b

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
