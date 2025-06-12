FROM docker.io/library/node:lts-bookworm@sha256:68cf33c392726bdda3f0bc5e53894ebd641fed6367fb0b4204277bbd2377a543

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.5.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
