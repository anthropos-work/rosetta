---
title: "Deferral Audit — M30 (close-milestone Phase 1b)"
date: 2026-06-14
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals: M30 introduces **zero** new deferrals. The field-bake is a proving + fix milestone;
  every gap it surfaced was landed Fate-1 (the 2 field bugs), and every non-passing secret-DNA gene is a
  documented **waived class** (waived-not-deferred — excluded from the gate denominator by design), not a
  punted scope item.
- No aged-out items: the inherited release-level backlog was re-signed 2026-06-14 at the v1.5 close (same
  calendar day) — no ageing trigger fires; M30 touches none of those surfaces (snapshot / blob-byte / v1.5).
- Every item has a clear fate decision.

## Summary
- Total deferrals in scope: 3 (all inherited release-level backlog; orthogonal to secret provisioning)
- Single deferrals: 0 new
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0
- New deferrals introduced by M30: 0

## Deferral Inventory

### Milestone-scoped (M30 decisions.md + progress.md + spec-notes.md)
None. The field-bake's outcomes are all Fate-1 landings or documented-correct waivers:
- **2 field bugs → LANDED Fate-1** (ext `m30/field-bake`, tag `stage-door-m30`):
  1. `sentinel/DB_CONNECTION` was critical/required but is compose-injected config (hardcoded `environment:`
     entry, never read from a `.env`) → reclassified `waived-config` + regression test pinning it.
  2. `up-injected.sh` only *checked* coverage, never *provisioned* + `preflight.sh` resolved REPO_ROOT one
     level too shallow (doubled `.agentspace/.agentspace/secrets` → demo gate silently skipped, exit 2) →
     both fixed: added the provision step + corrected the path to `EXT_ROOT/../..`.
- **Honesty residual = waived classes** (`waived-config`, `waived-aws-mount`, `waived-profile-gated`,
  `waived-optional`): excluded from the gate denominator by design — these are *not* deferrals. Each is a
  documented-correct reason a secret is not provisioned (compose-injected / AWS-mounted / profile-gated /
  optional). The Critical gate is 100%; zero critical short.
- No code TODO/FIXME/HACK in any M30-touched file (grep clean — rosetta skill docs + ext scripts).
- The one "defer" token in `spec-notes.md` ("DIRECTUS_TOKEN ... deferring to the injection override") is a
  runtime *behavior* (the fix16/17 non-rearm class — provision writes blank, the override strips), not a
  scope deferral.

### Resolved within-release (carried for completeness)
```yaml
- id: DEF-M27-01
  item: "Encrypted-zip (age/gpg) source ingestion support"
  status: DROPPED at M27 close (documented v1 scope boundary in M27-D3.5; no live destination)
- id: DEF-M27-02
  item: "Per-gene profile tag (vs default-graphql-profile scoping)"
  status: Fate-2 DISCHARGED at M28 close (the conditional never triggered — default graphql profile used)
```

### Inherited release-level backlog (orthogonal to secret provisioning — and to M30's prove+fix scope)
```yaml
- id: DEF-M10-01
  item: "Cloud SnapshotStore / S3 media blob bytes"
  origin_milestone: M10 (v1.2)
  destination: "roadmap-vision backlog (re-signed at v1.5 close 2026-06-14)"
  partial_attempted: no
- id: DEF-M21-01
  item: "(v1.5 carry — roadmap-vision backlog)"
  origin_milestone: M21 (v1.5)
  destination: "roadmap-vision backlog (re-signed at v1.5 close 2026-06-14)"
  partial_attempted: no
- id: M25-D9
  item: "(v1.5 carry — roadmap-vision backlog)"
  origin_milestone: M25 (v1.5)
  destination: "roadmap-vision backlog (re-signed at v1.5 close 2026-06-14)"
  partial_attempted: no
```

## Repeat-Deferral Patterns
None. No v1.6 item has been deferred across ≥2 milestones. DEF-M27-02 discharged; DEF-M27-01 dropped. The 3
release-level items are single-release-deferred (v1.5 → backlog), re-signed once at the v1.5 close.

## Aging Check
- M30 milestone-scoped: none to age.
- DEF-M10-01 / DEF-M21-01 / M25-D9: last re-decided 2026-06-14 at the v1.5 `/close-release` re-audit (same
  calendar day). Not aged — the ≥3-month, destination-milestone-closed, and area-touched triggers do not
  fire (M30 is the secret-provisioning field-bake; it does not touch the snapshot / blob-byte / v1.5 surfaces).

## Fate-1 Investigation

### The 2 field bugs M30 surfaced — LANDED Fate-1 (not deferred)
- **Fate-1 feasible:** yes — both fixed in full on ext `m30/field-bake`, tagged `stage-door-m30`, with a
  regression test for the sentinel reclassification + a corrected, shellcheck-clean preflight path. Tests
  green (Go `-race`/vet/gofmt clean; 99 demo-stack pytests pass). This is exactly the class a field-bake
  exists to catch + close in-milestone.
- **Recommendation:** LAND-NOW (done).

### DEF-M10-01 / DEF-M21-01 / M25-D9 — inherited backlog
- **Fate-1 feasible:** no — all orthogonal to secret provisioning (snapshot / blob-byte / v1.5 surfaces).
- **Recommendation:** KEEP-DEFERRED, no fresh action — re-signed GREEN 2026-06-14 at the v1.5 close; carry
  unchanged. (These will surface again at v1.6 `/close-release` for the release-level re-audit.)

## Recommendations
| Item | Fate verdict |
|---|---|
| The 2 field bugs (sentinel waive; provision-wiring + preflight path) | LAND-NOW (landed Fate-1 @ stage-door-m30) |
| Honesty-residual waived classes | N/A — waived-not-deferred (documented-correct, gate denominator excludes) |
| DEF-M10-01 / DEF-M21-01 / M25-D9 | KEEP-DEFERRED (re-signed at v1.5 close; carry to v1.6 close-release) |

## Applied Changes
- No plan edits required. M30 introduces no deferrals; the 2 surfaced bugs landed Fate-1; the within-release
  items are already resolved (DEF-M27-01 dropped, DEF-M27-02 discharged); the inherited release-level backlog
  is unchanged.
- This audit report is the audit trail.

## Blocking Items (require user decision)
None.
