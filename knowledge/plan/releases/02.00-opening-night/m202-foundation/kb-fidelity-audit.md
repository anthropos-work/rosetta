---
title: "KB Fidelity Audit — M202 Playthroughs Foundation"
date: 2026-07-01
scope: milestone:M202
invoked-by: build-milestone
---

## Verdict
GREEN

The one BLIND-AREA (`corpus/ops/demo/playthroughs.md`) is a **declared milestone deliverable** — the
`Delivers →` line already promotes knowledge production into M202's scope (the audit's sanctioned path for a
blind area). Every code-reuse anchor the milestone builds on is PAIRED and ALIGNED. No stale load-bearing
claims. The calling skill may proceed to Phase 1.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Playthroughs capability contract | `knowledge/plan/spec-drafts/playthroughs/spec.md` (v0.3) | (net-new `playthroughs` rext section) | DOC-ONLY (the spec M202 builds) |
| Manifest corpus (regression contract) | `.../m201-manifest-corpus/manifest-draft.yaml` + `overview.md` | (net-new validator) | DOC-ONLY (prose contract to validate against) |
| M42 e2e foundation (cockpit-login, section-assert, coverage-manifest, empty-states) | `corpus/ops/demo/coverage-protocol.md` | `stack-verify/e2e/lib/*.ts` | PAIRED |
| Seeding machinery + `--reset` + isolation guard | `corpus/ops/seeding-spec.md` | `stack-seeding/{cmd/stackseed,isolation,blueprint,dna}` | PAIRED |
| Re-run / `--reset` contract + N=0 `--force` guard | `corpus/ops/idempotency.md` | `stack-seeding/cmd/stackseed/main.go::doReset` | PAIRED |
| `datadna` conformance gate | `corpus/ops/seeding-spec.md` §data-DNA | `stack-seeding/cmd/datadna` + `stack-seeding/dna/` | PAIRED |
| **Playthroughs corpus runbook** | `corpus/ops/demo/playthroughs.md` **(does not exist)** | (net-new section) | **BLIND-AREA → declared `Delivers →` deliverable** |

## Fidelity Findings

1. **`--reset` full-fleet + N=0 `--force` guard** — Source: `seeding-spec.md:173,178,186` +
   `idempotency.md:41,57,157`. Expected: `stackseed --reset` truncates the full FK-ordered fleet
   (per-stack only) and refuses N=0 unless `--force`. Actual: `cmd/stackseed/main.go::doReset` (l.614-665)
   truncates `resetTables` (full child-first FK order) with the `n == 0 && !force` refusal at l.622.
   **Verdict: ALIGNED.**
2. **Idempotent seed = `ON CONFLICT DO NOTHING` + casbin `WHERE NOT EXISTS`** — Source: `idempotency.md:41`.
   Actual: matches the seeder fleet; `resetCasbin` does the targeted `DELETE … WHERE p_type='g2'` (l.671-682),
   never a TRUNCATE of the policy. **Verdict: ALIGNED.**
3. **Isolation guard = three classes, non-prod shared/external writes always blocked** — Source:
   `seeding-spec.md:84-106`. Actual: `isolation/isolation.go` — `PerStackIsolated`/`SharedPollutionRisk`/
   `External`; `Guard.CheckWrite` blocks shared/external on non-prod regardless of opt-in (l.159-178);
   `AuditLog.AssertClean` proves zero pollution post-run (`cmd/stackseed` l.509). **Verdict: ALIGNED.**
4. **cockpit-login handshake (roster → fake-FAPI → `?__clerk_identity=`)** — Source:
   `coverage-protocol.md` + spec §5.6. Actual: `stack-verify/e2e/lib/cockpit-login.ts` — `selectSeat`
   (`POST /v1/demo/select`) then a protected-route handshake; `loginAs` fails loud on a `/login` bounce.
   **Verdict: ALIGNED** (reused as-is for hero login).
5. **Coverage manifest is a TS descriptor model** — Source: spec §5.6. Actual:
   `coverage-manifest.ts` exports `VantageManifest`/`PageDescriptor`/`SectionDescriptor`. The playthroughs
   manifest is a distinct (YAML prose-intent) model per spec §5.3 — it EXTENDS the section-model idea, not
   the TS type. **Verdict: ALIGNED** (no claim of type reuse).

## Completeness Gaps
None load-bearing. The playthroughs manifest introduces a new YAML schema (§5.3) whose field set the spec
explicitly settles "during the first build" — i.e. M202 authors it; not a pre-existing undocumented behavior.

## Applied Fixes
- Recorded the topic→doc→code triples in `spec-notes.md` (this milestone).
- No doc edits needed — every audited claim is already ALIGNED.

## Open Items (require user decision)
None.

## Gate Result
GREEN — proceed. The single blind area is the milestone's own declared deliverable
(`corpus/ops/demo/playthroughs.md`, authored in the Docs section). No stale load-bearing claims.
