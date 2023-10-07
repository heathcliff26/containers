FROM docker.io/nginxinc/nginx-unprivileged:1.25.2@sha256:ad9a0ffaf09f6631f0f6a11f20a981e72a4b2a0c79a9b5429af1ee5709b7d69e

COPY --chown=nginx:nginx --chmod=644 index.html /usr/share/nginx/html/
COPY --chown=nginx:nginx --chmod=644 kundenumfrage /usr/share/nginx/html/kundenumfrage
