FROM docker.io/library/node:lts-bookworm@sha256:c85db7ec0c2cbedc73f417b890e0b08ecdf5316f858033316d6566ee11715158

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
