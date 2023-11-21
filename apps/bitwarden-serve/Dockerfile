FROM docker.io/library/node:lts-bookworm@sha256:3ae44db2fe4e00eb6845973554fd5b80b6ec76c78792d025f7c63008a9bcaa27

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.10.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
