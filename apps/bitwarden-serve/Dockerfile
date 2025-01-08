FROM docker.io/library/node:lts-bookworm@sha256:6781d62b7189e0dd9c2e17019caa25bd72b1a8cc9f9f8222aea2e856f51588fe

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.12.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
