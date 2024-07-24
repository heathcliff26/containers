FROM docker.io/library/node:lts-bookworm@sha256:1c211295f822a971b07048f04583cac16ee4c8cb0310b3c6e148133a1923ce33

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
