FROM docker.io/library/node:lts-bookworm@sha256:11dff346595c72894fe1166b08fbc9f7af40bb913f036724cb33e5c4ae6dc6e1

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.9.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
