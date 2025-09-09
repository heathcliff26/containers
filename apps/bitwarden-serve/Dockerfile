FROM docker.io/library/node:lts-bookworm@sha256:dc081610647ddc0010be25151ad49eac522e4d3a9bc0d614be8ff8b68c7c9654

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.8.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
