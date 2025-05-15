FROM docker.io/library/node:lts-bookworm@sha256:909d2840362f35b605ebf4176952b06d02fdcf32971213e5d5ea1090a41097c8

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.4.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
