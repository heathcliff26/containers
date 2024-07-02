FROM docker.io/library/node:lts-bookworm@sha256:940071b4e4c9918c40c632a4b51861aded3087696a2847a5d8ecff3546eb13ce

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
