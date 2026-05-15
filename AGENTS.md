# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

Go module `kubevault.dev/apimachinery` — the canonical API types, generated clientsets, OpenAPI definitions, and CRDs for [KubeVault](https://kubevault.com/), AppsCode's Kubernetes operator for HashiCorp Vault. Library only; downstream binaries (`operator`, `cli`, `unsealer`, `hq`, etc.) import this module.

Seven API group/versions (under domain `kubevault.com`):

- `kubevault:v1alpha1` and `kubevault:v1alpha2` — `VaultServer` and related cluster-of-Vault types.
- `catalog:v1alpha1` — `VaultServerVersion` catalog.
- `engine:v1alpha1` — Vault secrets engines as CRDs (`AWSRole`, `AzureRole`, `ElasticsearchRole`, `GCPRole`, `MariaDBRole`, `MongoDBRole`, `MySQLRole`, `PKIRole`, `SecretEngine`, etc.).
- `policy:v1alpha1` — Vault policy types.
- `config:v1alpha1` — runtime config.
- `ops:v1alpha1` — `VaultOpsRequest` (rotate, upgrade, etc.).

## Architecture

- `apis/` — Kubernetes API type definitions, one subdirectory per group:
  - `apis/catalog/`, `apis/config/`, `apis/engine/`, `apis/kubevault/`, `apis/ops/`, `apis/policy/`.
  - Each group has version subdirectories (`v1alpha1`, `v1alpha2`) with `*_types.go` (hand-written), `groupversion_info.go`, `install/`, `fuzzer/`, and generated `zz_generated.*.go` / `openapi_generated.go`.
  - `apis/constants.go`, `apis/doc.go` — cross-group shared constants and docs.
- `client/` — generated typed clientsets, listers, informers. Do not hand-edit.
- `crds/` — generated CRD YAML manifests (one per kind, named `<group>.kubevault.com_<plural>.yaml`) plus `doc.go`.
- `api/openapi-spec/` — generated aggregated OpenAPI spec.
- `pkg/openapi/` — OpenAPI definitions consumed by the operator's aggregated apiserver.
- `hack/` — codegen scripts.
- `third_party/protobuf/` — vendored protobuf assets used by codegen.
- `testdata/apis/` — fuzz test fixtures.
- `vendor/` — checked-in deps.

The CRD API group is `kubevault.com` (matches the project's vanity URL `kubevault.dev` modulo domain).

## Common commands

All Make targets run inside `ghcr.io/appscode/golang-dev` — Docker must be running.

- `make ci` — CI pipeline.
- `make gen` — regenerate everything: clientset, manifests, openapi. Run after any change to `apis/**/*_types.go`.
- `make manifests` — regenerate CRDs only (`gen-crds patch-crds label-crds`).
- `make clientset` — regenerate `client/` only.
- `make openapi` — regenerate `pkg/openapi/` and `api/openapi-spec/`.
- `make fmt`, `make lint`, `make unit-tests` / `make test` — standard.
- `make verify` — `verify-gen verify-modules`; `go mod tidy && go mod vendor` must leave the tree clean.
- `make add-license` / `make check-license` — manage license headers.

Run a single Go test (requires a local Go toolchain):

```
go test ./apis/engine/v1alpha1/... -run TestName -v
```

## Conventions

- Module path is `kubevault.dev/apimachinery` (vanity URL). Imports must use that.
- CRD domain is `kubevault.com`; do not rename without coordinating across every downstream KubeVault repo.
- License: Apache-2.0 (`LICENSE`).
- Sign off commits (`git commit -s`); contributions follow the DCO (`DCO`, `CONTRIBUTING.md`).
- Vendor directory is checked in — `go mod tidy && go mod vendor` must leave the tree clean (enforced by `verify-modules`).
- Do not hand-edit `zz_generated.*.go`, anything under `client/`, `pkg/openapi/openapi_generated.go`, `api/openapi-spec/`, or `crds/*.yaml` — change `apis/<group>/<v>/*_types.go` and re-run `make gen`.
- API groups are listed in `API_GROUPS` in the Makefile; new groups must be added there so codegen fans out correctly.
- Adding a new secrets engine: drop a new `*Role` (or `*Engine`) type into `apis/engine/v1alpha1/`, register it in `groupversion_info.go`, and re-run `make gen`.
