FROM docker.io/library/ubuntu:24.04

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Simple script to sync files with remote nodes via ssh from inside a kubernetes cluster" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/file-sync" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="file-sync"

RUN mkdir /root/.ssh && \
    apt-get update && \
    apt-get install -y yq rsync openssh-client netcat-traditional && \
    apt-get clean && \
    rm -rf \
        /tmp/* \
        /var/lib/apt/lists/* \
        /var/cache/apt/* \
        /var/tmp/*

COPY sync.sh /sync.sh

CMD [ "/sync.sh" ]
