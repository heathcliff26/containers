FROM docker.io/library/node:lts-bookworm@sha256:6a2972b55a5032ed961f5239c7a829c439fb54fc1d99ca5dccb2094860a47b97

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.5.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
