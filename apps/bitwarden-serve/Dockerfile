FROM docker.io/library/node:lts-bookworm@sha256:4a126f3116c37fbd8583209f13518efa9e9f6efc6bf18917141c6e6d4068fea1

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.3.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
