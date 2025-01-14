FROM docker.io/library/node:lts-bookworm@sha256:4f7fb7f5f716f2175f5e7716f2a6cb310e82f3c72152b78d8c1ff58068988c0b

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.12.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
