#!/bin/bash

set -e

config_file="/config/config.yaml"
known_hosts_file="/root/.ssh/known_hosts"
source_dir="/files"

readarray hosts < <(yq -r '.hosts[]' "${config_file}")
post_hook="$(yq -r .post_hook "${config_file}")"
ssh_key="$(yq -r .ssh_key "${config_file}")"
target_dir="$(yq -r .target_dir "${config_file}")"
skip_unavailable_hosts="$(yq -r .skip_unavailable_hosts "${config_file}")"

yq -r .known_hosts "${config_file}" >> "${known_hosts_file}"

echo "Found the following files to sync:"
ls -l "${source_dir}/"

echo ""
echo "Starting file sync"
echo ""

for host in "${hosts[@]}"; do
    host="${host/$'\n'/}"
    echo "--- ${host} ---"
    if ! nc -z "${host#*@}" 22; then
        if [ "${skip_unavailable_hosts}" == "true" ]; then
            echo "Host is down, skipping"
            continue
        fi
        echo "Host is down, exiting"
        exit 1
    fi

    echo "Syncing files"

    rc=0
    rsync_out="$(rsync --rsync-path="sudo rsync" -r -c -L --out-format='%n%L' -e "ssh -i ${ssh_key}" "${source_dir}"/* "${host}":"${target_dir}"/)" || rc=$?

    if [ $rc -ne 0 ]; then
        echo "Failed to rsync directory"
        echo "${rsync_out}"
        exit $rc
    fi

    if [ "${rsync_out}" != "" ]; then
        echo "Changed files:"
        echo "${rsync_out}"
        echo "Running post hook"
        ssh -i "${ssh_key}" "${host}" "${post_hook}"
    else
        echo "No changes, skipping post hook"
    fi
done
