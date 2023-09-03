FROM node:lts-bookworm@sha256:11e9c297fc51f6f65f7d0c7c8a8581e5721f2f16de43ceff1a199fd3ef609f95

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.8.2

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
