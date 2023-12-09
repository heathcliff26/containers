FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:64e7189a1d00634b292c3a9d0a82da2496489b37971a139061c4fc27df33d95b

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
