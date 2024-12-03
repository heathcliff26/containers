FROM docker.io/library/node:lts-bookworm@sha256:cb244536f6047c0057f1c6d90ae404880155c49b78fe2e8f25ef90c4fa7127a6

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.11.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
