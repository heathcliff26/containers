FROM quay.io/fedora/fedora-coreos:stable@sha256:962d96a8739036263308c024c45ec0b78e18396e02bb91f440f9c9f058636120

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.32.3

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)\.[0-9]+$
ENV CRIO_VERSION=1.32

# renovate: datasource=github-releases depName=heathcliff26/kube-upgrade extractVersion=^(?<version>.*)$
ENV KUBE_UPGRADE_VERSION=v0.4.6

COPY k8s-install.sh upgraded.service /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
