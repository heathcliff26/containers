FROM docker.io/library/node:lts-bookworm@sha256:2094ac60fa9f3df66c87f93e3ec4ab1e296f0a4d2ba3eed3f2d367529a3118b1

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.1.3

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
