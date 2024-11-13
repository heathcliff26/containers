FROM docker.io/library/node:lts-bookworm@sha256:6eb1af33c8fc1104f23efdd15f1e791570395eab3d0fb597087d18a4ab1f7b65

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.10.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
