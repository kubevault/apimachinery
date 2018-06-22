#!/bin/bash

set -x

GOPATH=$(go env GOPATH)
PACKAGE_NAME=github.com/kubevault/operator
REPO_ROOT="$GOPATH/src/$PACKAGE_NAME"
DOCKER_REPO_ROOT="/go/src/$PACKAGE_NAME"
DOCKER_CODEGEN_PKG="/go/src/k8s.io/code-generator"
apiGroups=(core/v1alpha1 extensions/v1alpha1)

pushd $REPO_ROOT

# for EAS types
docker run --rm -ti -u $(id -u):$(id -g) \
  -v "$REPO_ROOT":"$DOCKER_REPO_ROOT" \
  -w "$DOCKER_REPO_ROOT" \
  appscode/gengo:release-1.10 "$DOCKER_CODEGEN_PKG"/generate-internal-groups.sh "deepcopy,defaulter,conversion" \
  github.com/kubevault/operator/client \
  github.com/kubevault/operator/apis \
  github.com/kubevault/operator/apis \
  extensions:v1alpha1 \
  --go-header-file "$DOCKER_REPO_ROOT/hack/gengo/boilerplate.go.txt"

# for both CRD and EAS types
docker run --rm -ti -u $(id -u):$(id -g) \
  -v "$REPO_ROOT":"$DOCKER_REPO_ROOT" \
  -w "$DOCKER_REPO_ROOT" \
  appscode/gengo:release-1.10 "$DOCKER_CODEGEN_PKG"/generate-groups.sh all \
  github.com/kubevault/operator/client \
  github.com/kubevault/operator/apis \
  "core:v1alpha1 extensions:v1alpha1" \
  --go-header-file "$DOCKER_REPO_ROOT/hack/gengo/boilerplate.go.txt"

# Generate openapi
for gv in "${apiGroups[@]}"; do
  docker run --rm -ti -u $(id -u):$(id -g) \
    -v "$REPO_ROOT":"$DOCKER_REPO_ROOT" \
    -w "$DOCKER_REPO_ROOT" \
    appscode/gengo:release-1.10 openapi-gen \
    --v 1 --logtostderr \
    --go-header-file "hack/gengo/boilerplate.go.txt" \
    --input-dirs "$PACKAGE_NAME/apis/${gv},k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/api/resource,k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/version,k8s.io/api/core/v1" \
    --output-package "$PACKAGE_NAME/apis/${gv}"
done

# Generate crds.yaml and swagger.json
go run ./hack/gencrd/main.go

popd
