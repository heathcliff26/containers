FROM docker.io/library/node:lts-bookworm@sha256:b616fd91b610b30f5bfc2b561035477bd7eae56a30e7e643e059bc313f99d319

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.9.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
