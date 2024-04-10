FROM docker.io/library/node:lts-bookworm@sha256:3b676479124ebdb0348f540e5c8b64f959ec1358cea05080b711d965f5353a34

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.3.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
