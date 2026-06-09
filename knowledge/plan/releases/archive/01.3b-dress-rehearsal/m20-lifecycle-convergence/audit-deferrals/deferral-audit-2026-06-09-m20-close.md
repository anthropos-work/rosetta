---
title: "Deferral Audit — M20 close (release-wide M16→M20 pre-close sweep)"
date: 2026-06-09
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

M20 is the FINAL milestone of v1.3b "dress rehearsal", so this milestone-scope re-audit doubles as the
release-wide M16→M20 pre-close sweep before `/developer-kit:close-release`. Result: **0 new deferrals, 0 repeat,
0 chronic, 0 aged-out.** The sole open item across the whole release is the inherited, signed, cross-release
escape-hatch **DEF-M10-01** (→ v1.4), whose authority remains intact (all four aging triggers negative).

## Summary
- Total deferrals in scope: 1 (the single inherited cross-release escape-hatch)
- Single deferrals: 1 (DEF-M10-01, signed → v1.4)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0

## Deferral Inventory

### M20's own ledger — nothing deferred
- All four `overview.md` In-scope items landed Fate-1: set-dress chaining (#M20-D1), the demo `small-200`
  preset (#M20-D2), the atomicity contract (#M20-D3), the cold-start runbook (#M20-D4).
- **ISSUE-13 MCP-paging capture adapter is NOT a deferral** — it was a *spike* that resolved **document-only**
  (decision #M20-D4): the wired `postgres` MCP returns JSON rows, not COPY-format bytes, so an adapter would
  re-serialize COPY text for zero capability gain (the M9b-D9 reasoning). The sanctioned DSN-export / dump-restore
  cold-start path was authored in full (`corpus/ops/snapshot-cold-start.md`). A resolved spike that ships the
  documented path is a Fate-1 landing of the decision, not a punt.

### Release-wide scan (M16→M20) — no open within-release deferrals
- **M16-D8** — clerkenstein glossary stale-name fix: was a build-time boundary deferral (M16-D5), **reversed to
  Fate-1 and LANDED** at M16 close. Resolved, not open.
- **M17-D8** — the live docker-harness migrate-race *behavior* test: was M16's Fate-2 forward-route, **LANDED in
  M17** (reached its owning milestone). Resolved, not open.
- **M18 / M19** — frontend-port verification (M18→M19 Fate-2, landed in M19) and the true-zero-rebuild upstream
  PR (a documented OUT boundary — forbidden, edits platform repos — never a deferral). No open items.
- Code TODO/FIXME/HACK across M20-touched files (`dev-setdress.sh`, `up-injected.sh` + their tests, the rosetta
  docs): **none** (greps clean).

### The one open item (inherited, cross-release, signed)
```yaml
- id: DEF-M10-01
  item: "Cloud SnapshotStore backend + S3 media blob bytes (snapshot media carried as refs-only today)"
  origin_milestone: M10 (v1.2 "set dressing")
  first_deferred_on: 2026-06-07   # v1.2 close; signed escape-hatch (M9a-D4 + DEF-M10-01)
  last_seen_in: knowledge/plan/roadmap-vision.md (v1.4 staging, lines 28-30)
  destination: "v1.4 (signed RELEASE-SCOPE-DEFER)"
  reason_recorded: "S3 blob bytes + a cloud store change WHAT is transported and WHERE it is cached, not the
                    safety contract; refs-only is the deliberate current posture. A whole subsystem — out of
                    v1.3b's field-hardening scope by construction."
  partial_attempted: no
```

## Repeat-Deferral Patterns
None. DEF-M10-01 is a single signed cross-release escape-hatch, re-CONFIRMED (not re-deferred) at each v1.3b
milestone close — confirming a signed cross-release item is not the same as deferring it again. No item appears
in ≥2 milestones' deferral ledgers.

## Fate-1 Investigation

### DEF-M10-01 — "Cloud SnapshotStore + S3 media blob bytes"
- **Fate-1 (land now, complete) feasible:** no.
- **Why not:** a complete landing is a new subsystem (a cloud object-store `SnapshotStore` backend + blob-byte
  transport + the manifest/fidelity changes that ride it). That is categorically outside v1.3b "dress rehearsal",
  whose entire scope is field-hardening `/demo-up` with **tooling + docs only, zero platform-repo edits**. v1.3b
  never touched the snapshot-store/S3 area — confirmed by a file-level scan of the M16→M20 extensions history (no
  `SnapshotStore`/`store.go`/S3/blob file changed). Nothing in v1.3b incidentally unblocked it.
- **Fate:** escape-hatch (cross-release), already signed → v1.4. Home: `roadmap-vision.md`.

### Aging check (DEF-M10-01) — all four triggers negative
- Deferred across ≥2 milestones? **No** — single signed item; re-confirmed (≠ re-deferred) each close.
- Deferred ≥3 months ago? **No** — first deferred 2026-06-07; today 2026-06-09 (2 days).
- Destination milestone closed without landing? **No** — destination v1.4 is not yet opened.
- Feature area touched substantively by a later milestone? **No** — zero store/S3 touches across all of v1.3b.
- → **Authority intact. No fresh sign-off required.** It remains a clean signed v1.4 carrier.

## Recommendations
- **M20's own items:** all LAND-NOW (Fate-1), already landed in the milestone. No action.
- **DEF-M10-01 → KEEP-DEFERRED-WITH-SIGNOFF**, no fresh sign-off required (not aged, not repeat, context
  unchanged). Confirmed → v1.4. The release-close (`/developer-kit:close-release`) will carry it forward as the
  one signed escape-hatch, exactly as at the v1.3 close.

## Applied Changes
None required. No item converted, dropped, or newly escape-hatched. DEF-M10-01 unchanged in `roadmap-vision.md`.
This report records the release-wide pre-close verification.

## Blocking Items (require user decision)
None. 0 repeat-deferrals, 0 aged-out items, 0 chronic patterns. The audit is GREEN; close may proceed.
