FROM docker.io/library/node:lts-bookworm@sha256:fc57c776997ec560af4ed233c5273e1fc48895d345b6517c537077acb9d73af0

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.9.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
