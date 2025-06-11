FROM docker.io/library/node:lts-bookworm@sha256:f627d0e594bbc713de73c2c50d18033c89e114cb1645c3b520b201426cfe1975

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.5.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
