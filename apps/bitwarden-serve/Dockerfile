FROM docker.io/library/node:lts-bookworm@sha256:cbdfb793208fd6e63664b4d4e1a539a8b69e7c1f69c49dd55ca6e016074744fd

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.2.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
