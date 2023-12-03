FROM docker.io/nginxinc/nginx-unprivileged:1.25.3@sha256:a4cb3ee8f71116750ae3dc857031b7c74ec3b21df3d878e1c63204c206553e46

COPY --chown=nginx:nginx --chmod=644 index.html /usr/share/nginx/html/
COPY --chown=nginx:nginx --chmod=644 kundenumfrage /usr/share/nginx/html/kundenumfrage
