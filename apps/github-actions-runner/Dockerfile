FROM mcr.microsoft.com/dotnet/runtime-deps:8.0.18-jammy@sha256:df97e9d77bfec263c7f8bd7ca14f47009f95d3c9b3f20536b9234ea0312eca2d

# renovate: datasource=github-releases depName=actions/runner extractVersion=^v(?<version>.*)$
ARG RUNNER_VERSION=2.327.1
# renovate: datasource=github-releases depName=actions/runner-container-hooks extractVersion=^v(?<version>.*)$
ARG RUNNER_CONTAINER_HOOKS_VERSION=0.7.0
# renovate: datasource=github-tags depName=docker/cli extractVersion=^v(?<version>.*)$
ARG DOCKER_VERSION=28.3.3

ARG TARGETPLATFORM
ENV TARGETPLATFORM=${TARGETPLATFORM:-linux/amd64}

ENV DEBIAN_FRONTEND=noninteractive
ENV LANG=en_US.UTF-8
ENV RUNNER_MANUALLY_TRAP_SIG=1
ENV ACTIONS_RUNNER_PRINT_LOG_TO_STDOUT=1

RUN \
    # Prerequisites
    set -eux \
    && apt-get -qq update \
    && apt-get install -y \
        bash \
        ca-certificates \
        curl \
        git \
        jo \
        jq \
        locales \
        moreutils \
        sudo \
        tzdata \
        unrar \
        unzip \
        wget \
    && localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8 \
    && \
    case "${TARGETPLATFORM}" in \
        'linux/amd64') export ARCH='x64' && export DOCKER_ARCH='x86_64' ;; \
        'linux/arm64') export ARCH='arm64' && export DOCKER_ARCH='aarch64' ;; \
    esac \
    && \
    # Prepare user
    adduser --disabled-password --gecos "" --uid 1001 runner \
    && groupadd docker --gid 123 \
    && usermod -aG sudo runner \
    && usermod -aG docker runner \
    && echo "%sudo   ALL=(ALL:ALL) NOPASSWD:ALL" > /etc/sudoers \
    && echo "Defaults env_keep += \"DEBIAN_FRONTEND\"" >> /etc/sudoers \
    && \
    # Install Runner
    cd /home/runner \
    && \
    curl -fsSL -o runner.tar.gz "https://github.com/actions/runner/releases/download/v${RUNNER_VERSION}/actions-runner-linux-${ARCH}-${RUNNER_VERSION}.tar.gz" \
        && tar xzf ./runner.tar.gz \
        && rm runner.tar.gz \
    && \
    curl -fsSL -o runner-container-hooks.zip https://github.com/actions/runner-container-hooks/releases/download/v${RUNNER_CONTAINER_HOOKS_VERSION}/actions-runner-hooks-k8s-${RUNNER_CONTAINER_HOOKS_VERSION}.zip \
        && unzip ./runner-container-hooks.zip -d ./k8s \
        && rm runner-container-hooks.zip \
    && \
    # Install Docker
    curl -fsSL -o docker.tgz https://download.docker.com/linux/static/stable/${DOCKER_ARCH}/docker-${DOCKER_VERSION}.tgz \
        && tar zxvf docker.tgz \
        && rm -rf docker.tgz \
        && chown -R runner:docker . \
        && install -o root -g root -m 755 docker/* /usr/bin/ \
        && rm -rf docker \
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

WORKDIR /home/runner
USER runner
CMD ["/home/runner/run.sh"]
