FROM docker.io/library/node:lts-bookworm@sha256:e515259afd26f60db74957c62203c93d45760f2ba864d94accfa2edfc1ac17cf

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
