FROM docker.io/library/node:lts-bookworm@sha256:d77c6457d4318c6ef537dbf9fe86f36bfb997d280c9f949e3a1d968cf841390d

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.1.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
