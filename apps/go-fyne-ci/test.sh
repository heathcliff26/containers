#!/bin/bash

set -ex

base_dir="$(dirname "${BASH_SOURCE[0]}" | xargs realpath)"

export BUILDER_IMAGE="localhost/go-fyne-ci-test:local-test"

pushd "${base_dir}" > /dev/null

podman build -t "${BUILDER_IMAGE}" .

tmp_dir="$(mktemp -d /tmp/go-fyne-ci-test.XXXXXXXXXX)"

pushd "${tmp_dir}" > /dev/null

git clone --single-branch https://github.com/heathcliff26/infraspace-savegame-editor.git .

make release

popd > /dev/null
rm -rf "${tmp_dir}"

popd > /dev/null
