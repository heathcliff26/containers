# fcos-k8s

This is a layered fedora coreos version for running kubernetes.

## Rebasing to the image

In order to rebase to this image, run:
```
sudo rpm-ostree rebase -r ostree-unverified-registry:ghcr.io/heathcliff26/fcos-k8s:latest
```

## Installing kubernetes with kubeadm

Please not that when creating a cluster you need to change the `flex-volume-plugin-dir` to a writeable location.
This can be achieved using a kubeadm configuration file:
```
apiVersion: kubeadm.k8s.io/v1beta3
kind: ClusterConfiguration
controllerManager:
  extraArgs:
    flex-volume-plugin-dir: "/var/kubernetes/kubelet-plugins/volume/exec/"
```

## Links

Examples for coreos-layering:
https://github.com/coreos/layering-examples
