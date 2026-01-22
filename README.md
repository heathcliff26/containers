[![bitwarden-serve](https://github.com/heathcliff26/containers/actions/workflows/build-bitwarden-serve.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-bitwarden-serve.yaml)
[![default-webpage](https://github.com/heathcliff26/containers/actions/workflows/build-default-webpage.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-default-webpage.yaml)
[![fcos-k8s](https://github.com/heathcliff26/containers/actions/workflows/build-fcos-k8s.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-fcos-k8s.yaml)
[![fcos-k8s v1.32](https://github.com/heathcliff26/containers/actions/workflows/build-fcos-k8s-v1.32.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-fcos-k8s-v1.32.yaml)
[![fcos-k8s v1.33](https://github.com/heathcliff26/containers/actions/workflows/build-fcos-k8s-v1.33.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-fcos-k8s-v1.33.yaml)
[![fcos-k8s v1.34](https://github.com/heathcliff26/containers/actions/workflows/build-fcos-k8s-v1.34.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-fcos-k8s-v1.34.yaml)
[![file-sync](https://github.com/heathcliff26/containers/actions/workflows/build-file-sync.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-file-sync.yaml)
[![github-actions-runner](https://github.com/heathcliff26/containers/actions/workflows/build-github-actions-runner.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-github-actions-runner.yaml)
[![go-fyne-ci](https://github.com/heathcliff26/containers/actions/workflows/build-go-fyne-ci.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-go-fyne-ci.yaml)
[![keepalived](https://github.com/heathcliff26/containers/actions/workflows/build-keepalived.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-keepalived.yaml)
[![openwrt-package-sync](https://github.com/heathcliff26/containers/actions/workflows/build-openwrt-package-sync.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-openwrt-package-sync.yaml)
[![squid](https://github.com/heathcliff26/containers/actions/workflows/build-squid.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-squid.yaml)
[![tang](https://github.com/heathcliff26/containers/actions/workflows/build-tang.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-tang.yaml)
[![unbound](https://github.com/heathcliff26/containers/actions/workflows/build-unbound.yaml/badge.svg)](https://github.com/heathcliff26/containers/actions/workflows/build-unbound.yaml)

# Containers

This is a monorepo containing all the builds for the container images i use. I choose to do it in a monorepo because it is easier to re-use the workflows and renovate config.

## Build

All the Dockerfiles can be found under `apps/<name>/Dockerfile`. The associated workflow is named `build-<name>.yaml`.

More Details to the individual containers can be found in their respective README.md files.

## Directory structure for golang projects

The directory structure for golang projects in this repo follows [golang-standards](https://github.com/golang-standards/project-layout).

It has the following folders:
- `apps/<app>/config`: Config files if required
- `apps/<app>/cmd`: Go-files building the main package
- `apps/<app>/pkg`: Re-usable packages used for the app
- `apps/<app>/vendor`: Folder containing all dependencies needed to build the project
