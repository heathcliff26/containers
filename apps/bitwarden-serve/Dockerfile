FROM docker.io/library/node:lts-bookworm@sha256:4bc7ea5c063e6879b1888775595f48269d125b5f4ecaa6dfe6927397360938e8

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.8.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
