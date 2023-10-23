FROM docker.io/nginxinc/nginx-unprivileged:1.25.2@sha256:ce3a7e7cfeecb3265f331ae02046737fa29ac22b5131f06f59959aaefc551516

COPY --chown=nginx:nginx --chmod=644 index.html /usr/share/nginx/html/
COPY --chown=nginx:nginx --chmod=644 kundenumfrage /usr/share/nginx/html/kundenumfrage
