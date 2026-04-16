FROM docker.io/library/ubuntu:24.04@sha256:c4a8d5503dfb2a3eb8ab5f807da5bc69a85730fb49b5cfca2330194ebcc41c7b

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Running tang server in a lightweight ubuntu container" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/tang" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="tang"

RUN apt-get update && \
    apt-get install -y tang && \
    apt-get clean && \
    rm -rf \
        /tmp/* \
        /var/lib/apt/lists/* \
        /var/cache/apt/* \
        /var/tmp/*

EXPOSE 8080
VOLUME ["/var/db/tang"]

ENTRYPOINT [ "/usr/libexec/tangd", "-p", "8080", "-l", "/var/db/tang" ]
