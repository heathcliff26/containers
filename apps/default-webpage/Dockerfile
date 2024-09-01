FROM ghcr.io/heathcliff26/simple-fileserver:v1.0.1@sha256:cbd697f570cd43eeb9821a62c438bc3cf22e2ea5d60de8cfac9d280f1f0a0f93

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
