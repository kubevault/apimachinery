# Kubebuilder Migration — apimachinery state

Branch: `kubebuilder` (based on `nolgtm`)
Cross-repo plan: `../plan.md` (in the parent kubevault.dev/ directory; not part of any repo)

This file tracks the **apimachinery** side of the migration. See `../operator/MIGRATION_STATE.md` for operator-side state.

## How to resume

If the laptop went to sleep mid-migration, read the **Status** section below to find the last completed chunk, then continue with the next pending chunk. Each completed chunk has a commit hash so you can `git log` to verify.

The current Claude task list (in the assistant's task tracker) mirrors the chunks here. Memory entry: `project_kubebuilder_migration.md`.

## Why apimachinery first
This module is already heavily kubebuilder-style — types carry `+kubebuilder:*` markers, CRDs come from `controller-gen`. The work here is small (PROJECT file, hack/boilerplate.go.txt, optional file renames, Makefile cleanup) and lets us validate the toolchain before touching the operator.

## Status

Legend: `[x]` done · `[~]` in progress · `[ ]` pending

### Phase 1 — Scaffolding parity
- [x] Create `kubebuilder` branch from `nolgtm`
- [x] Add `MIGRATION_STATE.md` (this file)
- [x] Rename `apis/<group>/<version>/register.go` → `groupversion_info.go` (all 7 done)
  - [x] catalog/v1alpha1 (`c4e926ee`)
  - [x] config/v1alpha1, engine/v1alpha1, kubevault/v1alpha1+v1alpha2, ops/v1alpha1, policy/v1alpha1 (`6d2de305`)
- [ ] Skipped: separate `hack/boilerplate.go.txt` (existing `hack/license/go.txt` already serves this; duplicating would drift)
- [ ] Skipped for now: `PROJECT` file. Not required since types/CRDs already exist; only needed if running `kubebuilder create api`.
- [ ] Verify `controller-gen` produces identical `crds/` output (no diff). Deferred to Phase 1 follow-up.

### Phase 1 follow-up
- [x] Replace docker `ghcr.io/appscode/gengo:release-1.32` invocation with a local `controller-gen` for `gen-crds` (`e41b929f`). Pinned at `CONTROLLER_TOOLS_VERSION` (default `v0.16.5`), installed on demand via `go install` into `./bin/`. Other generators (clientset/lister/informer via `generate-groups.sh`, conversion-gen, openapi-gen, go-to-protobuf) are not part of controller-gen and stay on the docker image for now — they need a separate set of pinned binaries.
- [ ] Verify: run `make gen-crds` and confirm `git diff crds/` is empty against the previous docker-generated output. If non-empty, bump `CONTROLLER_TOOLS_VERSION` to match the version the legacy docker image ships.
- [ ] Regenerate `zz_generated.deepcopy.go` with `controller-gen object` — same toolchain swap, deferred.

### Phase 5 cleanup (BLOCKED — see consumer audit below)
- [ ] Delete `client/clientset`, `client/listers`, `client/informers`
- [ ] Drop client/lister/informer codegen targets from Makefile

#### Consumer audit (`grep -r kubevault.dev/apimachinery/client ...`)

Local repos that still import `kubevault.dev/apimachinery/client/...` and would break if the trees are deleted:

| Repo | Role |
|---|---|
| `kubevault.dev/operator` (`pkg/controller/*.go`, `pkg/admission/*.go`, `pkg/cmds/*`, `pkg/server/*`) | This repo's own legacy code — gone when Phase 5 lands here. |
| `kubevault.dev/cli` (`pkg/cmds/*`, `pkg/generate/*`) | Generates manifests for vault engines/roles using the typed clientset. Largest external consumer (~15 files). Needs porting to controller-runtime client before client/ can be deleted. |
| `kubedb.dev/schema-manager` (`pkg/server/server.go`, `test/e2e/suite_test.go`) | Imports the clientset to talk to kubevault CRs. |
| `kubedb.dev/webhook-server` (`pkg/server/server.go`) | Same. |
| `go.bytebuilders.dev/catalog-manager` (`pkg/cmds/run.go`) | Same. |
| `github.com/tamalsaha/{resource-listing-summary,schema-manager-demo}` (`main.go`) | Small one-off demos. |

Recommended sequence:
1. Migrate `kubevault.dev/cli` to the controller-runtime client (separate project).
2. Migrate the three downstream daemons (`schema-manager`, `webhook-server`, `catalog-manager`).
3. Then delete `client/clientset`, `client/listers`, `client/informers` here.

## Notes
- Hub/spoke conversion in `apis/kubevault/v1alpha1/conversion.go` (hand-written 517-line file converting v1alpha1 ↔ v1alpha2). Untouched in Phase 1.
- `apis/<group>/install/install.go` files stay — they're still useful and kubebuilder-compatible.
- `pkg/`, `third_party/`, `api/` (legacy) — review later.
