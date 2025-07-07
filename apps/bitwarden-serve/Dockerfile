FROM docker.io/library/node:lts-bookworm@sha256:df39165022a0463c1b9e5fe5198b326847a2cb22736ca2469bc7139785431673

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
