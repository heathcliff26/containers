FROM docker.io/library/node:lts-bookworm@sha256:b515e3af865b74f0a4da0a1f31881a6d1b9793bf5aaad782fcf2c933f3b66412

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
