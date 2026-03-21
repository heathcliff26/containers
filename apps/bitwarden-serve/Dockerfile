FROM docker.io/library/node:lts-bookworm@sha256:5a593d74b632d1c6f816457477b6819760e13624455d587eef0fa418c8d0777b

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Bitwarden cli installed via node" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/bitwarden-serve" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="bitwarden-serve"

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2026.2.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER nobody:nogroup

CMD ["/entrypoint.sh"]
