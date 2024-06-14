FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:aefe773b209e8e004d8e199d5cf212924a1e92539c7bc5dec9ef81812cdad39a

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
