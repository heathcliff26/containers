#!/bin/bash

set -ex

if [ "$(pwd)" == "/github/workspace" ] && [ "$(whoami)" == "root" ]; then
    exec su runner "$0" -- "$@"
fi

eval "${@}"
