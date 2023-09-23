FROM docker.io/nginxinc/nginx-unprivileged:1.25.2@sha256:436017b715fa21ed46e7313ec2dcd0b2d4d9b91860ebfd414061e1460856d186

COPY --chown=nginx:nginx --chmod=644 index.html /usr/share/nginx/html/
COPY --chown=nginx:nginx --chmod=644 kundenumfrage /usr/share/nginx/html/kundenumfrage
