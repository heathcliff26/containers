#!/bin/bash

set -e

script_dir="$(dirname "${BASH_SOURCE[0]}" | xargs realpath)/.."

APP="${1}"
OUT_DIR="${script_dir}/coverprofiles"

if [ ! -d "${OUT_DIR}" ]; then
    mkdir "${OUT_DIR}"
fi

go test -coverprofile="${OUT_DIR}/cover-${APP}.out" "./apps/${APP}/..."
go tool cover -html "${OUT_DIR}/cover-${APP}.out" -o "${OUT_DIR}/cover-${APP}.html"
