FROM docker.io/library/node:lts-bookworm@sha256:d764525456dfe2f96a436ba00f864ee8ae3690bfb457c9f12a3a2a35b2d8be41

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2023.9.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
