FROM docker.io/library/node:lts-bookworm@sha256:012715bdfd7227cc531be9c1190daf48eaff8d33431e9a46762adbde8ce9764b

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.3.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
