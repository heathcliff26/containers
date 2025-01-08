FROM docker.io/library/node:lts-bookworm@sha256:1f097426a7ddd1c5d0eacfe0402fdf91e38e4ecc37d23780428f6b87145ad2aa

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.12.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
