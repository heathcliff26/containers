FROM docker.io/library/node:lts-bookworm@sha256:18ad6b547976c931c7374c2f7b0a1dd27eda7b55dbf41c9cdb35fe4081996652

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.9.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
