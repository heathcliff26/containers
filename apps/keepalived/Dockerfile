FROM docker.io/library/alpine:3.20.1@sha256:ff5265e55d2f71d89d17ee63a634e37686637d2e2c8e76e57837e010c8666904

RUN apk add --no-cache keepalived && \
	rm -rf /etc/keepalived && \
    adduser -D keepalived_script

ENTRYPOINT [ "/usr/sbin/keepalived", "--dont-fork", "--log-console" ]
