# M248 — Progress

Section milestone. Checklist stub from the roadmap In-list.

## Sections

- [ ] **Re-point the sim `ManagerResultPath` builder** (`content_manifest.go:411-423`) to
  `/sim/<slug>/<userId>/result/<sessionId>` (`owner.UserID` on the ownerSlot; the per-`sim_type` kind
  branching collapses) — **verify-interview first** (rung-0: `/sim` route vs `/activity-dashboard/interviews`).
- [ ] **Update the graders + regenerate the manifest + docs** — the e2e grader `content-result-page.ts:459`
  shape + the `content-route-contract.unit.spec.ts` prefixes; `presets/content-manifest.json` regenerated via
  `stackseed --content-export` (honesty gate); `content-stories-spec.md` + `content-stories-routes.md`.

## Completeness Ledger

### Deferred

### Dropped
