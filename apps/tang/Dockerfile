FROM docker.io/library/ubuntu:24.04@sha256:353675e2a41babd526e2b837d7ec780c2a05bca0164f7ea5dbbd433d21d166fc

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
