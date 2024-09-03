FROM quay.io/fedora/fedora-coreos:stable@sha256:bfc2db4d0e0b6a98a52a8d31aca8237602ffea58e4dbb42c416eb501dffa39ff

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.31.0
# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=(?<version>.*)$
ENV CRIO_VERSION=v1.30

# renovate: datasource=github-releases depName=heathcliff26/kube-upgrade extractVersion=^(?<version>.*)$
ENV KUBE_UPGRADE_VERSION=v0.2.1

COPY k8s-install.sh upgraded.service  /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
