FROM docker.io/library/node:lts-bookworm@sha256:60ce62a127304f79d83551fe29a563d04375dae47ae2a45cf463759b5fde2b9f

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.10.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
