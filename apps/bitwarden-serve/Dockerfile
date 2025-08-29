FROM docker.io/library/node:lts-bookworm@sha256:9f1cfc5577021d71b9eeb0bb6e0820cb37ca858146ba78f7158a9c8fe2faf412

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.8.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
