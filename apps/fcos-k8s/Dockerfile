FROM quay.io/fedora/fedora-coreos:stable@sha256:9ccaf584660e7acf63f7f6658c13a3adf97e88f6e192b241a3b4abd364a55a2b

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.34.1

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)\.[0-9]+$
ENV CRIO_VERSION=1.33

# renovate: datasource=github-releases depName=heathcliff26/kube-upgrade extractVersion=^(?<version>.*)$
ENV KUBE_UPGRADE_VERSION=v0.4.12

COPY k8s-install.sh upgraded.service /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
