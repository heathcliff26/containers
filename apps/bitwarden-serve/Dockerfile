FROM docker.io/library/node:lts-bookworm@sha256:445acd9b2ef7e9de665424053bf95652e0b8995ef36500557d48faf29300170a

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.12.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
