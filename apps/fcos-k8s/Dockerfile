FROM quay.io/fedora/fedora-coreos:stable@sha256:3811dc030edbc0bf4ed59547c5661140ed1bc43473d6461ce8810677a79a6481

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.33.3

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)\.[0-9]+$
ENV CRIO_VERSION=1.33

# renovate: datasource=github-releases depName=heathcliff26/kube-upgrade extractVersion=^(?<version>.*)$
ENV KUBE_UPGRADE_VERSION=v0.4.10

COPY k8s-install.sh upgraded.service /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
