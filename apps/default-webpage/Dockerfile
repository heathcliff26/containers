FROM docker.io/nginxinc/nginx-unprivileged:1.25.2@sha256:acdde1b9566df603b4e34e5bc9e1127fa4c75642022e55df3f0f3dcbe4e5a4fd

COPY --chown=nginx:nginx --chmod=644 index.html /usr/share/nginx/html/
COPY --chown=nginx:nginx --chmod=644 kundenumfrage /usr/share/nginx/html/kundenumfrage
