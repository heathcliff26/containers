FROM docker.io/library/node:lts-bookworm@sha256:ab71b9da5ba19445dc5bb76bf99c218941db2c4d70ff4de4e0d9ec90920bfe3f

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
