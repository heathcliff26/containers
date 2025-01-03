FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.0@sha256:983024ed8e66fd198c36aa06d2b9166a39b187d318bdc09be8e415035747210a

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
