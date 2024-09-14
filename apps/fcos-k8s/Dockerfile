FROM quay.io/fedora/fedora-coreos:rawhide@sha256:209898e1f352cb842a964a3ff93a77df93536bb73ff58acd01deefbb3cb4fe0a

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.31.1

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)$
ENV CRIO_VERSION=1.30

# renovate: datasource=github-releases depName=heathcliff26/kube-upgrade extractVersion=^(?<version>.*)$
ENV KUBE_UPGRADE_VERSION=v0.3.0

COPY k8s-install.sh upgraded.service  /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
