FROM docker.io/nginxinc/nginx-unprivileged:1.25.2@sha256:23f009b2619ce410123917e7f274fbb810931e6162f89c8050c6c06c80e63bd6

COPY --chown=nginx:nginx --chmod=644 index.html /usr/share/nginx/html/
COPY --chown=nginx:nginx --chmod=644 kundenumfrage /usr/share/nginx/html/kundenumfrage
