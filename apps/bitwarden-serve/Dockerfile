FROM docker.io/library/node:lts-bookworm@sha256:1de802fdf24ea1b1b5dddb5bbd17049708e61226c9753879360d056298fd5c66

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION=2024.3.1

RUN npm install -g @bitwarden/cli@${BW_CLI_VERSION#cli-v}

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
