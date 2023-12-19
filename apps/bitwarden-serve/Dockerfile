FROM docker.io/library/node:lts-bookworm@sha256:3155b7fb5fd482060da75afc881857fad539b4a6e3ac9de7aab739de3145005a

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.12.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
