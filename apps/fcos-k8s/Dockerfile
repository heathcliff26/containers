FROM quay.io/fedora/fedora-coreos:next@sha256:3999564230f26d3ffc74364dcb1bf0439a7e9b32617fa85a201f7bcd47e058d3

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.31.1

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)\.[0-9]+$
ENV CRIO_VERSION=1.31

# renovate: datasource=github-releases depName=heathcliff26/kube-upgrade extractVersion=^(?<version>.*)$
ENV KUBE_UPGRADE_VERSION=v0.3.2

COPY k8s-install.sh upgraded.service  /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
