FROM docker.io/library/ubuntu:26.04@sha256:651ba3fe3a830441e3deaf70fafac40d808a6bd2800a6f2c43130055159f23e6

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
