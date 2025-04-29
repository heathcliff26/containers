FROM docker.io/library/node:lts-bookworm@sha256:c9397a508192cf0c591e1923fe88d7dad25c586cf5fa73ce4b8889427747f230

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.3.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
