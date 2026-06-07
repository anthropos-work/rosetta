---
milestone: M15
slug: safety-doc
version: v1.3 "stack party"
milestone_shape: section
status: planned
created: 2026-06-07
complexity: medium
delivers: corpus/ops/safety.md (new) + updates the rosetta-extensions/knowledge/ base to the v1.3 converged model
---

# M15 — Safety & security doc + dual-repo knowledge consolidation

## Goal
A single authoritative doc on **how the rosetta tooling stays safe** — it never reads private/customer data and it
never touches production data or services — plus a refresh of **both** knowledge bases to the v1.3 converged model.

## Why section
A finite documentation/consolidation deliverable over the shipped M12–M14 result. The M8/M11-analog closing milestone.

## Scope
- **In:**
  - A new **`corpus/ops/safety.md`** consolidating:
    - **read-side** (private-data avoidance): the tenant firewall `AssertPublicOnly`; the public predicates (`organization_id IS NULL`, `private = false AND tenant_id IS NULL AND status = 'published'`); the public-only data-DNA gene; capture is read-only.
    - **write-side** (prod-protection): the 3-layer isolation guard (CheckWrite / PreflightEnv / AssertClean); never-write shared Directus / prod-S3; the capture-source policy (dump-ingest → safe throttled primary read); the n=0-dev guards; the audit-proven zero-pollution assertion.
  - Cross-link from `snapshot-spec` / `seeding-spec` / `db-access` / `security_compliance`.
  - **Update the `rosetta-extensions/knowledge/`** base for the v1.3 converged stack model + the safety contract.
  - Refresh the root READMEs + the `demo/` recipe family for the unified `stack-*` skills + dev-as-peer.
- **Out:** nothing (closing milestone).

## Depends on
**M12 + M13 + M14** (documents their converged result). **Parallel with:** none (the closing milestone before `/developer-kit:close-release`).

## Open questions (resolve during build)
- Doc home — `corpus/ops/safety.md` vs `corpus/architecture/tooling-safety.md` (lean: `corpus/ops/safety.md`, ops-adjacent to seeding/snapshot/db-access).

## KB dependencies (read as contract)
- `corpus/ops/{snapshot-spec,seeding-spec,db-access}.md`, `corpus/architecture/security_compliance.md`
- the `rosetta-extensions/knowledge/` base (the repo's own KB)

## Delivers → corpus/ops/safety.md (new) + rosetta-extensions/knowledge/
- `corpus/ops/safety.md` (net-new — the read-side + write-side safety contract of the tooling).
- Updates the `rosetta-extensions/knowledge/` base to the v1.3 converged model.
