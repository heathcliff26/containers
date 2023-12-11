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
