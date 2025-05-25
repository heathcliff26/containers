FROM ghcr.io/heathcliff26/simple-fileserver:v1.2.5@sha256:31f483df1e2fbe94f4bd45b63d6f8964a9aeef3333a741692401c83c3491034b

ENV SFILESERVER_NO_INDEX="true"
ENV SFILESERVER_LOG="true"

COPY --chown=1001:1001 --chmod=644 index.html /webroot/
COPY --chown=1001:1001 --chmod=644 kundenumfrage /webroot/kundenumfrage
