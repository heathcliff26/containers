FROM docker.io/library/node:lts-bookworm@sha256:978e96850e9e828eda68610829a01f70317e353dcd3d146f2fff93880f13dcb4

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.8.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
