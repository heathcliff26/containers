FROM docker.io/library/node:lts-bookworm@sha256:48db4f6ea21d134be744207225753a1730c4bc1b4cdf836d44511c36bf0e34d7

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.9.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
