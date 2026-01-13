FROM docker.io/library/node:lts-bookworm@sha256:929c026d5a4e4a59685b3c1dbc1a8c3eb090aa95373d3a4fd668daa2493c8331

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Bitwarden cli installed via node" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/bitwarden-serve" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="bitwarden-serve"

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.12.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
