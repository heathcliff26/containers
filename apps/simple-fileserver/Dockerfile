###############################################################################
# BEGIN build-stage
# Compile the binary
FROM --platform=$BUILDPLATFORM docker.io/library/golang:1.21.5@sha256:ae34fbf671566a533f92e5469f3f3d34e9e6fb14c826db09956454da9a84c9a9 AS build-stage

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
