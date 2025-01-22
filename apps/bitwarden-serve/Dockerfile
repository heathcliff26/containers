FROM docker.io/library/node:lts-bookworm@sha256:671f4ef36157cf7f0651324c4117dd298eddd796a538f21e35c85fcadc5e963d

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.1.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
