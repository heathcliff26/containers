FROM docker.io/library/php:8.3.0-apache-bookworm@sha256:ac693528efd04f4a176938c818e035bf1d910940eeddfc6027c62c355fac664d

COPY fritzbox_dyndns.php /var/www/html/index.php
