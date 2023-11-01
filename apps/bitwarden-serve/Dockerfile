FROM docker.io/library/node:lts-bookworm@sha256:be4ad7f6808f747a611a374a786bf625382592310d47e6014004a97b473bdd4e

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.9.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
