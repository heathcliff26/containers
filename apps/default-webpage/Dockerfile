FROM docker.io/nginxinc/nginx-unprivileged:1.25.3@sha256:198442edfc1f70ffa15c21ba80976903ff376b9aba3c77b11148f2018f2928b9

COPY --chown=nginx:nginx --chmod=644 index.html /usr/share/nginx/html/
COPY --chown=nginx:nginx --chmod=644 kundenumfrage /usr/share/nginx/html/kundenumfrage
