#!/bin/sh

unbound-anchor
chown -R unbound:unbound /etc/unbound

exec unbound -d -p
