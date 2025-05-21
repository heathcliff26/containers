FROM docker.io/library/node:lts-bookworm@sha256:50730ab0b271c750c1a29bb4a60aff470e359a60a00fe40379c4747ac1e45796

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.4.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
