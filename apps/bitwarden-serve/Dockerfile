FROM docker.io/library/node:lts-bookworm@sha256:d9204281f05b497330fe26b29761e0445707dd9b4e484db73d491e01d47702c3

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.9.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
