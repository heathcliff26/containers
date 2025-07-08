FROM docker.io/library/node:lts-bookworm@sha256:2fa6c977460b56d4d8278947ab56faeb312bc4cc6c4cf78920c6de27812f51c5

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
