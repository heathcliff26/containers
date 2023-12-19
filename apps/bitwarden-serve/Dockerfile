FROM docker.io/library/node:lts-bookworm@sha256:2399607ee91a33ee9ef3302e5f194385a98d3f804e1d8baa6a5e314996fceb81

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.12.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
