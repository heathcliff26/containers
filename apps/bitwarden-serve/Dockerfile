FROM docker.io/library/node:lts-bookworm@sha256:af65937f2a7e69098ba4126188aaa70834dfe24c4a0761372bfa9cde1af9248f

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.9.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
