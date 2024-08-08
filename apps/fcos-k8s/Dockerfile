FROM quay.io/fedora/fedora-coreos:stable@sha256:35679392c36e6a164f490199ca31af4c6fb479c7904a9c41e95ac6b17f699bc0

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.30.2

COPY k8s-install.sh /var/kubernetes/k8s-install.sh

RUN /var/kubernetes/k8s-install.sh
