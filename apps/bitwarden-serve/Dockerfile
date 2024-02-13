FROM docker.io/library/node:lts-bookworm@sha256:75f2e314ac134e3e9973920c9debb24b3ce8b53d4e5de4a2421480c6c9c19fe8

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.2.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
