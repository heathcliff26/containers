FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:fd7327319cc67ddcc8ab18c8c695fca28abfd0e6ce21bd84b602eb300978dde3

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
