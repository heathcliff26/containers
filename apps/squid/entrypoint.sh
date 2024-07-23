#!/bin/sh

# Redirect logs to stdout
tail -F /var/log/squid/access.log 2>/dev/null &
tail -F /var/log/squid/error.log 2>/dev/null &
tail -F /var/log/squid/store.log 2>/dev/null &
tail -F /var/log/squid/cache.log 2>/dev/null &

# create missing cache directories and exit
/usr/sbin/squid -Nz

# Start squid
/usr/sbin/squid "$@"
