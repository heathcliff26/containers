FROM docker.io/library/node:lts-bookworm@sha256:0894ea26b282cfa4c2ed20bb3882e67be74f09b909b8e4c4bb0b0ab9e7f15c59

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.9.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
