FROM docker.io/library/node:lts-bookworm@sha256:e9df36b6e7e4c66e77ee132a63fd4d11740418fcad43e419b444e9165d34fd2b

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
