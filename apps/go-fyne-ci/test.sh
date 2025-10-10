#!/bin/bash

set -ex

base_dir="$(dirname "${BASH_SOURCE[0]}" | xargs realpath)"

export FYNE_CROSS_IMAGE="localhost/go-fyne-ci-test:local-test"

pushd "${base_dir}" > /dev/null

CONTAINER_ENGINE="podman"
if command -v docker >/dev/null 2>&1; then
    CONTAINER_ENGINE="docker"
fi

"${CONTAINER_ENGINE}" build -t "${FYNE_CROSS_IMAGE}" .

tmp_dir="$(mktemp -d /tmp/go-fyne-ci-test.XXXXXXXXXX)"

pushd "${tmp_dir}" > /dev/null

git clone --single-branch --depth 1 https://github.com/heathcliff26/infraspace-savegame-editor.git .

hack/fyne-cross.sh linux

popd > /dev/null
rm -rf "${tmp_dir}"

popd > /dev/null
