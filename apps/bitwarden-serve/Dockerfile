FROM docker.io/library/node:lts-bookworm@sha256:d7d3eb5b608bd14c2edf4a72709772fe194d5167d8e132377461197b256ee3d5

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
