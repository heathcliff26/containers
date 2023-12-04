###############################################################################
# BEGIN build-stage
# Compile the binary
FROM --platform=$BUILDPLATFORM docker.io/library/golang:1.21.4@sha256:9baee0edab4139ae9b108fffabb8e2e98a67f0b259fd25283c2a084bd74fea0d AS build-stage

ARG BUILDPLATFORM
ARG TARGETARCH

WORKDIR /app

COPY vendor ./vendor
COPY go.mod go.sum *.go ./
COPY collector ./collector
COPY speedtest ./speedtest

RUN CGO_ENABLED=0 GOOS=linux GOARCH="${TARGETARCH}" go build -ldflags="-w -s" -o /speedtest-exporter

#
# END build-stage
###############################################################################

###############################################################################
# BEGIN test-stage
# Run the tests in the container
FROM build-stage AS test-stage

RUN go test -v ./...

#
# END test-stage
###############################################################################

###############################################################################
# BEGIN combine-stage
# Combine all outputs, to enable single layer copy for the final image
FROM scratch AS combine-stage

COPY --from=build-stage /speedtest-exporter /

COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

#
# END combine-stage
###############################################################################

###############################################################################
# BEGIN final-stage
# Create final docker image
FROM scratch AS final-stage

WORKDIR /

COPY --from=combine-stage / /

EXPOSE 8080

USER 1001

ENTRYPOINT ["/speedtest-exporter"]

#
# END final-stage
###############################################################################
