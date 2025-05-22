FROM docker.io/library/node:lts-bookworm@sha256:bc77e2513f2034916f1f84299e88d17ff3612f45ad17655df07008ec6778a4bf

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2025.4.0

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
