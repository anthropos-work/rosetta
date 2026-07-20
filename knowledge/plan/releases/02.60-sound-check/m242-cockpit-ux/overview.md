---
milestone: M242
slug: cockpit-ux
version: v2.6 "sound check"
milestone_shape: section
status: planned
created: 2026-07-20
last_updated: 2026-07-20
depends_on: M240 + M241
delivers: corpus/ops/demo/cockpit-spec.md + corpus/ops/demo/content-stories-spec.md
---

# M242 — cockpit UX

**Status:** `planned`  ·  **Shape:** `section`  ·  **Complexity:** medium  ·  **Depends on:** M240 + M241

## Goal
The Content-stories tab reads clearly and the heroes are legible by role.

## Scope
### In
- **(1) row layout** — regroup by requirement tuple `(sim_type, modality)` → `target | passed login options | not-passed login options` on one row (render-only; fields exist).
- **(2) tab selector** — move into the white header, right, vertically centered (restructure `cockpit.py` header to flex; **preserve the byte-identical-when-no-content-manifest invariant**).
- **(3) hero icon bg by user-type** (manager = orange / employee = indigo, reuse the badge palette; derive a candidate color = `is_hiring && vantage != manager`).
- Extend the **cockpit specs**.

### Out
- Any data / seed change (M240 / M241).
- Platform edits.

## Open questions
- None blocking.

## Delivers
`corpus/ops/demo/cockpit-spec.md` + `corpus/ops/demo/content-stories-spec.md` (the row-regroup + header layout + role-color).

## Full design
See `knowledge/plan/roadmap.md` § Active — v2.6 "sound check" for the authoritative milestone design + the release-level decisions/risks (design spec: `releases/02.60-sound-check/design-notes.md`).
