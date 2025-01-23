FROM docker.io/library/node:lts-bookworm@sha256:ae2f3d4cc65d251352eca01ba668824f651a2ee4d2a37e2efb22649521a483fd

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.1.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
