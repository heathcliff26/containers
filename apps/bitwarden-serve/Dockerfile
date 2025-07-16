FROM docker.io/library/node:lts-bookworm@sha256:d4191686ddd8a87498fe948b2b75b65582fe0ef6659fe12ced60800cd55e4df4

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
