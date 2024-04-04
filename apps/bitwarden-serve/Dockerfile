FROM docker.io/library/node:lts-bookworm@sha256:fe75ec87317c6ead7a4d43e0a8eaeaadf6f8085b706ab547261cbc73c03403fe

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.2.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
