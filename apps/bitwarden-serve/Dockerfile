FROM docker.io/library/node:lts-bookworm@sha256:c91101b285c01518fa4dd73e264b5c963f8f38bb4a4dcf450715ae775250791d

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.6.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
