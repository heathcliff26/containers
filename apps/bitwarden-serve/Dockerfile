FROM docker.io/library/node:lts-bookworm@sha256:cd7fa8f136023f7500490e410ba70dd3982ccca21805264f2a260a3a97be7376

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.9.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
