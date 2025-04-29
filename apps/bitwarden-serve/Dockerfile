FROM docker.io/library/node:lts-bookworm@sha256:f57e74d3ded287ee0fca136fd04f0c21c3f5e3926d48fbf9d6b17a69536c4397

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.3.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
