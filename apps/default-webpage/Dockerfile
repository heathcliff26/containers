FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.9@sha256:4fca8cb7268a268afefd497e3f6877e07f52f1fbacc73058aed8d65be4e472fe

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
