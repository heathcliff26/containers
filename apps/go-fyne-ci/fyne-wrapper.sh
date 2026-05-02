#!/bin/bash

set -e

args=""

while [[ $# -gt 0 ]]; do
    if [[ "${1}" != -ldflags=* ]]; then
        args="$args $1"
    fi
    shift
done

fyne $args
