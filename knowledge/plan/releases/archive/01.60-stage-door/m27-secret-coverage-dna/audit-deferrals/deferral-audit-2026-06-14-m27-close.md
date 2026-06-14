---
title: "Deferral Audit — M27 (close-milestone Phase 1b)"
date: 2026-06-14
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals (M27 is the FIRST milestone of v1.6 — nothing inherited within-release to repeat).
- No aged-out items (the only milestone-scoped items were first-recorded this milestone; the inherited
  release-level backlog was re-signed <1 day ago at the v1.5 close).
- Every item has a clear fate decision.

## Summary
- Total deferrals in scope: 5 (2 milestone-scoped + 3 inherited release-level backlog)
- Single deferrals: 2 (both first-recorded in M27)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0

## Deferral Inventory

### Milestone-scoped (M27 decisions.md)
```yaml
- id: DEF-M27-01
  item: "Encrypted-zip (age/gpg) source ingestion support"
  origin_milestone: M27
  first_deferred_on: 2026-06-14
  last_seen_in: decisions.md M27-D3.5 ; spec-notes.md:32
  destination: "no committed owner — recorded v1 scope boundary ('attaches to a later milestone if a future need arises')"
  reason_recorded: "a new crypto dependency + key-management surface; no consumer needs it; plain dir/zip covers the M30 field-bake; module is deliberately stdlib-only (M27-D3.1) for a trivially-auditable values-blind surface"
  partial_attempted: no

- id: DEF-M27-02
  item: "Per-gene profile tag (vs default-graphql-profile scoping)"
  origin_milestone: M27
  first_deferred_on: 2026-06-14
  last_seen_in: decisions.md M27-D3.4 ; overview.md Open questions
  destination: "M28 (conditional — 'M28 may revisit if it wires non-default-profile bring-ups')"
  reason_recorded: "settled to graphql profile + waived-class device for v1; per-gene tag considered and rejected as not-needed-for-v1; the waived-class cleanly covers profile-gated keys"
  partial_attempted: no
```

### Inherited release-level backlog (orthogonal to secret provisioning)
```yaml
- id: DEF-M10-01
  item: "Cloud SnapshotStore / S3 media blob bytes"
  origin_milestone: M10 (v1.2)
  destination: "backlog / roadmap-vision (re-signed at v1.5 close 2026-06-14)"
  reason_recorded: "real-images-via-prod-public-links posture removed the user-facing sting; blob bytes not needed"
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
None. M27 is the first milestone of v1.6; there is no prior v1.6 milestone for an item to repeat across.
The 3 inherited items are single-release-deferred (v1.5 → backlog), re-signed once at the v1.5 close.

## Aging Check
- DEF-M27-01, DEF-M27-02: first-recorded TODAY (2026-06-14) — not aged on any trigger.
- DEF-M10-01, DEF-M21-01, M25-D9: last re-decided 2026-06-14 at the v1.5 `/close-release` deferral re-audit
  (same calendar day). Not aged (the ≥3-month / destination-milestone-closed / area-touched triggers do not
  fire — secret provisioning does not touch the snapshot/blob-byte area).

## Fate-1 Investigation

### DEF-M27-01 — "Encrypted-zip (age/gpg) source ingestion"
- **Fate-1 (land now, complete) feasible:** no
- **Why not:** a complete landing pulls in a crypto dependency + a key-management surface, directly contradicting
  the milestone's deliberate stdlib-only design (M27-D3.1) whose whole point is a trivially-auditable values-blind
  surface (no third-party code can ever see a value). No consumer needs it: the plain dir + plain-zip readers
  fully cover the M30 field-bake and the M28/M29 flows. A partial landing (a stub) is forbidden by the Fate-1
  no-partial rule.
- **Which fate:** this is NOT a tracked deferral with a live destination — it is a **documented v1 scope boundary**
  (M27-D3.5 states it "attaches to a later milestone if a future need arises"). There is no v1.6 milestone that
  owns or should own it (Fate 2/3 do not apply), and it is not release-scope-breaking (nothing in v1.6 needs it).
- **Recommendation: DROP** from the deferral ledger — it is a closed scope decision, not a pending item. It stays
  recorded as a known boundary in M27-D3.5 for any future need.

### DEF-M27-02 — "Per-gene profile tag"
- **Fate-1 (land now, complete) feasible:** the *open question* it came from is ALREADY landed (resolved this
  milestone): profile scoping is settled to the `graphql` profile + the waived-class device (M27-D3.4). The
  per-gene-tag VARIANT was consciously rejected as not-needed-for-v1.
- **Which fate:** the conditional residual ("M28 may revisit if it wires non-default-profile bring-ups") is
  **Fate 2** — naturally owned by M28's domain (M28 is the milestone that wires bring-ups). It is conditional,
  not committed scope, and M27-D3.4 already records the handoff, so **no M28 overview.md edit is required**.
- **Recommendation: LAND-NEXT (Fate 2, confirmed-covered)** — already noted in M27-D3.4; no plan edit.

### DEF-M10-01 / DEF-M21-01 / M25-D9 — inherited backlog
- **Fate-1 feasible:** no — all orthogonal to secret provisioning (snapshot/blob-byte/v1.5 surfaces).
- **Recommendation: KEEP-DEFERRED, no fresh action** — re-signed GREEN <1 day ago at the v1.5 close; carry
  unchanged. Re-fating them in an M27 close would be churn, not diligence.

## Recommendations
| Item | Fate verdict |
|---|---|
| DEF-M27-01 encrypted-zip | DROP (documented v1 boundary, not a pending deferral) |
| DEF-M27-02 per-gene profile tag | LAND-NEXT (Fate 2 — M28-owned-if-needed, no plan edit) |
| DEF-M10-01 / DEF-M21-01 / M25-D9 | KEEP-DEFERRED (re-signed <1 day ago at v1.5 close; carry) |

## Applied Changes
- No plan edits required. DEF-M27-01 is recorded as a closed scope boundary in M27-D3.5 (no ledger entry to
  carry forward). DEF-M27-02's Fate-2 handoff is already captured in M27-D3.4. The inherited items are
  unchanged (re-signed at the v1.5 close).
- This audit report is the audit trail.

## Blocking Items (require user decision)
None.
