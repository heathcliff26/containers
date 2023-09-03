#!/bin/bash

set -e

if [ "${BW_HOST}" != "" ]; then
    bw config server "${BW_HOST}"
fi

# shellcheck disable=SC2155
export BW_SESSION=$(bw login "${BW_USER}" --passwordenv BW_PASSWORD --raw)

bw unlock --check

# shellcheck disable=SC2016
echo 'Running `bw server` on port 8087'
bw serve --hostname 0.0.0.0 #--disable-origin-protection
