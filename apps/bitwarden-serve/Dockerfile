FROM docker.io/library/node:lts-bookworm@sha256:f0550f49012688961e8a30a0baa36836495b45fd8c283e9496dc8871ee78ebd1

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.4.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
