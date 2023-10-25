FROM docker.io/library/node:lts-bookworm@sha256:4a66e17d2cee338902c72d40a4a9cb5e00fa42a78d613073ba64440aefa747de

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.9.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
