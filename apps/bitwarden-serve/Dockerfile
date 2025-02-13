FROM docker.io/library/node:lts-bookworm@sha256:7c6b02a11eaa9e1fcaaac3fa50d26c19db87be4640e5be41d9175cf6ff493436

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.1.3

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
