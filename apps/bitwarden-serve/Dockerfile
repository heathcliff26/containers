FROM docker.io/library/node:lts-bookworm@sha256:3ab325a3e2b21e2289c4bfa11593a3da563d2a7559a4cb80493c7c0c9fe13cb6

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.1.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
