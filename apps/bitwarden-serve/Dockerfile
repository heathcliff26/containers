FROM docker.io/library/node:lts-bookworm@sha256:263acfc82e6c90dc5f80a937868a7ede7f205f1e0a1eecd94b254e3497829dbb

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.7.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
