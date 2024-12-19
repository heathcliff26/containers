FROM docker.io/library/ubuntu:24.04

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
