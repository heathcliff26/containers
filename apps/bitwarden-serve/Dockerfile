FROM docker.io/library/node:lts-bookworm@sha256:33bfe51ff32a17d6cfb5bca5ae2743a9b2390ac99dcffde543565bb83b1fb786

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.1.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
