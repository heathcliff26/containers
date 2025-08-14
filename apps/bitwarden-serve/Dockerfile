FROM docker.io/library/node:lts-bookworm@sha256:73297e219c73d998158945abd3dbc648c2756d610191da1498b2749c208614dd

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.7.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
