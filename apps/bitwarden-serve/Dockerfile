FROM docker.io/library/node:lts-bookworm@sha256:48921ca7e423bc46c9a6a058cf3abec92f34765150e43cc7adc3d9fa172ff25a

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.9.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
