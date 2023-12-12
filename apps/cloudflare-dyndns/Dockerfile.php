FROM docker.io/library/php:8.3.0-apache-bookworm@sha256:8281e439f1407ff597edd6c1c7aa184261fbee6428133094ecb1cbf60aabdfc6

COPY fritzbox_dyndns.php /var/www/html/index.php
