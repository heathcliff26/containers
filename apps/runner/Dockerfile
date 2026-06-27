###############################################################################
# BEGIN cosign-bin
# Compile the binary
FROM ghcr.io/sigstore/cosign/cosign:v2.6.1 AS cosign-bin

#
# END cosign-bin
###############################################################################

###############################################################################
# BEGIN runner-image
# Build the runner image
FROM docker.io/library/debian:13.5@sha256:d07d1b51c39f51188e60be9b64e6bf769fa94e187f092bc32b91305cfa34ba5a

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Basic CI image used for tekton pipeline tasks" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/runner" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="runner"

ARG TARGETARCH
# renovate: datasource=github-releases depName=mikefarah/yq extractVersion=^(?<version>.*)$
ENV YQ_VERSION=v4.53.3
# renovate: datasource=github-releases depName=fluxcd/flux2 extractVersion=^v(?<version>.*)$
ENV FLUX_VERSION=2.8.8
# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^(?<version>.*)$
ENV KUBERNETES_VERSION=v1.36.2

USER root

# Install Prerequisites
RUN set -eux \
    && apt-get -qq update \
    && apt-get -qq install -y --no-install-recommends --no-install-suggests \
        ca-certificates \
        curl \
        gcc \
        make \
        moreutils \
        wget \
        zstd \
        gettext-base \
        sudo \
        python3 \
        python3-pip \
        python3-yaml \
        podman \
        rootlesskit \
        qemu-user-static \
        binfmt-support \
    && \
    # Cleanup
    apt-get purge -y --auto-remove -o APT::AutoRemove::RecommendsImportant=false \
    && apt-get autoremove -y \
    && apt-get clean \
    && \
    rm -rf \
        /tmp/* \
        /var/lib/apt/lists/* \
        /var/cache/apt/* \
        /var/tmp/*

VOLUME /var/lib/containers

# Install yq
RUN curl -fsSL -o /usr/local/bin/yq "https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_${TARGETARCH}" \
        && chmod +x /usr/local/bin/yq

# Install flux
RUN curl -fsSL -o /tmp/flux.tar.gz "https://github.com/fluxcd/flux2/releases/download/v${FLUX_VERSION}/flux_${FLUX_VERSION}_linux_${TARGETARCH}.tar.gz" \
        && tar -xzf /tmp/flux.tar.gz -C /usr/local/bin flux \
        && rm -rf /tmp/flux.tar.gz

# Install kubectl
RUN curl -fsSL -o /usr/local/bin/kubectl "https://dl.k8s.io/release/${KUBERNETES_VERSION}/bin/linux/${TARGETARCH}/kubectl" \
        && chmod +x /usr/local/bin/kubectl

# Install cosign
COPY --from=cosign-bin /ko-app/cosign /usr/local/bin/cosign

# Install scripts
COPY crd-extractor.sh openapi2jsonschema.py /scripts/

# Add runner user
RUN useradd -ms /bin/bash runner \
    && usermod -aG sudo runner \
    && echo "runner ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/runner
USER runner
WORKDIR /home/runner

#
# END runner-img
###############################################################################
