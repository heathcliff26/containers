FROM docker.io/library/node:lts-bookworm@sha256:572a0114234c8142cef0814ebbd02ed5283d7ba5308156e353b4195545ba4ffe

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.1.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
