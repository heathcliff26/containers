FROM docker.io/library/node:lts-bookworm@sha256:77db97b595059d7f56643b81fd886ee3de634d6ea1ce5ca1200966ae601b086c

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.1.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
