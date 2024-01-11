FROM docker.io/library/node:lts-bookworm@sha256:9aa3de5470c99408fda002dc1f406e92a31daf0492eb33d857d8d9d252edcc52

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.1.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
