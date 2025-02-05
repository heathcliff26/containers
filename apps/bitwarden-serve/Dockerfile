FROM docker.io/library/node:lts-bookworm@sha256:3962f5a3aa1ebc1d82a1897a4a5f2db9ebdf2233edf521bdd7bdfb754284b366

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.1.2

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
