FROM docker.io/library/node:lts-bookworm@sha256:d6925dc84f8c0d1c1f8df4ea6a9a54e57d430241cb734b1b0c45ed6d26e8e9c0

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.4.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
