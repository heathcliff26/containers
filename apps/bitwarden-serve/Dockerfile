FROM docker.io/library/node:lts-bookworm@sha256:02cd2205818f121c13612721876f28c18bd50148bb8af532ea121c96ffcb59bf

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
