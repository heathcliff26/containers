#!/bin/bash

set -ex

# Enable the kubernetes repo and install the packages as described in:
# https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/

KUBERNETES_SHORT_VERSION="${KUBERNETES_VERSION%.*}"

# This overwrites any existing configuration in /etc/yum.repos.d/kubernetes.repo
cat <<EOF | tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://pkgs.k8s.io/core:/stable:/v${KUBERNETES_SHORT_VERSION}/rpm/
enabled=1
gpgcheck=1
gpgkey=https://pkgs.k8s.io/core:/stable:/v${KUBERNETES_SHORT_VERSION}/rpm/repodata/repomd.xml.key
EOF

rpm-ostree install "kubelet-${KUBERNETES_VERSION}" "kubeadm-${KUBERNETES_VERSION}" "kubectl-${KUBERNETES_VERSION}" cri-o
rpm-ostree cleanup -m

systemctl enable kubelet.service crio.service

# Prepare the host as described in:
# https://docs.fedoraproject.org/en-US/quick-docs/using-kubernetes/

cat <<EOF | tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF

cat <<EOF | tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
net.ipv6.conf.all.forwarding        = 1
EOF

# Ensure we have shell completion
kubeadm completion bash > /etc/bash_completion.d/kubeadm
kubectl completion bash > /etc/bash_completion.d/kubectl
crictl completion bash > /etc/bash_completion.d/crictl

ostree container commit
