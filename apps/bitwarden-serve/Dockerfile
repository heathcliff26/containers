FROM docker.io/library/node:lts-bookworm@sha256:35e3b7fc91f95fc3c9dd76348f7bc5a7aa7525cbf6bc6363333b4c3d6a41e0f4

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.8.2

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
