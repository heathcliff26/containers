FROM docker.io/nginxinc/nginx-unprivileged:1.25.2@sha256:80cc0d8dba216b9364d1449368fe44fdd1bf2050c49ade2f1da2fa8d183598f5

COPY --chown=nginx:nginx --chmod=644 index.html /usr/share/nginx/html/
COPY --chown=nginx:nginx --chmod=644 kundenumfrage /usr/share/nginx/html/kundenumfrage
