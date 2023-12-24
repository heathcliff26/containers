#!/bin/bash

set -e

script_dir="$(dirname "${BASH_SOURCE[0]}" | xargs realpath)/.."

declare -a APPS=(cloudflare-dyndns simple-fileserver speedtest-exporter)

coverprofiles=""
webviews=""

for app in "${APPS[@]}"; do
    "${script_dir}/hack/go-coverprofile.sh" "${app}"
    coverprofiles+="    <br><a href="cover-${app}.out">${app}</a></br>"
    webviews+="    <br><a href="cover-${app}.html">${app}</a></br>"
done

cat << EOF > "${script_dir}/coverprofiles/index.html"
<html>
<title>heathcliff26/containers</title>
<body>
    <h2>Generated webviews per app:</h2>
${webviews}

    <h2>Go test coverprofiles per app:</h2>
${coverprofiles}
</body>
</html>
EOF
