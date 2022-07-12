#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

go mod vendor

bash ./vendor/k8s.io/code-generator/generate-groups.sh "all" \
  github.com/ferryproxy/client-go/generated github.com/ferryproxy/api/apis \
  traffic:v1alpha2 \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --output-base "${SCRIPT_ROOT}/../../.."
