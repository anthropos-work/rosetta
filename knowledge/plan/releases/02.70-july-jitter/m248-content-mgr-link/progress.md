# M248 — Progress

Section milestone. Checklist stub from the roadmap In-list.

## Sections

- [x] **Re-point the sim `ManagerResultPath` builder** (`content_manifest.go`) — routed by sim_type:
  NON-interview → `/sim/<slug>/<userId>/result/<sessionId>` (`owner.UserID`); INTERVIEW → its dedicated
  `/enterprise/activity-dashboard/interviews/<simId>/<membershipId>` route. **verify-interview resolved LIVE**
  (D3): the `/sim` interview manager report is flag/data-gated (renders "Coming Soon" on a demo), so interview
  KEEPS its route; only the non-interview family moves to `/sim`. Go tests + honesty gate GREEN.
- [x] **Update the graders + regenerate the manifest + docs** — `shapeFor` selects by sim_type; `manager-scored`
  keys on the SCORE (language-agnostic, collapse-proof); `manager-interview` = M236 activity-dashboard shape;
  `content-route-contract` + `content-result-page` unit specs re-pointed (174 unit specs GREEN, tsc clean);
  `presets/content-manifest.json` regenerated (honesty gate GREEN, 21 non-interview → /sim, 2 interview →
  activity-dashboard); `content-stories-spec.md` + `content-stories-routes.md` updated for the mixed routing.

## Live render-confirm (demo-2)
Warm content-stories sweep: **LANDED 43/47**. 18/21 non-interview managers land on `/sim` + both interview
managers land on activity-dashboard; **direct browser drives** confirm the `/sim` manager route renders the
full scored result (asmt 4516 · train 5406 · asmt-voice-fail 2981 chars, score present). Residual (not M248
code defects, → M254 fresh-seed re-confirm): **3** non-interview manager pages render a header-only shell at the
sweep's settle budget (per-session render state; the route itself is proven to render full) + **1** academy
player env failure (ant-academy `:23077` down on demo-2).

## Completeness Ledger

### Deferred
- **CARRY-M248-01 → M254 (Fate-3):** re-confirm the content-stories manager pairs land on the FRESH billion
  reset-to-seed. 3 non-interview manager `/sim` pages rendered a header-only shell on demo-2's (M246-era, warm)
  seed at a 20 s settle; the route is proven to render full results, so this is a per-session warmth/seed-data
  artifact to re-check on a cold fresh seed — not an M248 projection/grader defect. (Also the academy `:23077`
  env failure is a demo-2 host state, re-checked on billion.)

### Dropped
