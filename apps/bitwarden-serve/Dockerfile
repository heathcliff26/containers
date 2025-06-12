FROM docker.io/library/node:lts-bookworm@sha256:2040569338f7a279161576ec5ff332d97ef53f46baf94c85fe06fd3ed81a8f5f

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.5.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
