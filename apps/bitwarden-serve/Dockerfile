FROM docker.io/library/node:lts-bookworm@sha256:1b0c1d80fce6d8816c7d3a870e2100e7ddaf8617ac9810d7be2c97633377910d

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.12.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
