FROM docker.io/library/node:lts-bookworm@sha256:2ff27bf5505de4b7dd9b92f079c995a1863fb16db7e462ceec0b96782f07e350

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
