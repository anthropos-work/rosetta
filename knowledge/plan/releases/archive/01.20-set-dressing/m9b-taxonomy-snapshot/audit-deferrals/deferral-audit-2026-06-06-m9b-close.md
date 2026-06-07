---
title: "Deferral Audit — M9b close"
date: 2026-06-06
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN (resolved by user fate decision — re-run 2026-06-06, close attempt 2)

- The one AGED_OUT / REPEAT deferral (the **offline pg_dump-FILE reader**, DEF-M9b-02b) has been
  **DROPPED by the user** (M9b-D9) — cut from the roadmap, NOT deferred. The drop is recorded in
  `decisions.md`, the source carry-forward (the M9a-retro note) has been cut so it cannot re-surface,
  and it is NOT seeded to `roadmap-vision.md` / v1.3. With the item RESOLVED (dropped-by-user), no
  unresolved repeat-deferral remains → GREEN. The companion correctness fix DEF-M9b-02a lands in
  Phase 7 (docs+code stop claiming an offline reader exists).
- _Prior verdict (attempt 1): RED — see "Original blocking analysis" below, retained for the trail._

## Original blocking analysis (attempt 1 — superseded by the user DROP above)

- One AGED_OUT / REPEAT deferral: the **offline dump-ingest reader** was the M9a-scoped boundary,
  carried forward to M9b "as part of the first real capture", and is closing in M9b without landing.
  Aging trigger fires (the milestone it was deferred *to* is closing without it) → a fresh user
  fate-decision is required before close proceeds.

## Summary
- Total deferrals in scope: 3 (the M9a → M9b carry-forward list) + 0 M9b-internal
- Single deferrals: 2 (both landed — see inventory)
- Repeat deferrals: 1 (the offline dump-ingest reader — REPEAT/AGED_OUT)
- Chronic patterns flagged: 0 (this is the item's second appearance, not a chronic ≥3 pattern, but
  it is AGED_OUT because its destination milestone is closing without it)

## Deferral Inventory

M9b carries no internal deferrals (all 7 sections checked; M9b-D1…D8 resolved; Q1/Q2/Q3 resolved to
D-numbers; **zero TODO/FIXME/HACK** in any M9b-touched Go source or test). The inventory is the
M9a→M9b carry-forward list (m9a retro.md §"Carried forward → M9b"):

```yaml
- id: DEF-M9b-01
  item: "Prove the framework on the real public taxonomy (capture + parent-scope + bulk-COPY + index rebuild)"
  origin_milestone: M9a
  destination: M9b
  status: LANDED — the taxonomy surface (10-table enumeration, public filters, parent-scope
    predicate, vector columns + REINDEX, FK replay order), the firewall parent-scope build, the
    fidelity measure, and the DAG node all landed + unit-proven (M9b-D1…D8). The live 2.1 GB
    capture is exercised via --dsn (a restored dump or safe primary read, per M9a-D3).

- id: DEF-M9b-02
  item: "Wire the offline dump-ingest READER (read an existing prod pg_dump file directly, no restore+DSN)"
  origin_milestone: M9a
  first_deferred_on: 2026-06-06 (M9a close — the scoped boundary, m9a retro §What-didn't)
  last_seen_in: stack-snapshot/cmd/stacksnap/main.go:181-189 (comment + the --dsn requirement)
  destination: M9b ("as part of the first real capture") — NOT LANDED
  reason_recorded: "the dump-format work belongs with the first real surface"
  partial_attempted: no — the --dump <path> flag selects the dump-ingest KIND but is then dropped;
    main.go still forces --dsn for every source, so dump-ingest needs a restored dump + DSN.

- id: DEF-M9b-03
  item: "Tag stack-snapshot-m9b after the final harden pass, not at build-end"
  origin_milestone: M9a (the one M9a close finding)
  destination: M9b close
  status: OPEN — owned by this close (Phase 11); tag will be set at the final hardened HEAD f9fabc3.
```

## Repeat-Deferral Patterns

### REPEAT / AGED_OUT: "offline dump-ingest reader"
- **First deferred:** M9a close, 2026-06-06 — explicit scoped boundary: "dump-ingest selects but
  still needs a DSN in M9a … the offline pg_dump reader is deferred to the M9b path."
- **Deferred again (de facto):** M9b — carried forward in m9a retro §"Carried forward → M9b" item 2
  ("wire the offline dump-ingest reader as part of the first real capture"); M9b implemented the
  taxonomy surface but NOT the offline file reader. The M9a-era comment + the --dsn requirement are
  unchanged at HEAD.
- **Current destination:** none recorded — it would silently drop at M9b close.
- **Time in limbo:** 2 milestones (M9a → M9b).
- **Aging trigger:** "the milestone it was deferred *to* (M9b) is closing without landing it" →
  AGED_OUT → blocking, fresh fate decision required.
- **Pattern flag:** not CHRONIC (2 appearances, not ≥3). DRIFT-adjacent: the reason each time is
  "belongs with the next/first-real surface," which has now run out of next-surfaces for the read
  path (M10 is Directus, M11 is recipes — neither is a skiller dump consumer).

## Fate-1 Investigation

### DEF-M9b-02 — "offline dump-ingest reader"
- **Fate-1 (land now, complete) feasible:** no (as a full feature) — implementing a real offline
  pg_dump-archive reader (parse custom/plain pg_dump, route per-table COPY streams through the
  firewall offline) is a meaningful new capability, not a slice. A partial "stub the reader" is a
  disguised deferral the three-fate rule rejects. Critically, it is an **optimization over a working
  path**, not a gap in M9b's delivered scope: the capture-source POLICY (dump-ingest as default
  precedence) IS landed; the operator drives dump-ingest today by restoring the dump and pointing
  `--dsn` at it (a sanctioned, documented path — the dump is "ingested" via restore-then-read).
- **The Fate-1 part that IS landed now (correctness, not the feature):** the doc/comment
  CONTRADICTION the deferral left behind — `source.go:9-17` claims "this build ships (1) dump-ingest"
  (offline) while `main.go:181-189` forces `--dsn`, and the CLI exposes a dead `--dump <path>` flag.
  Whatever the feature's fate, the docs+comments+CLI must stop claiming an offline reader exists.
  This is landed in Phase 7 of this close (see Applied Changes).
- **If no (the feature):** Fate options are Fate 3 (annotate a release-milestone) or escape-hatch
  (cross-release, v1.3 infra). Neither M10 (Directus content) nor M11 (recipes/presets) is a natural
  home for a skiller pg_dump-archive reader. The offline reader is infrastructure tooling that pairs
  with the v1.3 restore-from-snapshot / read-replica / cloud-store upgrades (all already v1.3, behind
  the same AWS-access gate). The honest fate is **escape-hatch → v1.3**, with the working restore+DSN
  path covering M9b/v1.2's need. This requires user sign-off (aged-out + escape-hatch).

## Recommendations

| ID | Item | Recommendation |
|----|------|----------------|
| DEF-M9b-01 | Prove framework on real taxonomy | NONE — landed. |
| DEF-M9b-02a | Doc/comment/CLI contradiction (offline reader claimed but absent) | **LAND-NOW** (Fate 1) — correctness fix in Phase 7: align `source.go` doc, `main.go` comment, remove the dead `--dump` flag, and fix `corpus/ops/snapshot-spec.md` to state the truth: **dump-ingest = restore-the-dump-into-Postgres-then-`--dsn`** (plus the safe-primary-read path). No offline file-reader — and (per the DROP) it is NOT a v1.3 upgrade either. |
| DEF-M9b-02b | Offline pg_dump-file reader (the feature) | **DROPPED-by-user** (M9b-D9, 2026-06-06) — cut from the roadmap, not deferred. Two working data-movement paths (safe primary read via `--dsn`; restore-then-`--dsn`) already cover the need; the direct file-reader adds no new capability and no reliable speed gain, and dropping it ends the repeat-defer pattern. NOT seeded to v1.3 / `roadmap-vision.md`. |
| DEF-M9b-03 | Tag after final harden | In progress — owned by this close, Phase 11. |

## Applied Changes (close attempt 2 — user fate decision applied)
- **DEF-M9b-02b — DROPPED by the user** (M9b-D9, 2026-06-06). Cut from the roadmap, not deferred. No
  `RELEASE-SCOPE-DEFER` decision, no `roadmap-vision.md` / v1.3 entry. The source carry-forward (the
  M9a-retro "Carried forward → M9b" note) has been **cut** so the item cannot re-surface in a future
  audit. The dead `--dump` CLI flag is removed and `source.go` / `snapshot-spec.md` are corrected
  (Phase 7) so docs+code no longer claim an offline file-reader exists.
- DEF-M9b-02a (the correctness fix) lands in Phase 7 — recorded as part of M9b-D9.

## Blocking Items — RESOLVED
1. **DEF-M9b-02b — the offline pg_dump-file reader.** ~~AGED_OUT + escape-hatch; user fate required.~~
   **RESOLVED: the user FATED it DROP** (M9b-D9, 2026-06-06). Cut entirely; restore-then-`--dsn` and
   the safe primary read cover v1.2's need. No remaining blocking items → verdict GREEN.
