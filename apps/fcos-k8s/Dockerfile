FROM quay.io/fedora/fedora-coreos:stable@sha256:e64639857fcb8474f6db2699f7caebd82fa3fc6890e9bcc0d7d5e6675f83c224

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Fedora CoreOS with Kubernetes and CRI-O" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/fcos-k8s" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="fcos-k8s"

LABEL quay.expires-after=2y

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.36.2

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)\.[0-9]+$
ENV CRIO_VERSION=1.36

COPY k8s-install.sh /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
