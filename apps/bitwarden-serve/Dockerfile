FROM docker.io/library/node:lts-bookworm@sha256:f496dba5f685ef33797ed5882b4ce209053db67f88b50c1484ecccba6531bfde

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.10.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
