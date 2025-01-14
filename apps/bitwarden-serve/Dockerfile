FROM docker.io/library/node:lts-bookworm@sha256:816f04d578545be8f3faadaefaa0926c65e67056d9bf2864009976380c2b0713

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.12.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
