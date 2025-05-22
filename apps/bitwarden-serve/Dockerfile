FROM docker.io/library/node:lts-bookworm@sha256:f466376b5f5ca8327da5914fddbcb5eee21c1c8482bed548a5ee84750d0b1b69

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.4.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
