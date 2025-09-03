FROM docker.io/library/ubuntu:24.04@sha256:9cbed754112939e914291337b5e554b07ad7c392491dba6daf25eef1332a22e8

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
