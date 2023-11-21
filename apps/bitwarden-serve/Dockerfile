FROM docker.io/library/node:lts-bookworm@sha256:715ecafe1d62e5d452f7945c32f81733425aba550e40dbf2ba5b9596e73bdb92

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.10.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
