FROM quay.io/fedora/fedora-coreos:stable@sha256:c34e95a37b9f62ed9391928d29b403defdae8f530221ab218ece11f2f5b9bed8

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.31.3

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)\.[0-9]+$
ENV CRIO_VERSION=1.31

# renovate: datasource=github-releases depName=heathcliff26/kube-upgrade extractVersion=^(?<version>.*)$
ENV KUBE_UPGRADE_VERSION=v0.4.1

COPY k8s-install.sh upgraded.service  /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
