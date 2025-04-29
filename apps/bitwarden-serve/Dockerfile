FROM docker.io/library/node:lts-bookworm@sha256:120a74c2ae5bf7f51ab253caef82e7ec3491ac85ef5b1c25751e6c3b55f49cf8

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.3.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
