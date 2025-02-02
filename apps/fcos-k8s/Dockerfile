FROM quay.io/fedora/fedora-coreos:stable@sha256:dfc958657e56d105d6cfb3278893b845eb8d223a063a7593a1baf7db103cc60f

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.32.1

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)\.[0-9]+$
ENV CRIO_VERSION=1.32

# renovate: datasource=github-releases depName=heathcliff26/kube-upgrade extractVersion=^(?<version>.*)$
ENV KUBE_UPGRADE_VERSION=v0.4.4

COPY k8s-install.sh upgraded.service /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
