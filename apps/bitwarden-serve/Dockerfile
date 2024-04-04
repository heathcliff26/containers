FROM docker.io/library/node:lts-bookworm@sha256:6d7557fb008745f2cc1d9ed8965e04ab00a3087ca466ca80f17ea5676469d215

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.2.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
