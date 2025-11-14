FROM docker.io/library/ubuntu:24.04@sha256:4a5ffc07dd015f03d6cb56f209db27ed2378056f7cc8a70ae11d0c35718ed3e6

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
