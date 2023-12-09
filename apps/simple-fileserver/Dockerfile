###############################################################################
# BEGIN build-stage
# Compile the binary
FROM --platform=$BUILDPLATFORM docker.io/library/golang:1.21.5@sha256:63769aa11f833b6fcd4ed5a2122a07507c92b5d832aed4b3c19ae85d595bc739 AS build-stage

ARG BUILDPLATFORM
ARG TARGETARCH

WORKDIR /app

COPY vendor ./vendor
COPY go.mod go.sum *.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH="${TARGETARCH}" go build -ldflags="-w -s" -o /simple-fileserver

#
# END build-stage
###############################################################################

###############################################################################
# BEGIN test-stage
# Run the tests in the container
FROM build-stage AS test-stage

COPY testdata ./testdata

RUN go test -v ./...

#
# END test-stage
###############################################################################

###############################################################################
# BEGIN final-stage
# Create final docker image
FROM scratch AS final-stage

WORKDIR /

COPY --from=build-stage /simple-fileserver /

EXPOSE 8080

USER 1001

ENTRYPOINT ["/simple-fileserver", "-webroot", "/webroot"]

#
# END final-stage
###############################################################################
