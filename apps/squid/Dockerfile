FROM docker.io/library/alpine:3.23.4@sha256:c7989ac7a27b473e1795973c98d714f62b4dd0b134594d36880505ce0bfd716b

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
