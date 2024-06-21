FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:494243c773a3d860edb2aff0960c13474b0dde47e8e47bb4237f0d14fee537de

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
