#!/bin/bash

set -e

script_dir="$(dirname "${BASH_SOURCE[0]}" | xargs realpath)/.."

# shellcheck source=common.sh
source "${script_dir}/hack/common.sh"

for app in "${GO_APPS[@]}"; do
    pushd "${script_dir}/apps/${app}" >/dev/null
    echo "Updating ${app}"
    go get -u ./...
    go mod tidy
    go mod vendor
    popd >/dev/null
done
