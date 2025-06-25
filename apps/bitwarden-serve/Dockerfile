FROM docker.io/library/node:lts-bookworm@sha256:0c0734eb7051babbb3e95cd74e684f940552b31472152edf0bb23e54ab44a0d7

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.5.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
