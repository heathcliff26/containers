FROM docker.io/library/node:lts-bookworm@sha256:6326b52a508f0d99ffdbfaa29a69380321b215153db6f32974835bac71b38fa4

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
