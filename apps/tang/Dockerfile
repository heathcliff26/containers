FROM docker.io/library/ubuntu:24.04@sha256:590e57acc18d58cd25d00254d4ca989bbfcd7d929ca6b521892c9c904c391f50

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
