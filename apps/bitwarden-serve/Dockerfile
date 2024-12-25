FROM docker.io/library/node:lts-bookworm@sha256:7bea049c66b5846c4ce2786b1b4e32865ef11b10fa446c1bfd791daea412c299

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.12.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
