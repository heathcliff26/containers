FROM docker.io/nginxinc/nginx-unprivileged:1.25.2@sha256:3aecdd9321c1aac4587d49114b3c4ef8d83327602638d8e48c4d2ad7e36789ac

COPY --chown=nginx:nginx --chmod=644 index.html /usr/share/nginx/html/
COPY --chown=nginx:nginx --chmod=644 kundenumfrage /usr/share/nginx/html/kundenumfrage
