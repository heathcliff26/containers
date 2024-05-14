FROM docker.io/library/node:lts-bookworm@sha256:32834b043faec33925138c0f8fe5d6d2b2ba68eb14eb59e4f994dcadc95acf32

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.4.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
