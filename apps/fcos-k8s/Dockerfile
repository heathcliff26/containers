FROM quay.io/fedora/fedora-coreos:stable@sha256:02f523ab0c7772e59577e003bc2cd3b2ebc9d287a971f3d8d0c934f314db69e9

LABEL   org.opencontainers.image.authors="heathcliff@heathcliff.eu" \
        org.opencontainers.image.description="Fedora CoreOS with Kubernetes and CRI-O" \
        org.opencontainers.image.source="https://github.com/heathcliff26/containers/tree/main/apps/fcos-k8s" \
        org.opencontainers.image.licenses="Apache-2.0" \
        org.opencontainers.image.title="fcos-k8s"

LABEL quay.expires-after=2y

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.34.1

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)\.[0-9]+$
ENV CRIO_VERSION=1.34

COPY k8s-install.sh migration-remove-upgraded.service /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
