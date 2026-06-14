---
title: "Deferral Audit — M28 (close-milestone Phase 1b)"
date: 2026-06-14
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals (M28 introduces no new deferral; the one inherited within-release item, DEF-M27-02,
  resolved cleanly this milestone rather than repeating).
- No aged-out items (the inherited release-level backlog was re-signed 2026-06-14 at the v1.5 close — same
  calendar day; no ageing trigger fires).
- Every item has a clear fate decision.

## Summary
- Total deferrals in scope: 4 (1 inherited within-release + 3 inherited release-level backlog)
- Single deferrals: 1 (DEF-M27-02, resolved this milestone)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0
- New deferrals introduced by M28: 0

## Deferral Inventory

### Inherited within-release (from M27)
```yaml
- id: DEF-M27-02
  item: "Per-gene profile tag (vs default-graphql-profile scoping)"
  origin_milestone: M27
  first_deferred_on: 2026-06-14
  last_seen_in: M27 decisions.md M27-D3.4 ; M27 close audit (Fate-2, M28-owned-if-needed)
  destination: "M28 (conditional — revisit IFF M28 wires non-default-profile bring-ups)"
  reason_recorded: "settled to graphql profile + waived-class device for v1; per-gene tag rejected as not-needed-for-v1"
  partial_attempted: no
```

### Inherited release-level backlog (orthogonal to secret provisioning)
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

M28's own `decisions.md` carries M28-D1/D2/D3 — none is a deferral. M28-D3 ("DIRECTUS_TOKEN non-rearm:
write BLANK … defer to the override") uses "defer" in the sense of runtime deference to the injection
override, not a scope deferral; it is a fully-landed, test-pinned design decision (no pending work).

## Repeat-Deferral Patterns
None. The only within-release inherited item (DEF-M27-02) did not repeat — it resolved (see Fate-1 below).
The 3 release-level items are single-release-deferred (v1.5 → backlog), re-signed once at the v1.5 close.

## Aging Check
- DEF-M27-02: first-recorded 2026-06-14 (yesterday's-same-day); resolved this milestone — not aged.
- DEF-M10-01 / DEF-M21-01 / M25-D9: last re-decided 2026-06-14 at the v1.5 `/close-release` re-audit (same
  calendar day). Not aged: the ≥3-month, destination-milestone-closed, and area-touched triggers do not fire
  (secret provisioning does not touch the snapshot / blob-byte / v1.5 surfaces).

## Fate-1 Investigation

### DEF-M27-02 — "Per-gene profile tag"
- **Fate-1 (resolved this milestone):** the conditional was "M28 revisits IFF it wires non-default-profile
  bring-ups." M28 wired the pre-flight into `/dev-up` + `/demo-up` on the **default `graphql` profile**; it did
  NOT introduce a non-default-profile bring-up, so the conditional **did not trigger**. Profile scoping was
  reused unchanged from M27 (default `graphql` + the `waived-profile-gated` class) — M28 progress.md item
  "Profile-scoping decision settled + implemented" confirms this.
- **Which fate:** the conditional Fate-2 is **discharged** — the triggering condition never occurred, so there
  is no residual work. The per-gene-tag variant remains a documented, consciously-rejected v1 option (M27-D3.4)
  available if a future milestone ever wires a non-default-profile stack.
- **Recommendation: CLOSE (Fate-2 discharged)** — no plan edit, no carry-forward.

### DEF-M10-01 / DEF-M21-01 / M25-D9 — inherited backlog
- **Fate-1 feasible:** no — all orthogonal to secret provisioning (snapshot / blob-byte / v1.5 surfaces).
- **Recommendation: KEEP-DEFERRED, no fresh action** — re-signed GREEN 2026-06-14 at the v1.5 close; carry
  unchanged. Re-fating them in an M28 close would be churn, not diligence.

## Recommendations
| Item | Fate verdict |
|---|---|
| DEF-M27-02 per-gene profile tag | CLOSE (Fate-2 discharged — conditional never triggered) |
| DEF-M10-01 / DEF-M21-01 / M25-D9 | KEEP-DEFERRED (re-signed at v1.5 close; carry) |

## Applied Changes
- No plan edits required. DEF-M27-02's conditional resolved by non-triggering (M28 used the default `graphql`
  profile); recorded here as discharged. The inherited release-level items are unchanged.
- This audit report is the audit trail.

## Blocking Items (require user decision)
None.
