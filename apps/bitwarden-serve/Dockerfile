FROM docker.io/library/node:lts-bookworm@sha256:a12ce152ab63dbeb232fcbcd0e9fdfbee6e6443d57785915967893ce9e7c58ba

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
