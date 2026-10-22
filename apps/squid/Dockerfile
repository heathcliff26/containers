FROM docker.io/library/alpine:3.24.3@sha256:b3413572228e34bd1c9245f5565e6cd127005655e8a7899aef33d8e83255c9e6

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Running squid proxy in a lightweight alpine container" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/squid" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="squid"

RUN apk add --no-cache squid

EXPOSE 3128
VOLUME /var/log/squid \
	/var/spool/squid

COPY entrypoint.sh /usr/local/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
CMD ["-f", "/etc/squid/squid.conf", "-NYC"]
