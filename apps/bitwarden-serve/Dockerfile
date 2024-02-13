FROM docker.io/library/node:lts-bookworm@sha256:489f37b8c21a3210adf4b2b42fbabd14bd8d30853a0a42194fecfaca2e3b0e78

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.2.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
