## Development Guide
This document is intended to be the canonical source of truth for things like supported toolchain versions for building Guard.
If you find a requirement that this doc does not capture, please submit an issue on github.

This document is intended to be relative to the branch in which it is found. It is guaranteed that requirements will change over time
for the development branch, but release branches of Guard should not change.

### Build Guard
Some of the Guard development helper scripts rely on a fairly up-to-date GNU tools environment, so most recent Linux distros should
work just fine out-of-the-box.

#### Setup GO
Guard is written in Google's GO programming language. Currently, Guard is developed and tested on **go 1.8.3**. If you haven't set up a GO
development environment, please follow [these instructions](https://golang.org/doc/code.html) to install GO.

#### Download Source

```console
$ go get github.com/kubevault/operator
$ cd $(go env GOPATH)/src/github.com/kubevault/operator
```

#### Install Dev tools
To install various dev tools for Guard, run the following command:

```console
# setting up dependencies for compiling steward...
$ ./hack/builddeps.sh
```

#### Build Binary
```
$ ./hack/make.py
$ steward version
```

#### Dependency management
Guard uses [Glide](https://github.com/Masterminds/glide) to manage dependencies. Dependencies are already checked in the `vendor` folder.
If you want to update/add dependencies, run:
```console
$ glide slow
```

#### Build Docker images
To build and push your custom Docker image, follow the steps below. To release a new version of Guard, please follow the [release guide](/docs/developer-guide/release.md).

```console
# Build Docker image
$ ./hack/docker/setup.sh; ./hack/docker/setup.sh push

# Add docker tag for your repository
$ docker tag soter/vault-operator:<tag> <image>:<tag>

# Push Image
$ docker push <image>:<tag>
```

#### Generate CLI Reference Docs
```console
$ ./hack/gendocs/make.sh
```
