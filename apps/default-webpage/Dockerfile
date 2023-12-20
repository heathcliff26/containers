FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:fe5d652ea4404d01fb967cfeb5e701924840f4fd749aee7b5c1dc354f0013318

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
