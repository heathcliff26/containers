###############################################################################
# BEGIN fetch-stage
# Fetch the binary
FROM --platform=$BUILDPLATFORM docker.io/library/ubuntu:24.04 AS fetch-stage

# renovate: datasource=github-releases depName=bitwarden/clients extractVersion=^cli-v(?<version>.*)$
ENV BW_CLI_VERSION="cli-v2026.2.0"

RUN apt-get update && \
    apt-get install -y curl unzip

ARG TARGETARCH

WORKDIR /fetch

RUN set -ex; \
    if [ "${TARGETARCH}" = "amd64" ]; then \
        PLATFORM_TUPLE="linux"; \
    elif [ "${TARGETARCH}" = "arm64" ]; then \
        PLATFORM_TUPLE="linux-arm64"; \
    else \
        echo "Unsupported architecture: ${TARGETARCH}"; \
        exit 1; \
    fi; \
    curl -SL -o bw.zip "https://github.com/bitwarden/clients/releases/download/${BW_CLI_VERSION}/bw-${PLATFORM_TUPLE}-${BW_CLI_VERSION#cli-v}.zip"

RUN unzip bw.zip

RUN chmod +x bw

#
# END fetch-stage
###############################################################################

###############################################################################
# BEGIN final-stage
# Create final docker image
FROM docker.io/library/ubuntu:24.04@sha256:186072bba1b2f436cbb91ef2567abca677337cfc786c86e107d25b7072feef0c AS final-stage

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Bitwarden cli in an ubuntu image" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/bitwarden-serve" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="bitwarden-serve"

COPY --from=fetch-stage /fetch/bw /usr/local/bin/bw

COPY entrypoint.sh /

ENV HOME=/tmp

USER nobody:nogroup

CMD ["/entrypoint.sh"]

#
# END final-stage
###############################################################################
