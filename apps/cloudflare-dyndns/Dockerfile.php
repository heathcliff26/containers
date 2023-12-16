FROM docker.io/library/php:8.3.0-apache-bookworm@sha256:110e09c94d7a1962e4b9c334300a3cced6c2578c87c98fd08ddf93647db0e22e

COPY fritzbox_dyndns.php /var/www/html/index.php
