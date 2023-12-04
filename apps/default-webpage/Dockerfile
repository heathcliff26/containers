FROM ghcr.io/heathcliff26/simple-fileserver:latest@sha256:9ce5f279c6ef538a90df04e98017fd32a6a1c8450514b0c485383b3da1c840bb

ENV SFILESERVER_NO_INDEX="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
