FROM docker.io/library/node:lts-bookworm@sha256:7ca973caa50c822fb1f4ea11c6ae1e322daaf9b4d9bc1b56c51c2b22315c9866

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.2.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
