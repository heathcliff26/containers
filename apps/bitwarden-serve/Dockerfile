FROM docker.io/library/node:lts-bookworm@sha256:99981c3d1aac0d98cd9f03f74b92dddf30f30ffb0b34e6df8bd96283f62f12c6

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.12.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
