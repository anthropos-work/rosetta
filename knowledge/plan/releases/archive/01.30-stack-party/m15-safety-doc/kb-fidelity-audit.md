---
title: "KB Fidelity Audit — M15 (Safety & security doc + dual-repo KB)"
date: 2026-06-07
scope: milestone:M15
invoked-by: build-milestone
---

## Verdict
**GREEN** — every safety claim the milestone will consolidate into `corpus/ops/safety.md` is ALIGNED with the actual extensions code. The new doc can be authored against verified facts.

This is the load-bearing-doc audit: M15's deliverable (safety.md) is a *claim surface* a future reader trusts, so the audit verifies the underlying code BEFORE the doc restates it. Every read-side / write-side claim in the overview was cross-checked against the real implementation; the symbol names, predicates, and behaviors match. One naming nuance and one over-broad source-comment surfaced — captured below as accuracy notes, NOT blockers.

## Topic Inventory

| Topic | Knowledge doc(s) | Code path(s) | Status |
|---|---|---|---|
| Read-side tenant firewall (`AssertPublicOnly`) | `snapshot-spec.md`, `db-access.md` | `stack-snapshot/firewall/firewall.go` (`AssertPlan` + `AssertCaptured`) | PAIRED |
| Public predicates (taxonomy + directus) | `snapshot-spec.md`, `db-access.md` | `firewall.go` (`PublicFilter`, `DefaultPredicate`, directus predicate) | PAIRED |
| Capture-source policy (dump-ingest → primary-read, MVCC, READ ONLY) | `snapshot-spec.md` | `stack-snapshot/source/source.go` | PAIRED |
| 3-layer write isolation guard (`CheckWrite`/`PreflightEnv`/`AssertClean`) | `seeding-spec.md` | `stack-seeding/isolation/isolation.go` + `audit.go` | PAIRED |
| Never-write shared Directus / prod-S3 | `seeding-spec.md` | `isolation.go` registry + `audit.go` `PreflightEnv` | PAIRED |
| n=0-dev guards | (none yet — safety.md will be first) | `dev-stack/dev-setdress.sh:55-57`, `stackseed/main.go:181` | CODE-ONLY (expected — safety.md documents it) |
| Audit-proven zero pollution | `seeding-spec.md` | `audit.go` `AssertClean` | PAIRED |
| safety.md itself | — (net-new, M15 delivers) | — | DOC-ONLY (the deliverable) |

## Fidelity Findings

1. **`AssertPublicOnly` is a CONCEPTUAL name, not a Go symbol.** `firewall.go` documents "AssertPublicOnly is invoked twice, defense in depth" but the actual exported functions are `AssertPlan` (plan-time gate) and `AssertCaptured` (post-capture gate). The existing docs (db-access, snapshot-spec) already use `AssertPublicOnly` as the umbrella name — **ALIGNED** with that convention. safety.md should follow suit but name the two real functions so a reader can grep them. **Verdict: ALIGNED** (doc convention is intentional and pre-existing).

2. **Public predicates match exactly.** Taxonomy = `organization_id IS NULL` (`firewall.PublicFilter`); Directus = `private = false AND tenant_id IS NULL AND status = 'published'` (the directus `PublicPredicate.PublicFilter`, referenced in `firewall.go:72` and the parent-scope examples). **Verdict: ALIGNED.**

3. **Capture-source policy precedence is dump-ingest [default] → primary-read [fallback] → restore-from-snapshot / read-replica [upgrades, not wired].** `source.go` `DefaultPrecedence` + `Kind.Available()` match. The bounded read session applies `SET TRANSACTION READ ONLY` + `statement_timeout` + `idle_in_transaction_session_timeout` (`source.go` `SetupSQL`). MVCC = no write-blocking note is in the package doc. **The offline pg_dump-FILE reader was DROPPED (M9b-D9)** — a dump is ingested by RESTORING into Postgres and read over a DSN. safety.md must NOT claim an offline file reader. **Verdict: ALIGNED** (and an accuracy guardrail for the new doc).

4. **3-layer guard signatures match seeding-spec exactly.** `Guard.CheckWrite(store, class, t)` (`isolation.go:159`), `Guard.PreflightEnv(env, t)` (`audit.go:97`), `AuditLog.AssertClean(t)` (`audit.go:64`). PreflightEnv forces `STORAGE_S3_PUBLIC_BUCKET=""` always, rejects live-Clerk URLs + live Directus write tokens on non-prod. The non-prod asymmetry (shared/external writes ALWAYS blocked on non-prod, opt-in only relaxes on prod) is the structural guarantee. **Verdict: ALIGNED.**

5. **n=0-dev guard scope (the one nuance to document precisely).** The guard exists in TWO places: `dev-setdress.sh` (refuses to auto-set-dress N=0 without `--force`) and `stackseed/main.go` (refuses `--reset` of N=0 without `--force`). **`stacksnap` (snapshot replay) has NO N=0 guard** — and correctly so: replay writes only public reference data into the stack's OWN isolated Directus/Postgres, so replaying into the main dev stack is harmless. The `dev-setdress.sh:20` comment ("`stacksnap`/`stackseed` independently refuse N=0 too") slightly over-states it for `stacksnap`. safety.md should describe the n=0 guard as protecting against (a) unsolicited auto-set-dressing and (b) destructive `--reset` — NOT as a blanket "all tools refuse N=0". **Verdict: ALIGNED with a precision note** (recorded as M15-D… in decisions.md so the new doc gets it right; the source comment itself is a pre-existing minor inaccuracy, not in M15's diff — flagged, not fixed here to avoid scope creep into M13 code).

## Completeness Gaps
None load-bearing. The "doubled in M13" framing for the n=0 guard refers to its presence in both the auto-set-dress path and the reset path (two independent enforcement points), which is accurate — see finding 5.

## Applied Fixes
None needed — the existing KB docs (snapshot-spec, seeding-spec, db-access, security_compliance) are accurate. The audit's job here is to certify the facts BEFORE safety.md restates them; no stale claim to repair.

## Open Items (require user decision)
None.

## Gate Result
**GREEN: proceed to Phase 1.** All safety claims verified against code. Two accuracy guardrails carried into the build (no offline-file-reader claim; precise n=0-guard scoping) — recorded in `decisions.md`.
