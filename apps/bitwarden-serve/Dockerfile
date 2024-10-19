FROM docker.io/library/node:lts-bookworm@sha256:d7c25dbb3da3f8d9990073780e3c1eb0f0c68a55ef5adf402aa96fb4ee49f6e1

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.9.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
