FROM docker.io/library/node:lts-bookworm@sha256:17ed4afc1466aefe91add68c6c585f8f688ed0f2494b355269f038b522d784c6

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.9.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
