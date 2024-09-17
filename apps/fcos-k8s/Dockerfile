FROM quay.io/fedora/fedora-coreos:rawhide@sha256:77a40922655cc84c2c71687a5dfa49a93dda9b7b4f5ac97136a519c0dc2b6333

# renovate: datasource=github-tags depName=kubernetes/kubernetes extractVersion=^v(?<version>.*)$
ENV KUBERNETES_VERSION=1.31.1

# renovate: datasource=github-releases depName=cri-o/cri-o extractVersion=^v(?<version>.*)$
ENV CRIO_VERSION=1.30

# renovate: datasource=github-releases depName=heathcliff26/kube-upgrade extractVersion=^(?<version>.*)$
ENV KUBE_UPGRADE_VERSION=v0.3.0

COPY k8s-install.sh upgraded.service  /var/kubernetes/

RUN /var/kubernetes/k8s-install.sh
