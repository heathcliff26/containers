FROM docker.io/library/node:lts-bookworm@sha256:d6ba961be11a08d00369a01dec7566fc8f2dd895a8edd5b983915123f0e81c63

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.8.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
