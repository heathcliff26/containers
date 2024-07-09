FROM docker.io/library/node:lts-bookworm@sha256:26aa76c4aaf39c2da2e679941ee1589243afb46173c861aa0d77d3a34fe2488c

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
