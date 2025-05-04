FROM docker.io/library/golang:1.24.2@sha256:30baaea08c5d1e858329c50f29fe381e9b7d7bced11a0f5f1f69a1504cdfbf5e

# renovate: datasource=github-releases depName=golangci/golangci-lint
ARG GOLANGCI_LINT_VERSION=v2.1.6
# renovate: datasource=github-releases depName=ziglang/zig
ARG ZIG_VERSION=0.14.0
# renovate: datasource=github-releases depName=fyne-io/fyne
ARG FYNE_VERSION=v2.6.0
# renovate: datasource=github-releases depName=boxboat/fixuid
ARG FIXUID_VERSION=0.6.0

# Install tools
RUN set -eux; \
    apt-get update; \
    apt-get install -y -q --no-install-recommends \
        make \
        unzip \
        xz-utils \
        zip \
    ; \
    apt-get clean; \
    rm -r /var/lib/apt/lists/*;

# Add zig to path
ENV PATH=/usr/local/zig:$PATH

# Install Zig
RUN set -eux; \
    arch="$(dpkg --print-architecture)"; \
    url=; \
    sha256=; \
    case "$arch" in \
        'amd64') \
            # dev release
            # url="https://ziglang.org/builds/zig-linux-x86_64-${ZIG_VERSION}.tar.xz";\
            # stable release
            url="https://ziglang.org/download/${ZIG_VERSION}/zig-linux-x86_64-${ZIG_VERSION}.tar.xz";\
            ;; \
        'arm64') \
            # dev release
            # url="https://ziglang.org/builds/zig-linux-aarch64-${ZIG_VERSION}.tar.xz";\
            # stable release
            url="https://ziglang.org/download/${ZIG_VERSION}/zig-linux-aarch64-${ZIG_VERSION}.tar.xz";\
            ;; \
        *) echo >&2 "error: unsupported architecture '$arch'"; exit 1 ;; \
    esac; \
    curl -sSL ${url} -o zig.tar.xz; \
    tar -C /usr/local -Jxvf zig.tar.xz; \
    mv /usr/local/zig-* /usr/local/zig; \
    rm zig.tar.xz; \
    zig version;

# Install the fyne CLI tool
RUN set -eux; \
    go install -ldflags="-s" -v "fyne.io/fyne/v2/cmd/fyne@${FYNE_VERSION}"; \
    mv /go/bin/fyne /usr/local/bin/fyne; \
    fyne version; \
    go clean -cache -modcache; \
    mkdir -p "$GOPATH/pkg/mod" && chmod -R 777 "$GOPATH"

# Install the golangci-lint tool
RUN set -eux; \
    go install -ldflags="-s" -v "github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}"; \
    mv /go/bin/golangci-lint /usr/local/bin/golangci-lint; \
    golangci-lint version; \
    go clean -cache -modcache; \
    mkdir -p "$GOPATH/pkg/mod" && chmod -R 777 "$GOPATH"

# Install fixuid see #41
RUN arch="$(dpkg --print-architecture)"; \
    addgroup --gid 1000 docker; \
    adduser --uid 1000 --ingroup docker --home /home/docker --shell /bin/sh --disabled-password --gecos "" docker; \
    curl -SsL https://github.com/boxboat/fixuid/releases/download/v${FIXUID_VERSION}/fixuid-${FIXUID_VERSION}-linux-${arch}.tar.gz | tar -C /usr/local/bin -xzf -; \
    chown root:root /usr/local/bin/fixuid; \
    chmod 4755 /usr/local/bin/fixuid; \
    mkdir -p /etc/fixuid; \
    printf "user: docker\ngroup: docker\n" > /etc/fixuid/config.yml

# Install linux libraries
RUN set -eux; \
    dpkg --add-architecture arm64; \
    dpkg --add-architecture amd64; \
    apt-get update; \
    apt-get install -y -q --no-install-recommends \
        libgl-dev:amd64 \
        libx11-dev:amd64 \
        libxrandr-dev:amd64 \
        libxxf86vm-dev:amd64 \
        libxi-dev:amd64 \
        libxcursor-dev:amd64 \
        libxinerama-dev:amd64 \
         # deps to support wayland
        libxkbcommon-dev:amd64 \
        libdecor-0-dev:amd64 \
        libegl-dev:amd64 \
    ; \
    apt-get install -y -q --no-install-recommends \
        libgl-dev:arm64 \
        libx11-dev:arm64 \
        libxrandr-dev:arm64 \
        libxxf86vm-dev:arm64 \
        libxi-dev:arm64 \
        libxcursor-dev:arm64 \
        libxinerama-dev:arm64 \
         # deps to support wayland
        libxkbcommon-dev:arm64 \
        libdecor-0-dev:arm64 \
        libegl-dev:arm64 \
    ; \
    # remove static libX11 to allow zig build against shared X11 lib
    rm -rf /usr/lib/*/libX11.a; \
    apt-get -qy autoremove; \
    apt-get clean; \
    rm -r /var/lib/apt/lists/*;

WORKDIR /app
