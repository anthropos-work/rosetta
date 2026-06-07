---
title: "Deferral Audit — M10 close"
date: 2026-06-06
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals; no aged-out items; no chronic patterns. The one M10-internal deferral (the
  **S3 blob BYTES**, `MediaCaveat`) is a **first-appearance** item with a clear, already-recorded
  home → **Fate 2 → v1.3** (the cloud snapshot store / S3 backend seed in `roadmap-vision.md`,
  user note #4, 2026-06-06). The milestone floor (file refs + structure) **landed**; the bytes are a
  user-accepted operational add gated on S3 read access that is not wired here.
- All inherited v1.2 deferrals (M9a → M9b) were already RESOLVED at the M9b close (DEF-M9b-02b
  DROPPED by user; DEF-M9b-01 LANDED; DEF-M9b-03 tag mechanics completed). Nothing carries into M10.

## Summary
- Total deferrals in scope: 1 (M10-internal) + 0 unresolved inherited
- Single deferrals: 1 (the S3 blob bytes — first-appearance)
- Repeat deferrals: 0
- Aged-out: 0
- Chronic patterns flagged: 0

## Deferral Inventory

M10 carries no TODO/FIXME/HACK in any M10-touched Go source or test (grep clean). All 7 sections in
`progress.md` checked; M10-D1…D4 resolved; Q1/Q2/Q3 resolved to D-numbers. The single inventory item:

```yaml
- id: DEF-M10-01
  item: "Mirror the public Directus media BLOB BYTES into a per-stack-isolated store (the file REFS
    + ref-columns + 1,311 directus_files rows are captured/replayed; the S3 byte payloads are not)"
  origin_milestone: M10
  first_deferred_on: 2026-06-06 (M10 build — first appearance; no prior milestone touched media bytes)
  last_seen_in: spec-notes.md §"Media / blobs (M10-D4)"; stack-snapshot/directus/media.go
    (BlobBytesAvailable() == false)
  destination: v1.3 — the cloud snapshot store / S3-backend seed (roadmap-vision.md "v1.3 seeds")
  reason_recorded: "Blob bytes live in Directus's OWN S3 bucket. No confirmed S3 read access here →
    BlobBytesAvailable() is false. Until S3-read is wired, the per-stack Directus serves refs with a
    local-storage adapter + placeholder assets (a believable structural demo). Blob mirroring →
    per-stack-isolated bucket only, never the shared prod S3."
  partial_attempted: no — the file-REFS path is the milestone FLOOR (fully landed: ref columns,
    ReferencedFilesFilter, the 1,311 public directus_files rows). The bytes are a separate, cleanly
    bounded operational capability, not a half-built slice of the refs work.
```

### Inherited (v1.2 M9a → M9b) — all RESOLVED before M10, listed for the trail
- **DEF-M9b-01** (prove framework on real taxonomy) — LANDED at M9b.
- **DEF-M9b-02b** (offline pg_dump-file reader) — **DROPPED by user** (M9b-D9, 2026-06-06). The M9b
  audit cut the source carry-forward so it cannot re-surface. NOT seeded to v1.3. Confirmed absent
  from M10's ledger.
- **DEF-M9b-02a** (the doc/CLI correctness fix) — LANDED at M9b Phase 7.
- **DEF-M9b-03** (tag after final harden) — completed (`stack-snapshot-m9b` @ `55ee0e6`).

## Repeat-Deferral Patterns
None. DEF-M10-01 is a first-appearance item — no prior milestone deferred media bytes (M9a was the
framework, M9b was taxonomy; neither has a media surface). The aging/repeat logic does not fire.

## Fate-1 Investigation

### DEF-M10-01 — "S3 media blob bytes"
- **Fate-1 (land now, complete) feasible:** no — landing the bytes requires **confirmed S3 read
  access to Directus's own private bucket** (a credential/infra provisioning step) plus a
  per-stack-isolated mirror bucket to write into. Neither is wired in this environment, and provisioning
  S3 access is out of M10's content-snapshot scope. A partial "stub a byte-copier" would be a disguised
  deferral the three-fate rule rejects. Critically, the bytes are **not a gap in M10's delivered scope**:
  the structural floor (refs + ref-columns + the 1,311 public file rows + the local-storage/placeholder
  adapter path) IS landed and makes a believable structural demo — exactly what the user accepted upfront
  ("replay blobs too" with the explicit caveat that the bytes need S3 access not wired here).
- **The Fate-1 part that IS landed now:** the entire media-REFS path — `media.go`'s ref-column
  enumeration (uuid-typed `directus_files` references), `ReferencedFilesFilter()` (the 1,311 public-scoped
  rows), and the per-stack Directus's local-storage + placeholder-asset serving. `BlobBytesAvailable()`
  is the honest, tested gate that reports the bytes are not yet mirrorable.
- **If no (the bytes):** the natural home is the **v1.3 cloud snapshot store / S3 backend** seed —
  already explicitly recorded in `roadmap-vision.md` ("Cloud snapshot store … swaps the `SnapshotStore`
  backend for cloud object storage (S3) … Explicitly named as a v1.2 → v1.3 follow-up, user 2026-06-06
  note #4"). Blob byte-mirroring (read from Directus S3 → write to a per-stack-isolated bucket) is the
  same S3-access-gated infra class and pairs with that seed. This is **Fate 2** — a future-release home
  already exists in the vision doc; no new roadmap entry is invented at close. It is NOT escape-hatch:
  the user accepted it upfront as an operational add, and v1.2's scope is not broken by routing it (the
  structural floor delivers M10's "100% coverage" thesis — content is real, refs replayed).

## Recommendations

| ID | Item | Recommendation |
|----|------|----------------|
| DEF-M10-01 | S3 media blob bytes | **LAND-NEXT (Fate 2 → v1.3)** — confirmed-covered by the existing `roadmap-vision.md` "Cloud snapshot store / S3 backend" seed (user note #4, 2026-06-06). No roadmap edit needed; record the confirmation + backref in `decisions.md` (M10-D5). The structural floor (refs) landed in M10; the bytes are the S3-access-gated operational add. |
| Inherited M9a/M9b | (all 4) | NONE — resolved at M9b close (LANDED / DROPPED / tag-done). |

## Applied Changes
- DEF-M10-01 recorded as **Fate 2 → v1.3**, confirmed-covered by the existing `roadmap-vision.md`
  cloud-store/S3 seed. A confirmation decision (M10-D5) is added in `decisions.md` at Phase 7 with the
  backref to `roadmap-vision.md`. No `roadmap-vision.md` edit is required (the home already exists);
  no `RELEASE-SCOPE-DEFER` (this is Fate 2, not escape-hatch).
- No inherited-item changes (all were resolved at the M9b close).

## Blocking Items (require user decision)
None. DEF-M10-01 is a single, first-appearance Fate-2 item with an already-recorded v1.3 home. No
repeat deferrals, no aged-out items → verdict GREEN. No escape-hatch sign-off required.
