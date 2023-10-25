FROM docker.io/library/node:lts-bookworm@sha256:76985c6cb58f5679c724ee3deacf77a8afb86ee47ab12d871a2b6cc363822794

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.9.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
