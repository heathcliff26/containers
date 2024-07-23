FROM docker.io/library/node:lts-bookworm@sha256:34cc48658dd212db8b5494311c4a0d1c31121a172f89475e2dbd3dbe851d515c

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
