FROM quay.io/fedora/fedora-coreos:stable@sha256:ef0f7f4ff37ef622c0ddc95c4af164b58da9107db18d3402977afd59f032db94

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.33.2

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)\.[0-9]+$
ENV CRIO_VERSION=1.33

# renovate: datasource=github-releases depName=heathcliff26/kube-upgrade extractVersion=^(?<version>.*)$
ENV KUBE_UPGRADE_VERSION=v0.4.10

COPY k8s-install.sh upgraded.service /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
