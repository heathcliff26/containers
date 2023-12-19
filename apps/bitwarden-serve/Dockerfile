FROM docker.io/library/node:lts-bookworm@sha256:c454cb8e4376ab2ac080eb5358ab4d4348a0e66da42886241f972348aa912daa

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.12.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
