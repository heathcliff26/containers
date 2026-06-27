#!/bin/bash

set -ex

base_dir="$(dirname "${BASH_SOURCE[0]}" | xargs realpath)"

curl -SL -o "${base_dir}/crd-extractor.sh" https://raw.githubusercontent.com/datreeio/CRDs-catalog/main/Utilities/crd-extractor.sh
chmod +x "${base_dir}/crd-extractor.sh"
curl -SL -o "${base_dir}/openapi2jsonschema.py" https://raw.githubusercontent.com/yannh/kubeconform/master/scripts/openapi2jsonschema.py
