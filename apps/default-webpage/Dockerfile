FROM docker.io/nginxinc/nginx-unprivileged:1.25.2@sha256:14aea4dc576f24c0152bae4bcb98286c9fbfde8b54d96a60d3d6515127f3dfbf

COPY --chown=nginx:nginx --chmod=644 index.html /usr/share/nginx/html/
COPY --chown=nginx:nginx --chmod=644 kundenumfrage /usr/share/nginx/html/kundenumfrage
