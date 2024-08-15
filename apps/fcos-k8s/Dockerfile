FROM quay.io/fedora/fedora-coreos:stable@sha256:5240392ed4ffa33f4a0dbfb94ea9409583d42eaa5981403236ebeddd987f7cea

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.31.0

COPY k8s-install.sh /var/kubernetes/k8s-install.sh

RUN /var/kubernetes/k8s-install.sh
