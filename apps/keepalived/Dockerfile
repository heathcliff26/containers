FROM docker.io/library/alpine:3.22.0@sha256:a8bf99cad4d74d1ec021670fa473e6ff716e401b0542918ce68e94c3234910e1

RUN apk add --no-cache keepalived && \
	rm -rf /etc/keepalived && \
    adduser -D keepalived_script

ENTRYPOINT [ "/usr/sbin/keepalived", "--dont-fork", "--log-console" ]
