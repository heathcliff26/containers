FROM docker.io/library/node:lts-bookworm@sha256:f00d0bb31d4a4c4212625884fceb52488f10fa60ba5a128f3e3e2d65fb550094

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.9.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
