FROM docker.io/library/node:lts-bookworm@sha256:854735b7b4dd7e70dced63752542346d42068b5a5a05120b3a4699be59dc0535

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.10.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
