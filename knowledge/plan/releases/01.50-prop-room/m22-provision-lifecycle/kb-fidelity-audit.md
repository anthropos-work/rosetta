---
title: "KB Fidelity Audit — M22 (Executed provisioning + per-stack Directus lifecycle)"
date: 2026-06-13
scope: milestone:M22
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Executed provisioning in the shared engine | `corpus/ops/snapshot-spec.md` §store-fork + `directus-local.md` §"What's still future work" | `dev-stack/dev-setdress.sh` (`snapshot_step`), `stack-snapshot/directus/provision.go`, `cmd/provision-plan/main.go` | PAIRED (doc names it print-only-today → M22 executes) |
| Compose-service emission | `directus-local.md` §future (M22) | `stack-injection/gen_injected_override.py` (`build_lines`/`frontend_lines` net-new-block pattern), `stack-core/gen_override.py` | PAIRED (the append-block pattern exists; M22 adds a directus block) |
| Idempotent re-provision | `corpus/ops/idempotency.md` (M17 contract) | `cmd/stacksnap/autoprovision.go` (gap-gated apply), `dev-setdress.sh` | PAIRED (M17 re-run contract is the model; M22 adds bootstrap/container guards) |
| Directus verify probes | `corpus/ops/verification.md` (M18 net) | `stack-verify/lib/services.sh` (SERVICES registry), `live/autoverify.sh` (cheap-win asserts), `lib/readiness.sh` | PAIRED (M18 framework; M22 adds a directus row + cheap-win) |
| 12 GB-VM preflight | `corpus/ops/demo/frontend-tier.md` + `directus-local.md` | `demo-stack/up-injected.sh` (`preflight_vm_ram`) | PAIRED (M19 preflight; M22 extends accounting) |
| Per-stack Directus recipe / env contract | `directus-local.md` (full) + `snapshot-spec.md` §store-fork | `stack-snapshot/directus/provision.go` (`ProvisionPlan`/`EnvContract`/`Validate`/`DefaultEnvContract`) | PAIRED — verified ALIGNED |

## Fidelity Findings

1. **directus-local.md bootstrap empirics vs `DefaultEnvContract`** — doc claims `admin@<stack>.example.com`, `.local` rejected, `DB_SEARCH_PATH=directus` required, split host/container DSN. Code (`provision.go:154-171`, `ProvisionPlan` steps) matches exactly. **ALIGNED.**
2. **directus-local.md exit semantics table vs `autoprovision.go`** — doc's "bootstrapped gap + structure → auto-provision exit 0 / diverged → exit 5 / empty → exit 4" matches `tryAutoProvision` gap-gating (`nUser > 0` no-op) + the probe ordering. **ALIGNED.**
3. **idempotency.md M17 verdict table vs replay/seed code** — TRUNCATE-then-reload, idempotent COPY, casbin WHERE-NOT-EXISTS all still present. **ALIGNED.** (M22 adds new rows, does not contradict.)
4. **verification.md cheap-win asserts vs `autoverify.sh`** — `/api/health` + `casbin_rules > 0` + scoped verify live all present and offset/scope-aware. **ALIGNED.** (M22 adds a directus cheap-win in the same pattern.)
5. **snapshot-spec.md "print-only for BOTH stack types" (line 386) vs `dev-setdress.sh` snapshot_step** — the engine prints the recipe + firewall-checks the env, executes no boot. **ALIGNED** — this IS the current truth M22 changes. Not a stale claim; M22's Phase 5 retires it.

## Completeness Gaps

1. **(incidental) `provision.go:108` anchor in `directus-local.md:40`** — the doc's historical aside ("previously dead-ended at the print-only `provision.go:108` placeholder") points at a line that, post-M21, holds the content-schema step Detail (no longer a placeholder). The "previously … was" framing is historically accurate, but the numeric anchor now lands on different content. M22 rewrites this exact area (the lifecycle section), so the fix lands naturally in M22 Phase 5 — prefer dropping the numeric anchor for a symbol/prose reference. Tracked as **KB-1** in decisions.md. Non-blocking (incidental, self-healing in M22's own doc work).

## Applied Fixes
None applied inline — the one completeness gap (KB-1) is best fixed by M22's own doc rewrite of the lifecycle section (Phase 5), not a pre-emptive 1-line edit that the rewrite would re-touch.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to Phase 1. Every M22 topic is PAIRED with a truthful doc; the "print-only / reads-live-from-prod / M22-future" markers across `directus-local.md`, `snapshot-spec.md`, `verification.md`, `idempotency.md` are the *current* contract M22 executes against — no blind areas, no stale load-bearing claims. KB-1 (the one incidental anchor) is tracked for M22 Phase 5.
