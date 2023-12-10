FROM docker.io/library/php:8.3.0-apache-bookworm@sha256:c55d99c94f804ee54177ba00961d2441333b277a67e6a6901341cb251b47f638

COPY fritzbox_dyndns.php /var/www/html/index.php
