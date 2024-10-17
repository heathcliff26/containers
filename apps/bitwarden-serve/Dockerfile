FROM docker.io/library/node:lts-bookworm@sha256:7c11e52b3bec6e84949d8b6c7c9e9379130ae0e8b5d31ca6a555afda0224de19

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.9.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
