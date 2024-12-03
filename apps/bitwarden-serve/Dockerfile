FROM docker.io/library/node:lts-bookworm@sha256:fd453a224f4461243f1b631f8874b0b9d5e5c4e84ab2441f00d8f09aa3525fce

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.11.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
