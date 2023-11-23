FROM docker.io/library/node:lts-bookworm@sha256:36e29862d0593675aa0f99b22c3c0f1165a24731b6ccb1f317f181fe96321b64

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.10.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
