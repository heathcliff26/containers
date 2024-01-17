FROM docker.io/library/node:lts-bookworm@sha256:ffebb4405810c92d267a764b21975fb2d96772e41877248a37bf3abaa0d3b590

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.1.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
