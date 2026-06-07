# M15 — Spec notes

Technical notes accumulate here during build.

## Pre-flight audits — §1 read-side
- **Phase 0b KB-fidelity: GREEN** (2026-06-07). Report: `kb-fidelity-audit.md`.
- Topic→doc→code triples (verified ALIGNED — the facts safety.md restates):
  - read firewall → `snapshot-spec.md` / `db-access.md` → `stack-snapshot/firewall/firewall.go` (`AssertPlan` plan-time + `AssertCaptured` post-capture; conceptual umbrella name `AssertPublicOnly`).
  - public predicates → same docs → `firewall.go` `PublicFilter` (`organization_id IS NULL`) + directus `PublicPredicate` (`private = false AND tenant_id IS NULL AND status = 'published'`).
  - capture-source policy → `snapshot-spec.md` → `stack-snapshot/source/source.go` (`DefaultPrecedence`: dump-ingest→primary-read→[restore-snapshot/read-replica, not wired]; `BoundedSession.SetupSQL` = `SET TRANSACTION READ ONLY` + timeouts; MVCC = no write-block; **no offline file reader** — M9b-D9).
  - 3-layer write guard → `seeding-spec.md` → `stack-seeding/isolation/isolation.go` (`Guard.CheckWrite`, 3 IsolationClasses, non-prod asymmetry) + `audit.go` (`Guard.PreflightEnv` forces `STORAGE_S3_PUBLIC_BUCKET=""` + rejects live Clerk/Directus on non-prod; `AuditLog.AssertClean` = zero-pollution proof).
  - n=0 guard → (safety.md is first doc) → `dev-stack/dev-setdress.sh:55-57` (auto-set-dress) + `stack-seeding/cmd/stackseed/main.go:181` (`--reset`). `stacksnap` replay has NO N=0 guard (correct — see M15-D4).

## safety.md — read-side (never reads private data)
_Tenant firewall AssertPublicOnly; public predicates (org_id IS NULL; private=false AND tenant_id IS NULL AND status=published); public-only data-DNA gene; capture read-only._

## safety.md — write-side (never touches prod)
_3-layer isolation guard; never-write shared Directus / prod-S3; capture-source policy; n=0-dev guards; audit-proven zero pollution._

## Dual-repo KB refresh
_rosetta corpus + rosetta-extensions/knowledge/ to the v1.3 converged stack model._
