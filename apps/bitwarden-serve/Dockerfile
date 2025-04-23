FROM docker.io/library/node:lts-bookworm@sha256:853006d5c7835ed1247944b0f221e4add0c55b365a933c2c507ce2f004e322ef

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.3.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
