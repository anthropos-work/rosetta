---
title: "Deferral Audit — milestone M37 (close)"
date: 2026-06-23
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals. No chronic patterns. No aged-out items. Every M37 scope item landed (Fate 1) or is
  correctly owned by M38 (Fate 2, already-planned). No new deferral introduced at M37 close.

## Summary
- Total deferrals in scope: 0 (M37 introduced none)
- Single deferrals: 0
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out items: 0

## Deferral Inventory

**M37 own ledger** — grep of `decisions.md` / `progress.md` / `spec-notes.md` for
`defer|postpone|later|out of scope|future milestone|tracked for|follow-up|backlog|escape`: **zero matches.**
M37's two open questions both *resolved* (not deferred):
- **O11** (seat-switch mechanism: token-injection vs FAPI handshake) → RESOLVED to the parameterized
  FAPI handshake (server-authoritative); decision recorded in `decisions.md § O11`. Both options spiked
  per the overview's mitigation. Not a deferral.
- **KB-1** (browser-login handshake doc undocumented on `main`) → LANDED in full during M37's Phase-5 docs
  (folded into `architecture.md` as part of the wip-branch reconcile). Not a deferral.

**M37-touched rext code** (`registry.go`, `server.go`, `cmd/fake-fapi/main.go`, `alignment/cmd/multirun/`):
grep for `TODO|FIXME|HACK|XXX` in product files → **zero matches.**

**Inherited release backlog** (audited GREEN at the M34/M35/M36 closes): DEF-M10-01 (cloud SnapshotStore /
S3 blob bytes), DEF-M21-01..04 (prop-room residuals), M25-D9, M33. All orthogonal to the `clerkenstein`
ext section M37 touched — unchanged, no repeat-defer, none re-touched by M37.

## Repeat-Deferral Patterns
None.

## Fate-1 Investigation

The only items adjacent to M37 that could *look* like deferrals are M37's explicit **Out:** scope — the
presenter cockpit, the "login as"/"jump to" UI, the literal live browser seat-switch render, and the
seeder-side `roster export` producer. Each is **Fate 2 (already-owned by M38)**, not a deferral:

### M37-OUT-1 — "the presenter cockpit panel + [Login as]/[Jump to] UI + live seat-switch render"
- **Fate-1 (land now, complete) feasible:** no — it is a different ext section (`demo-stack`, a standalone
  served panel on an offset port) and a different deliverable (a UI that *drives* the capability M37 ships).
- **Which fate applies:** **Fate 2.** M38's `overview.md` `In:` list explicitly owns it ("the panel; … login-as
  wired to M37; jump-to"). M37 ships the Clerkenstein *consumer* (registry + active-seat selection + a golden
  roster fixture); the cockpit that drives the handshake is M38's domain. Confirmed-covered, no plan edit needed.

### M37-OUT-2 — "the seeder-side roster-export producer (FAKE_FAPI_ROSTER feed)"
- **Fate-1 (land now, complete) feasible:** no — by the ARCH decision in `decisions.md`, Clerkenstein owns the
  *consumer* (the `Registry` + roster-JSON contract + a golden roster fixture matching the stories preset). The
  *producer* that exports the roster from the seeder's own derivation is the demo-tooling/M38 integration seam.
- **Which fate applies:** **Fate 2.** Owned by M38 (the cockpit reads `stack.stories.yaml` and wires login-as
  to M37 — the roster export is part of that integration). Confirmed-covered.

## Recommendations
- M37-OUT-1 → **LAND-NEXT (Fate 2)** — confirmed owned by M38; no plan edit.
- M37-OUT-2 → **LAND-NEXT (Fate 2)** — confirmed owned by M38; no plan edit.
- Inherited backlog (DEF-M10-01 / DEF-M21-01..04 / M25-D9 / M33) → unchanged; orthogonal to M37; already
  GREEN-audited at M34/M35/M36 closes.

## Applied Changes
None required — no deferral introduced, no repeat pattern, no aged-out item. The two Out-scope items were
already correctly Fate-2-owned by M38 at design time (recorded in M37 `overview.md § Out` + M38 `overview.md
§ In`); this audit confirms that ownership rather than mutating any plan.

## Blocking Items (require user decision)
None.
