FROM docker.io/library/node:lts-bookworm@sha256:9a4a149d3c060538ccdecd568b236fb7d6d765b42699c1f17b63dfe77cadede8

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
