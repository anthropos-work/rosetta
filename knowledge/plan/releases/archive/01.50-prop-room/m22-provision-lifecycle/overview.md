---
milestone: M22
slug: provision-lifecycle
version: v1.5 "prop room"
milestone_shape: section
status: archived
created: 2026-06-11
last_updated: 2026-06-13
complexity: medium
delivers: corpus/ops/directus-local.md (lifecycle half) + verification.md/idempotency.md rows + executed provisioning + compose-service emission + Directus verify probes
backlog_refs: NEW-2 (per-stack-Directus recipe execution/boot automation)
---

# M22 — Executed provisioning + per-stack Directus lifecycle

## Goal
Turn the print-only 4-step recipe into an **executed** bring-up step that boots a per-stack Directus as a
**compose service** in the stack's override — idempotent, verified, torn down with the stack — so demo (default)
and opt-in dev stacks come up with a live local Directus.

## Why section
The shape is fully known once M21 lands: the recipe steps are written (print-only today), the override generators
exist, the verify framework exists, the idempotency contract exists. M22 is concrete edits + wiring — no exploratory
uncertainty. Build with `/developer-kit:build-milestone`.

## Repo split
- **`rosetta-extensions`** (authoring → tag `prop-room-m22` → consume): the executed provisioning in the shared
  `dev-setdress.sh` engine, the compose-service emission, the idempotent re-provision guards, the Directus verify
  probes, the 12 GB-VM preflight extension.
- **`rosetta`**: `corpus/ops/directus-local.md` (the lifecycle half) + new rows in `verification.md` / `idempotency.md`
  + the `rosetta_demo.md` registry/teardown note.

## Scope
- **In (`rosetta-extensions`):**
  - **Execute** bootstrap → apply-structure → replay → boot inside the shared `dev-setdress.sh` engine (replacing the
    print-only block); **demo default-on / dev opt-in** (`--local-content`; `dev-N≥1` direct, `N=0` additionally
    behind the existing `--force` n=0 guard).
  - **Emit the Directus container into the per-stack override as a compose service** (offset port `8055+N·10000`,
    joins the stack's app-network, named `<project>-directus-1`) — **not** a bespoke `docker run` (so `demo-down`/
    `dev-down`, the port registry, and `stack-verify`'s naming convention all cover it with no new lifecycle code —
    the maintainability constraint).
  - **Idempotent re-provision** (bootstrap-on-non-empty + container-name-conflict guards, matching the M17 re-run
    contract); **register the Directus offset port**.
  - **Directus verify probes** in `stack-verify`: a SERVICES row + `/server/health`, `directus` added to the
    expected-schemas list, a **"registered collections > 0"** cheap-win (the silent-failure analog of the casbin
    assert), and a **no-prod-read env assert**.
  - The **12 GB-VM preflight** accounting extended to include the Directus container; **non-fatal** (a failed boot
    degrades to the prod-read path with an honest status line — never blocks a good stack).
- **In (`rosetta`):** update `corpus/ops/verification.md` + `corpus/ops/idempotency.md` (new rows) +
  `corpus/ops/rosetta_demo.md` (registry/teardown).
- **Out:** the env re-point (M23 — M22 boots + verifies the instance; M23 points services at it); referential
  closure (M23).

## Depends on / parallel with
Depends on: M21 (a replay that exits 0 is the prerequisite for a Directus that serves content). Parallel with: none.

## Open questions
- Compose-service vs sidecar override file (lean: a service block in the existing injected/dev override — reuses the
  proven generator, nothing bespoke).
- How loud a degraded "boot failed → still reading prod" status should be (lean: a clear ⚠ line in the set-dress
  status, consistent with the M18 verify block).

## KB dependencies
`corpus/ops/snapshot-spec.md` (the store-fork recipe), `corpus/ops/verification.md`, `corpus/ops/idempotency.md`,
`corpus/ops/rosetta_demo.md`, the `dev-setdress`/`stack-verify` sections.

## Risk (data-safety — blocks-prod-safety)
An executed provision must only ever write the per-stack-isolated offset Directus/Postgres — the `EnvContract.Validate`
firewall moves from a print-time check to a **load-bearing executed gate** (hard-abort before any write if the env
resolves to prod); tests pin the target class, as M17 did for TRUNCATE.
