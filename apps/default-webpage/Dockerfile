FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:431d4abe4c6b15fb7a6880e6a6270dd73b07de7e3822ed857c04e212bf59fbfc

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
