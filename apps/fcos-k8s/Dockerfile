FROM quay.io/fedora/fedora-coreos:next@sha256:5544074a29cda4f7f5e3a3e98f873b35eb8eb233a28f97b65430134ee30c1f59

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.31.1

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)\.[0-9]+$
ENV CRIO_VERSION=1.31

# renovate: datasource=github-releases depName=heathcliff26/kube-upgrade extractVersion=^(?<version>.*)$
ENV KUBE_UPGRADE_VERSION=v0.3.1

COPY k8s-install.sh upgraded.service  /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
