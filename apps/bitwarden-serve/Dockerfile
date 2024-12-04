FROM docker.io/library/node:lts-bookworm@sha256:ebf4cb185bca60e3d33a3a9d7df436280ebd38ae715ee6e7e82271f2dee04434

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.11.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
