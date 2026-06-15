---
title: "Deferral Audit — M32 (close)"
date: 2026-06-15
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals. M32 is v1.7's SECOND and FINAL milestone; the only prior in-release milestone (M31)
  closed GREEN with zero open deferrals ([../../m31-mkcert-fapi-cert/audit-deferrals/deferral-audit-2026-06-15-m31-close.md](../../m31-mkcert-fapi-cert/audit-deferrals/deferral-audit-2026-06-15-m31-close.md)).
- Every item surfaced has a clear, already-recorded fate. No CHRONIC/DRIFT patterns; no AGED_OUT items.

## Summary
- Total deferrals in scope: 0 in-milestone punts + 5 inherited release-vision backlog items re-confirmed
- Single deferrals: 0
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out items: 0

## Deferral Inventory

**M32's own docs carry ZERO punted work.** The grep for `defer|postpone|later|out of scope|future milestone|
tracked for|follow-up|backlog|escape-hatch` across M32's `decisions.md` returned nothing; `overview.md` /
`progress.md` surfaced only two design-time scope boundaries, neither a punt:

```yaml
- id: DEF-M32-cand-01
  item: "the cert work"
  origin_milestone: M32 (overview.md:42 — Out:)
  reason_recorded: "Out: the cert (M31)"
  classification: SIBLING-MILESTONE BOUNDARY (not a deferral)
  why: the FAPI TLS cert is M31's delivered scope (closed + merged). M32 is the sibling studio-desk fix.
       Listed Out: only to mark the seam; nothing to land in M32.

- id: DEF-M32-cand-02
  item: "ant-academy demo liveness"
  origin_milestone: M32 (overview.md:42 — Out:)
  reason_recorded: "Out: ... ant-academy (backlog)"
  classification: RELEASE-LEVEL ROUTING decided at v1.7 design-roadmap (→ M33 / roadmap-vision backlog, repro-first)
  why: routed by the v1.7 design pass, not by M32. roadmap-vision.md §Unscheduled backlog records "M33 — ant-academy
       demo liveness (deferred from v1.7 design, 2026-06-15, repro-first)". Not M32's to land; not an in-release item.
```

**The close-time live Playwright smoke is NOT a deferral.** The smoke box was satisfied-by-composition at close
(M32-D5): `NODE_ENV=production` is pinned (regression test, mutation-checked 4 ways) → `isProduction=true` → the
production `sendFile` block (no dead-`:9100` 302), and that block covers every dev-block route (code-read,
M32-D1, "NO GAP"). Necessary + sufficient — the work is DONE, not punted. A fresh `/demo-up` re-demonstrates it
live on demand (operator action, not pending work). This mirrors M31-D7's accepted composition close.

### Inherited release-vision backlog (re-confirmed, all out of milestone scope)

The five items the orchestrator named live in `roadmap-vision.md` §"Unscheduled backlog (not a planned release)"
— cross-release / unscheduled by design, NOT in-release v1.7 milestone deferrals. Re-aged this pass:

```yaml
- id: M33   (ant-academy demo liveness) — re-signed at v1.7 design 2026-06-15, repro-first. NOT aged-out.
- id: M26   (self-contained-demo, orphaned ext branch) — awaits its OWN design-roadmap pass for a version. Unchanged; NOT aged-out (status is "unplaced", not "deferred-within-a-release").
- id: DEF-M10-01 (cloud SnapshotStore + S3 blob bytes) — gated on eu-west-1 S3 read access landing; re-signed v1.5 design. NOT aged-out (external dependency, not time).
- id: DEF-M21-01 (replayCmd conn-seam hermetic test) — LANDED at v1.5 close-release 2026-06-14 (survives the branch merge); pick up when the replay path next opens. Resolved-shaped, not pending.
- id: M25-D9 (dev-N taxonomy replay rc=4 migrate-ordering) — dev-only follow-up, orthogonal to the content-serve done-bar (GREEN). NOT aged-out.
```

None is touched by M32 (studio-desk single-port / `:9100` sweep), so the "feature area touched by a later
milestone" aging trigger does not fire. None has been deferred across ≥2 *in-release* milestones (they are
cross-release/unscheduled). All were re-signed at the v1.7 design-roadmap pass (2026-06-15) — fresh authority.

## Repeat-Deferral Patterns
None. M32 punts nothing of its own; the prior in-release milestone (M31) closed with zero open deferrals.
The inherited backlog items live in roadmap-vision (cross-release), so they cannot form an in-release
repeat-deferral pattern.

## Fate-1 Investigation
- DEF-M32-cand-01 (cert): Fate-1 N/A — M31's delivered, merged scope. Sibling boundary, no work to land.
- DEF-M32-cand-02 (ant-academy liveness): Fate-1 NO — separate area (M33), repro-first, no repro yet; landing
  in M32 would be net-new release scope. Already routed at v1.7 design (roadmap-vision).
- Live Playwright smoke: LANDED-by-composition this close (M32-D5). Nothing to fate.
- Inherited backlog (M33/M26/DEF-M10-01/DEF-M21-01/M25-D9): Fate-1 NO across the board — each is genuinely
  out of v1.7's demo-UI-hardening thesis (external gate / orphaned-effort needing its own design / dev-only
  orthogonal / already-landed). All correctly cross-release or unscheduled.

## Recommendations
- DEF-M32-cand-01 → no action (sibling boundary, M31 delivered).
- DEF-M32-cand-02 → no action (already routed to M33/roadmap-vision at v1.7 design).
- Live smoke → no action (resolved-by-composition this close, M32-D5).
- Inherited backlog → no action (all correctly cross-release/unscheduled in roadmap-vision; re-signed at v1.7 design).

(Outward-facing release-level git carry-over noted in state.md — push the v1.5/v1.6 ext tags to origin; the
orphaned `m26/self-contained-demo` + `wip/clerkenstein-browser-login` branches — is release/roadmap-vision
hygiene for the upcoming `/developer-kit:close-release`, NOT an M32 milestone deferral. Out of this audit's
milestone scope; flagged for the orchestrator / close-release.)

## Applied Changes
None. M32's docs carry no punted work; every candidate already carries its correct fate (overview Out: lines +
M32-D5 + roadmap-vision). No new task added, no decision rewritten, no roadmap edit needed.

## Blocking Items (require user decision)
None.
