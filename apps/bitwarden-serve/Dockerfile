FROM docker.io/library/node:lts-bookworm@sha256:d38bdc98e1be70113c21f5c1b3069b30a75cfd2f494c5c82454b0fd2e7955809

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.7.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
