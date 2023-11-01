FROM docker.io/library/node:lts-bookworm@sha256:ff339f04927bd81eafbca8592f96c45cb97e58a5c58ae1113cd7325ef016e96f

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.9.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
