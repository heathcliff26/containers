FROM docker.io/library/node:lts-bookworm@sha256:c3ef15af9be4505fde55589eadf42b4757a91e6b1b3be796bdec0f86560205e9

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.1.3

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
