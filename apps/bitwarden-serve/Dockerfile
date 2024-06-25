FROM docker.io/library/node:lts-bookworm@sha256:611644c7d67a42ce9f99ede8a6ad70717eda00857d9c2a4e4718845e8b426f04

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
