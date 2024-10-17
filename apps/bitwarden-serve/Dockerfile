FROM docker.io/library/node:lts-bookworm@sha256:ed4e1d34925491770279e4d8023ce9355fb93dd4d10b433ba3af6ed53276bba6

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.9.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
