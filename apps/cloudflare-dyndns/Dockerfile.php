FROM docker.io/library/php:8.3.0-apache-bookworm@sha256:df17ddfded53708a37a41fd7a9eca94e523182d93f7fde31ba442c002dc9e446

COPY fritzbox_dyndns.php /var/www/html/index.php
