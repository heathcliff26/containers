FROM docker.io/library/node:lts-bookworm@sha256:66f3e4c476ce6d27edf1c0d14c8cc6a7098d53329931e4a4a757bfa04b1604b1

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.9.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
