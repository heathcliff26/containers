FROM docker.io/library/node:lts-bookworm@sha256:83355e9c59b3201495ef828a92feb8b16f57f50f5ac125781f17439f12c62e82

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.2.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
