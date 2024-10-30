FROM docker.io/library/node:lts-bookworm@sha256:de4c8be8232b7081d8846360d73ce6dbff33c6636f2259cd14d82c0de1ac3ff2

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.10.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
