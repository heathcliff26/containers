FROM docker.io/library/php:8.3.0-apache-bookworm@sha256:45ed72f21858fcb440f9a6af04972e1a17fc0642a647cd159bc3c92f07ab6b09

COPY fritzbox_dyndns.php /var/www/html/index.php
