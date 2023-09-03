FROM docker.io/library/ubuntu:22.04

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION="cli-v2023.7.0"

RUN apt update && \
    apt install -y wget unzip && \
    wget https://github.com/bitwarden/clients/releases/download/${BW_CLI_VERSION}/bw-linux-${BW_CLI_VERSION#cli-v}.zip && \
    unzip bw-linux-${BW_CLI_VERSION#cli-v}.zip && \
    chmod +x bw && \
    mv bw /usr/local/bin/bw && \
    rm -rfv *.zip

COPY entrypoint.sh /

USER 1000

CMD ["/entrypoint.sh"]
