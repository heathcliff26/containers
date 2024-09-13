FROM quay.io/fedora/fedora-coreos:rawhide@sha256:653b3633698ae504c76547bbab6f94ed1f0537dc82252939f47538853320d3b3

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.31.1

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)$
ENV CRIO_VERSION=1.30

# renovate: datasource=github-releases depName=heathcliff26/kube-upgrade extractVersion=^(?<version>.*)$
ENV KUBE_UPGRADE_VERSION=v0.3.0

COPY k8s-install.sh upgraded.service  /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
