FROM docker.io/library/node:lts-bookworm@sha256:a6835f377da19e4d98dd598b26a9a0cded4a9c993b1bf59718aa539d3a581a16

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.1.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
