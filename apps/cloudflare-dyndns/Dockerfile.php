FROM docker.io/library/php:8.3.0-apache-bookworm@sha256:f51a5168c4832b923fc22a52f1ef5475cdd02333c8efd924e5c4f773000fb0fc

COPY fritzbox_dyndns.php /var/www/html/index.php
