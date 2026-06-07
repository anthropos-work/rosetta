---
title: "Deferral Audit — release (v1.2 close)"
date: 2026-06-06
scope: release
invoked-by: close-release
release: v1.2 "set dressing" (01.20-set-dressing)
---

## Verdict
**GREEN**

Release-wide re-audit across all 4 milestones (M9a → M9b → M10 → M11). Every deferral raised during v1.2
has reached a terminal fate:

- **Exactly one item remains open** — **DEF-M10-01** (S3 media blob BYTES), fated **Fate 2 → v1.3** at the
  M10 close, confirmed-covered by the existing `roadmap-vision.md` "Cloud snapshot store / S3 backend" seed
  (user note #4, 2026-06-06). Re-confirmed unchanged at M11 close and here. This is the only expected open
  item; it does not block the release.
- The one **REPEAT/AGED_OUT** item that ever fired (the offline pg_dump-FILE reader, DEF-M9b-02b) was
  **DROPPED by the user** (M9b-D9) — cut entirely, source carry-forward severed, not re-surfacing, not
  seeded to v1.3. Verified absent from prod source (`--dump` flag gone) and pinned-gone by a regression test
  (`TestDroppedDumpFlagStaysGone`).
- **No repeat-deferral, no chronic pattern, no unresolved aged-out item** carries into the release close.
- **Zero TODO/FIXME/HACK** in any milestone-touched Go source across `stack-snapshot` + `stack-seeding`.

## Summary
- Total deferrals raised across v1.2: **7** (DEF-M9b-01/02a/02b/03 + DEF-M10-01 + the M9a-design Out-list
  scope-partition items + the M9a-D4 cloud-store escape-hatch)
- Open at release close: **1** (DEF-M10-01 — Fate 2 → v1.3)
- Resolved (LANDED / DROPPED / tag-done): **6**
- Repeat deferrals unresolved: **0** (the one repeat, DEF-M9b-02b, was DROPPED)
- Chronic patterns flagged: **0**
- Aged-out unresolved: **0** (the one aged-out, DEF-M9b-02b, was user-fated DROP)
- New deferrals introduced by M11: **0**
- Code TODO/FIXME/HACK in milestone-touched source: **0**

## Deferral Inventory (release-wide, terminal state)

```yaml
# --- M9a (framework) ---
- id: DEF-M9a-scope-1..3        # taxonomy / Directus content / recipes+presets
  origin_milestone: M9a (overview Out:)
  fate: Fate 2 (in-release scope partition) → LANDED by M9b / M10 / M11 respectively
  status: RESOLVED — each landed in its owning milestone's In: list (the M7a→M7c precedent)

- id: DEF-M9a-D4               # cloud/S3 SnapshotStore backend
  origin_milestone: M9a-D4 / M9a-Q4→D
  fate: KEEP-DEFERRED-WITH-SIGNOFF (escape hatch) → v1.3
  status: OPEN-BY-DESIGN — user-signed at design (2026-06-05/06); SnapshotStore interface seam built;
    home recorded in roadmap-vision.md. Same S3-access-gated class that DEF-M10-01 rides.

# --- M9b (taxonomy) ---
- id: DEF-M9b-01               # prove framework on real taxonomy
  fate: LANDED (M9b-D1…D8)
  status: RESOLVED

- id: DEF-M9b-02a              # doc/CLI correctness fix (offline-reader lie)
  fate: LAND-NOW (Fate 1)
  status: RESOLVED — landed M9b Phase 7 (source.go/main.go/README/snapshot-spec corrected)

- id: DEF-M9b-02b             # offline pg_dump-FILE reader (the feature) — the ONLY repeat/aged-out
  fate: DROP (user, M9b-D9)
  status: RESOLVED — cut from roadmap, source carry-forward severed, NOT seeded to v1.3.
    Verified: --dump absent from prod source; TestDroppedDumpFlagStaysGone pins it gone.

- id: DEF-M9b-03              # tag stack-snapshot-m9b after final harden
  fate: LAND-NOW (process)
  status: RESOLVED — tag @ 55ee0e6

# --- M10 (Directus content) — the one OPEN item ---
- id: DEF-M10-01
  item: "Mirror the public Directus media BLOB BYTES into a per-stack-isolated store (file REFS +
    ref-columns + 1,311 directus_files rows ARE captured/replayed; the S3 byte payloads are not)"
  origin_milestone: M10 (first appearance, 2026-06-06)
  last_seen_in: stack-snapshot/directus/media.go (BlobBytesAvailable()==false, MediaCaveat); M10-D5
  fate: Fate 2 (confirmed-covered) → v1.3
  destination: roadmap-vision.md "Cloud snapshot store / S3 backend" seed (user note #4, 2026-06-06)
  status: OPEN — expected; structural floor (refs) landed; bytes are the S3-access-gated operational add.
  partial_attempted: no — refs path is the FLOOR (fully landed); bytes are a clean separate add.

# --- M11 (recipes/presets/corpus) ---
# none — every section landed Fate-1; zero deferrals introduced.
```

## Repeat-Deferral Patterns
**One ever fired, now resolved.** The offline dump-ingest reader (DEF-M9b-02 → 02b) appeared across M9a→M9b
(2 milestones, DRIFT-adjacent: each defer reasoned "belongs with the next/first-real surface," which ran out
of next-surfaces). It tripped the AGED_OUT trigger at M9b close (destination milestone closing without it) →
RED at M9b close attempt 1 → user FATED it **DROP** (M9b-D9). Not chronic (2 appearances, not ≥3). With the
DROP applied and the source carry-forward severed, the pattern is dead and cannot re-surface. No other group
reaches size ≥ 2.

## Fate-1 Investigation (the one open item)

### DEF-M10-01 — "S3 media blob bytes"
- **Fate-1 (land now, complete) feasible:** no. Landing the bytes requires confirmed S3 read access to
  Directus's own private bucket + a per-stack-isolated mirror bucket — neither wired in this environment, and
  provisioning S3 access is outside v1.2's content-snapshot scope. A "stub a byte-copier" partial is a
  disguised deferral the three-fate rule rejects. Critically, the bytes are **not a gap in v1.2's delivered
  scope**: the structural floor (refs + ref-columns + the 1,311 public `directus_files` rows + the
  local-storage/placeholder adapter) IS landed and delivers M10's 100%-coverage thesis.
- **Fate applied:** **Fate 2 (confirmed-covered) → v1.3.** The v1.3 "Cloud snapshot store / S3 backend" seed
  (roadmap-vision.md, user note #4) owns the same S3-access-gated infra class; the media bytes ride that swap.
  No new roadmap entry invented; NOT escape-hatch (v1.2 scope is not broken). `BlobBytesAvailable()` is the
  honest, tested `false` gate; flipping it once S3 access is wired is a single switch.

## Recommendations
| ID | Item | Recommendation |
|----|------|----------------|
| DEF-M10-01 | S3 media blob bytes | **LAND-NEXT (Fate 2, v1.3)** — no action needed; destination seed exists + is unchanged. Re-confirmed at release close. |
| DEF-M9a-D4 | Cloud/S3 SnapshotStore backend | **KEEP (escape hatch, v1.3)** — user-signed at design; interface seam in place; same v1.3 home as DEF-M10-01. No action. |
| All others (6) | — | NONE — resolved (LANDED / DROPPED / tag-done) at their owning milestone closes. |

## Applied Changes
None required. The single open item (DEF-M10-01) retains its M10 fate and its recorded v1.3 destination; this
release audit re-confirms (does not re-fate) it. The DEF-M9a-D4 cloud-store escape-hatch retains its
design-time sign-off and shares the same v1.3 destination. No `roadmap-vision.md` edit, no new decision, no
`RELEASE-SCOPE-DEFER` needed (the homes already exist).

## Blocking Items (require user decision)
**None. GREEN.** No repeat-deferral without resolution, no aged-out item without a fresh fate, no chronic
pattern. close-release proceeds.
