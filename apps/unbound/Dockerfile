FROM docker.io/library/alpine:3.22.0@sha256:a8bf99cad4d74d1ec021670fa473e6ff716e401b0542918ce68e94c3234910e1

RUN apk add --no-cache unbound && \
	rm -f /etc/unbound/unbound.conf

COPY entrypoint.sh /usr/local/bin

WORKDIR /etc/unbound

ENTRYPOINT [ "entrypoint.sh" ]
