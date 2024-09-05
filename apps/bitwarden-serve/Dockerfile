FROM docker.io/library/node:lts-bookworm@sha256:e3de1da7f2cc6ca3065fdd013513b72791f10a6dbdfdc05563800b5318a309cc

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.8.2

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
