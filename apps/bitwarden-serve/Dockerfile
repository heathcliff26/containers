FROM docker.io/library/node:lts-bookworm@sha256:b6e7fef05c7609784c320e5c746aaf8d49221ec0eb3c1242fb673099511b0583

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.4.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
