FROM docker.io/library/node:lts-bookworm@sha256:f1f8564f958d623629d96ac170e148057a91cc24f5fe258a0ecdf62286582ff3

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.10.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
