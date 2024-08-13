FROM docker.io/library/node:lts-bookworm@sha256:d3c8ababe9566f9f3495d0d365a5c4b393f607924647dd52e75bf4f8a54effd3

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.7.2

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
