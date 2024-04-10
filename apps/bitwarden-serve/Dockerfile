FROM docker.io/library/node:lts-bookworm@sha256:cad16010fccab8c3697b5f3577953c2d9a396e9a205673598cd11264ee6d0581

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.3.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
