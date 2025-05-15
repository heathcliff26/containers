FROM docker.io/library/node:lts-bookworm@sha256:a9316925500a0ceef7510f5587cd7147a969cc5c41d0cce2c11affaf65a57ae0

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.4.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
