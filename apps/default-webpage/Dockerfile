FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:eed6c691a96f9c6c22a741cb1c370d8106d3a982098c3f38b28a85fd59b02f7d

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
