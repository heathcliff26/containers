FROM docker.io/library/node:lts-bookworm@sha256:73ff989b1ae110bbdd3e5221e1bd110f4212a9e9ad40dc6ac3f6411dbe12fc59

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
