FROM quay.io/fedora/fedora-coreos:stable@sha256:a18cc8dcf65cd6ce4916b6e6d39820e4b67bb6f52d2dd28ac4d47b241fc3d770

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Fedora CoreOS with Kubernetes and CRI-O" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/fcos-k8s" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="fcos-k8s"

LABEL quay.expires-after=2y

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.35.3

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)\.[0-9]+$
ENV CRIO_VERSION=1.35

COPY k8s-install.sh migration-remove-upgraded.service /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
