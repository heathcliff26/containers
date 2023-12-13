FROM docker.io/library/php:8.3.0-apache-bookworm@sha256:88cec9a79de8d33d75d47c2c9da80f926744f984bec2960ce0c2b67e28c35b79

COPY fritzbox_dyndns.php /var/www/html/index.php
