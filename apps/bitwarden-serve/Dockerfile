FROM docker.io/library/node:lts-bookworm@sha256:146bbe4eaee99ae885be2a0a767f63a4b96032141d70d80e590f10e0d7ebabcb

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.10.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
