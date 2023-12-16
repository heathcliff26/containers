FROM docker.io/library/php:8.3.0-apache-bookworm@sha256:b32b8dea83c94461c44e1bb31ad3524b1183d15cbfcc66b5981e9a019b817072

COPY fritzbox_dyndns.php /var/www/html/index.php
