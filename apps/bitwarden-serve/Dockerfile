FROM docker.io/library/node:lts-bookworm@sha256:f172e4a7393cef2252c16f5f38d7abcf2cc0fd46c181c0ed5006cf3332e01173

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
