FROM docker.io/library/node:lts-bookworm@sha256:24da5a5fcd6a01a3d8e2bc0f8401f242850e80a5a4395d56e99d60a469853c22

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.8.2

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
