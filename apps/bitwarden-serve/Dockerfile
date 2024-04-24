FROM docker.io/library/node:lts-bookworm@sha256:ec31b4e7762901e0a1cbf5a82ffd66c9798057c50f1b6f33a0089cf0bbbcffed

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.3.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
