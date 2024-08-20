#!/bin/bash

set -ex

# Enable the kubernetes repo and install the packages as described in:
# https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/

KUBERNETES_SHORT_VERSION="${KUBERNETES_VERSION%.*}"

echo "Create kubernetes.repo file"
# This overwrites any existing configuration in /etc/yum.repos.d/kubernetes.repo
cat <<EOF | tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://pkgs.k8s.io/core:/stable:/v${KUBERNETES_SHORT_VERSION}/rpm/
enabled=1
gpgcheck=1
gpgkey=https://pkgs.k8s.io/core:/stable:/v${KUBERNETES_SHORT_VERSION}/rpm/repodata/repomd.xml.key
EOF

echo "Install packages"
rpm-ostree install "kubelet-${KUBERNETES_VERSION}" "kubeadm-${KUBERNETES_VERSION}" "kubectl-${KUBERNETES_VERSION}" cri-o
rpm-ostree cleanup -m

echo "Enable kubelet and crio services"
systemctl enable kubelet.service crio.service

# Prepare the host as described in:
# https://docs.fedoraproject.org/en-US/quick-docs/using-kubernetes/

echo "Ensure kernel modules are loaded"
cat <<EOF | tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF

echo "Ensure sysctl settings are updated"
cat <<EOF | tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
net.ipv6.conf.all.forwarding        = 1
EOF

echo "Add bash autocompletion"
# Ensure we have shell completion
kubeadm completion bash >/etc/bash_completion.d/kubeadm
kubectl completion bash >/etc/bash_completion.d/kubectl
crictl completion bash >/etc/bash_completion.d/crictl

arch="$(uname -m)"
case "${arch}" in
    x86_64|amd64)
        arch="amd64"
        ;;
    aarch64|arm64)
        arch="arm64"
        ;;
esac

upgraded_url="https://github.com/heathcliff26/kube-upgrade/releases/download/${KUBE_UPGRADE_VERSION}/upgraded-${arch}"
echo "Install upgraded ${KUBE_UPGRADE_VERSION} from ${upgraded_url}"
curl -SL -o /usr/libexec/upgraded "${upgraded_url}"
chmod 755 /usr/libexec/upgraded
/usr/libexec/upgraded version

echo "Install and enable upgraded.service"
cp /var/kubernetes/upgraded.service /etc/systemd/system/
systemctl enable upgraded.service
mkdir /etc/kube-upgraded

echo "Commit changes"
ostree container commit
