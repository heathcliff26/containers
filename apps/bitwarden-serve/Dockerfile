FROM docker.io/library/node:lts-bookworm@sha256:74066d050011dd18c56390523cb255b38bce09b332221a4ec15ab0d8e1ac61fe

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.4.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
