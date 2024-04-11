FROM docker.io/library/node:lts-bookworm@sha256:4b2369e0ff40bdc8522a281278afe3ea257b2db63a95eceb68ac2447ff0d8cda

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.3.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
