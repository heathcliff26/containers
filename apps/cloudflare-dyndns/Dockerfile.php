FROM docker.io/library/php:8.3.0-apache-bookworm@sha256:0497b44a7619b7b88662097760c404d4f3fe810f74a9febe09fced71072212c0

COPY fritzbox_dyndns.php /var/www/html/index.php
