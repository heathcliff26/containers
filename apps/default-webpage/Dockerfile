FROM docker.io/nginxinc/nginx-unprivileged:1.25.3@sha256:d47e89176fcf7a83aaf220ea1c560cb232cb9646144eec5be742575df5d9ccc5

COPY --chown=nginx:nginx --chmod=644 index.html /usr/share/nginx/html/
COPY --chown=nginx:nginx --chmod=644 kundenumfrage /usr/share/nginx/html/kundenumfrage
