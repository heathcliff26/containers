FROM docker.io/library/alpine:3.22.0@sha256:a8bf99cad4d74d1ec021670fa473e6ff716e401b0542918ce68e94c3234910e1

RUN apk add --no-cache squid

EXPOSE 3128
VOLUME /var/log/squid \
	/var/spool/squid

COPY entrypoint.sh /usr/local/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
CMD ["-f", "/etc/squid/squid.conf", "-NYC"]
