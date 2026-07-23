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

## M248: Hardening

### Pass 1 — 2026-07-23
Section milestone; real code is in `rosetta-extensions` (gitignored). Scope manifest (M248 diff
`9678f5f^..c0ec407`, 2 source files): `stack-seeding/seeders/content_manifest.go` (the sim_type
`ManagerResultPath` router — co-located `content_manifest_test.go`) and
`stack-verify/e2e/lib/content-result-page.ts` (the grader — co-located `content-result-page.unit.spec.ts`
+ `content-route-contract.unit.spec.ts`).

**Assessment:** coverage across the three target areas is deep and live-verified.
- **sim_type routing branch** (non-interview → `/sim`, interview → activity-dashboard): exhaustively
  pinned in Go — `TestBuildContentProducts_SimulationProjection` (asmt/interview/hire specifics) +
  `TestBuildContentProducts_SeatAndSessionSingleSource` (per-session drift BOTH directions, with the
  interview-membership-not-userid + non-interview-not-activity-dashboard regression guards) — and in TS
  (`content-route-contract`: shape-by-sim_type, the CQ-1 slug-not-keyword manager twin, anti-fall-through,
  uuid-shape). No gap.
- **honesty gate / projection**: fail-closed drops, downgrade, no-host, no-owner, denominator pin,
  language-consistency teeth, `products:[]` wire-format — all present. No gap.
- **manager-scored-by-score grader**: one genuine mutation-surviving gap — see below.

**Tests added** (`content-result-page.unit.spec.ts`, +2 mutation-style):
- 1 the `hasSkills`-only acceptor: an EXPANDED manager render (Evaluated Skills section, NO N/100 score)
  lands on `hasSkills` alone. Every prior positive `manager-scored` test carried a score, so dropping
  `!hasSkills` from `if (!hasScore && !hasSkills)` would have survived them all (silent false-FAIL of a
  real expanded render). The manager twin of player-scored's "evaluated-skills section alone lands".
- 1 the 400-char floor (decision D2), distinct from player-scored's 300: a score-bearing sub-400 page
  does NOT land as manager, but the SAME page DOES land as player-scored — isolating the deliberate
  100-char divergence nothing sat in 300..400 to pin.

**Mutation-verified:** removing `!hasSkills` and collapsing `400→300` in the source each turn the matching
test red; reverted. Full unit suite **174 → 176 GREEN**, tsc clean, Go seeders GREEN, 3× flake-clean.

**Bugs fixed inline:** none (the branches were correct; only untested).

**Flakes stabilized:** none observed.

**Knowledge backfill:** none warranted — this pass PINS behavior already documented (decision D2 records
the score-based gate + 400 floor; the grader source comment already states "`hasSkills` is kept as an
additional acceptor for an expanded render"). No new invariant surfaced.

**rext consumption tag:** `july-jitter-m248-harden` @ `6e0ed2c` (pushed to origin, rung-zero verified).

### Stop condition
One pass. The full six-dimension scan found exactly one mutation-surviving branch (now closed); the Go
and honesty-gate/routing layers are already exhaustively + live-covered. Right-sized per the harden
mandate — did NOT force further passes on adequately-covered code.
