FROM docker.io/library/node:lts-bookworm@sha256:87ec76e4bb2ae5b7e2f10af227d98a8fd244c78cbfc53fa2985e30d9dec7fc6e

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.3.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
