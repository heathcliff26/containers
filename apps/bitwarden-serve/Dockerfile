FROM docker.io/library/node:lts-bookworm@sha256:71972a1c6569a4b2e8b555d1558a5b08993142b110f93e16c2729feac6aa27f6

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.12.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
