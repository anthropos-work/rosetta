---
title: "KB Fidelity Audit — M28 Provisioning engine + coverage/verify gate"
date: 2026-06-14
scope: milestone:M28
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| PreflightEnv / never-write-prod / 3-layer isolation guard | `corpus/ops/safety.md` (§2.1–2.5) | `.agentspace/rosetta-extensions/stack-seeding/isolation/`; `safety.md:156-205` cited | PAIRED |
| N=0 main-dev-stack guard precedent | `corpus/ops/safety.md` §2.5 | `stack-seeding/cmd/stackseed/main.go` `doReset`; `dev-stack/dev-setdress.sh:96` | PAIRED |
| Injection override + DIRECTUS_TOKEN strip (fix16/17) | `corpus/ops/rosetta_demo.md`, `corpus/ops/safety.md` §2.3 | `stack-injection/gen_injected_override.py:273`; `stack-core/gen_override.py:84` | PAIRED |
| Run-it-twice / copy-if-absent idempotency contract | `corpus/ops/idempotency.md` | M27 source reader (DNA-driven open); M28 provision (new) | PAIRED |
| Non-fatal pre-flight (warn standard / fail critical) | `corpus/ops/verification.md` | `dev-stack/dev-setdress.sh`, `dev-stack/dev-stack` cmd_up, `demo-stack/up-injected.sh` | PAIRED |
| Demo Clerk minting path (demo-aware coverage) | `corpus/services/clerkenstein.md` | `stack-injection/inject.py` (PK_DEMO + sk_test_<demo>); `gen_injected_override.py:60-95` | PAIRED |
| M27 base: DNA + ingestion reader + check/measure scorer | `corpus/ops/secrets-spec.md` (M29 deliverable, not yet authored) | `stack-secrets/secretdna/`, `stack-secrets/source/`, `stack-secrets/cmd/stacksecrets/` | DOC-ONLY (planned M29) |

## Fidelity Findings
1. **safety.md:156-205 PreflightEnv anchor** — ALIGNED. Line 156 is the `Guard.PreflightEnv` bullet; 205 is the prod-S3 `""` override line. The overview's `safety.md:156-205` citation is exact (304-line file). Fix owner: none.
2. **DIRECTUS_TOKEN strip is real in both override generators** — ALIGNED. `gen_injected_override.py:273` appends `DIRECTUS_TOKEN=` (empty) to every emitted service; `gen_override.py:84` strips it for the dev data-plane consumer. The DNA gene `platform/DIRECTUS_TOKEN` (note: "key-present only, no nonempty") matches: a deliberately-blanked non-prod value still passes coverage. Fix owner: none.
3. **N=0 guard precedent (stackseed --reset)** — ALIGNED. `doReset` refuses `n == 0 && !force` with exactly the message M28's provision will mirror. `safety.md §2.5` documents the doubled n=0-dev guards. Fix owner: none.
4. **Non-fatal pre-flight convention** — ALIGNED. `verification.md:48-49`: a failing check produces a `⚠` block + a hint, never an abort; mirrors `dev-setdress.sh`'s default-on + non-fatal pattern. This is the exact convention M28's `check` wiring follows. Fix owner: none.
5. **Demo Clerk minting** — ALIGNED. `clerkenstein.md` documents `mintpk` (authoritative pk minter) + `inject.py` mints PK_DEMO / `sk_test_<demo>`, baked into `.env.demo-N` + frontend build args — NOT sourced from the secret dir. This is exactly the demo-aware-coverage surface M28 extends. Fix owner: none.

## Completeness Gaps
1. (incidental) The M27 `stack-secrets` code has no corpus doc yet — `secrets-spec.md` is M29's `Delivers →` line (overview "Out:"). The contract M28 reads (the DNA shape, the source-reader layout contract, the exit-code contract) is fully documented IN-CODE (package doc-comments in `dna.go`, `source.go`, `main.go`) which is the authoritative source for this milestone's build. Not a blind area — a planned future doc.

## Applied Fixes
None needed — every load-bearing claim ALIGNED on first check.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to Phase 1. All M28 topics are PAIRED + ALIGNED, or DOC-ONLY-by-plan (secrets-spec.md is M29). No blind area, no stale load-bearing claim, no critical undocumented behavior in code M28 extends. The M27 contract M28 builds on is documented in-code (the authoritative source for the build).
