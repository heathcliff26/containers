# Containers

This is a monorepo containing all the builds for the container images i use. I choose to do it in a monorepo because it is easier to re-use the workflows and renovate config.

## Build

All the Dockerfiles can be found under `apps/<name>/Dockerfile`. The associated workflow is named `build-<name>.yaml`.
