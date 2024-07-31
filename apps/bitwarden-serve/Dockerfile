FROM docker.io/library/node:lts-bookworm@sha256:1ae9ba874435551280e95c8a8e74adf8a48d72b564bf9dfe4718231f2144c88f

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.7.2

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
