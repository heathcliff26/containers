FROM docker.io/library/node:lts-bookworm@sha256:615b1a42aa5c644b84cc9631c734aa1af1827d5c7096294f7d510a85277d87ae

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.2.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
