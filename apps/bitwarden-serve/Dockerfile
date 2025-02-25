FROM docker.io/library/node:lts-bookworm@sha256:a27967171a54d13b85387675fe0362f77e969c7a8abad2faaee45490ad48fd52

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.1.3

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
