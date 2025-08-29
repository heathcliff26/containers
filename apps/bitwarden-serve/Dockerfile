FROM docker.io/library/node:lts-bookworm@sha256:58f97e9f835c1e48fe950bca42fc7336a90e56bd0e209fa95c9552f97025de2a

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.8.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
