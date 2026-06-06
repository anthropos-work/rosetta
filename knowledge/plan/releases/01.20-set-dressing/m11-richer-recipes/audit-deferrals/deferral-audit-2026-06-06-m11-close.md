---
title: "Deferral Audit — milestone (M11 close)"
date: 2026-06-06
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals; no aged-out items; no chronic patterns.
- M11 itself added **zero** deferrals — every section landed Fate-1 (presets, recipes, the
  `/demo-snapshot` skill, corpus cross-links, the §5 `stacksnap --help` hygiene fix M11-D4).
- The single open v1.2 deferral — **DEF-M10-01** (S3 media blob BYTES / `MediaCaveat`) — was
  fated at the M10 close (**Fate 2 → v1.3**, GREEN) and is **unchanged** at M11: M11 touched no
  media surface (it is the docs/discoverability layer), and its recipe/preset wording honors the
  ref-floor-vs-bytes boundary. Destination seed exists in `roadmap-vision.md` ("Cloud snapshot
  store / S3 backend", user note #4, 2026-06-06).
- All earlier inherited v1.2 deferrals (M9a → M9b) were RESOLVED at the M9b close (DEF-M9b-02b
  DROPPED by user; DEF-M9b-01 LANDED; tag mechanics completed). Nothing chronic carries forward.

## Summary
- Total deferrals in scope: 1 (inherited; carried unchanged from M10)
- Single deferrals: 1 (DEF-M10-01 — first-appearance at M10, fated Fate-2 there)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0
- New deferrals introduced by M11: 0

## Deferral Inventory

```yaml
- id: DEF-M10-01
  item: "Mirror the public Directus media BLOB BYTES into a per-stack-isolated store (file REFS +
    ref-columns + 1,311 directus_files rows ARE captured/replayed; the S3 byte payloads are not)"
  origin_milestone: M10
  first_deferred_on: 2026-06-06 (M10 build — first appearance)
  last_seen_in: m10/spec-notes.md §"Media / blobs (M10-D4)"; stack-snapshot/directus/media.go
    (BlobBytesAvailable() == false). NOT re-touched in M11.
  destination: v1.3 — the cloud snapshot store / S3-backend seed (roadmap-vision.md "v1.3 seeds")
  reason_recorded: "Blob bytes live in Directus's OWN S3 bucket; no confirmed S3 read access here →
    BlobBytesAvailable() is the honest tested false gate. Refs + structure are the milestone FLOOR
    and landed; bytes are a separate S3-access-gated infra add, same class as the cloud store."
  partial_attempted: no — refs path is the FLOOR (fully landed); bytes are a clean separate add.
```

## Repeat-Deferral Patterns
None. DEF-M10-01 is a single first-appearance item, fated at first appearance. M11 introduced no
new deferrals and re-touched no deferred surface. No group reaches size ≥ 2; aging/repeat logic
does not fire.

## Fate-1 Investigation

### DEF-M10-01 — "S3 media blob bytes"
- **Fate-1 (land now, complete) feasible:** no — and not in M11's domain. M11 is the
  recipes/presets/corpus discoverability milestone; it ships no capture/replay code and no media
  surface. Landing the bytes still requires confirmed S3 read access to Directus's own private
  bucket + a per-stack-isolated mirror bucket — neither wired in this environment, and provisioning
  S3 access is out of v1.2 scope.
- **Fate applied:** Fate 2 (confirmed-covered) → v1.3. The v1.3 "Cloud snapshot store / S3 backend"
  seed in `roadmap-vision.md` owns the same S3-access-gated infra class; the media bytes ride that
  swap. No new roadmap entry invented; not escape-hatch (v1.2's 100%-coverage thesis is delivered by
  the structural floor — refs + replayed public templates). Confirmed at M10 close; re-confirmed
  unchanged here.

## Recommendations
- **DEF-M10-01 → LAND-NEXT (Fate 2, v1.3)** — no action needed; destination seed already exists and
  is unchanged. Re-confirmed at M11 close.

## Applied Changes
None required. M11 introduced no deferrals; the one inherited item retains its M10 fate and its
recorded v1.3 destination. This audit re-confirms (does not re-fate) DEF-M10-01.

## Blocking Items (require user decision)
None. GREEN.
