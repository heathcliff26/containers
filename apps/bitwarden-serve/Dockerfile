FROM docker.io/library/node:lts-bookworm@sha256:933bcfad91e9052a02bc29eb5aa29033e542afac4174f9524b79066d97b23c24

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.8.2

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
