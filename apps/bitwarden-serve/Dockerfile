FROM docker.io/library/node:lts-bookworm@sha256:0d2dae60fe3ad12418e7e2c994f356065554b44824b3461c80d8b9ea645ba1ac

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.2.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
