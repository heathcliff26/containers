FROM docker.io/library/node:lts-bookworm@sha256:89b86535844dc9a7a2540b1cd3b3b1b84315fbc89369dc6d8f09591ca939562e

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.3.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
