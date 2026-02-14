FROM ghcr.io/actions/actions-runner:2.331.0@sha256:dced476aa42703ebd9aafc295ce52f160989c4528e831fc3be2aef83a1b3f6da

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Self-hosted github actions runner" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/github-actions-runner" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="github-actions-runner"

ARG TARGETARCH
# renovate: datasource=github-releases depName=mikefarah/yq extractVersion=^(?<version>.*)$
ARG YQ_VERSION=v4.52.4
# renovate: datasource=github-releases depName=fluxcd/flux2 extractVersion=^v(?<version>.*)$
ARG FLUX_VERSION=2.7.5

USER root

# Install Prerequisites
RUN set -eux \
    && apt-get -qq update \
    && apt-get -qq install -y --no-install-recommends --no-install-suggests \
        ca-certificates \
        gcc \
        jo \
        make \
        moreutils \
        wget \
        zstd \
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


# Install yq
RUN curl -fsSL -o /usr/local/bin/yq "https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_${TARGETARCH}" \
        && chmod +x /usr/local/bin/yq

# Install flux
RUN curl -fsSL -o /tmp/flux.tar.gz "https://github.com/fluxcd/flux2/releases/download/v${FLUX_VERSION}/flux_${FLUX_VERSION}_linux_${TARGETARCH}.tar.gz" \
        && tar -xzf /tmp/flux.tar.gz -C /usr/local/bin flux \
        && rm -rf /tmp/flux.tar.gz

USER runner

CMD ["/home/runner/run.sh"]
