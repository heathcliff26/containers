FROM docker.io/library/node:lts-bookworm@sha256:b6778b063d7ba271af6560b62b1c23463d8dd14185dde5c5dd41ecd2bfef3000

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.7.2

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
