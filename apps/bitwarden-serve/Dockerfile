FROM docker.io/library/node:lts-bookworm@sha256:a4d1de4c7339eabcf78a90137dfd551b798829e3ef3e399e0036ac454afa1291

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.8.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
