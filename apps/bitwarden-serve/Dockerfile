FROM docker.io/library/node:lts-bookworm@sha256:f7354aa281a8d45e44f4d28269200e1c412e9b41188599326a49506e01a89f97

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.12.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
