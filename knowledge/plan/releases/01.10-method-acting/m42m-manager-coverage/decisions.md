# M42m — Decisions

Implementation decisions with rationale, numbered `M42m-D1`, `M42m-D2`, … . Cross-iter decisions live here;
per-iter detail lives in each `iter-NN/decisions.md`. The strategy-of-record `TOK-NN` entries also live here
(the milestone's strategy-evolution chain).

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M42m-D1 | The next-web Studio left-nav link (`STUDIO_URL`) is NOT env-rewritable to the demo-local studio-desk without a platform-source edit → RE-SCOPE TRIGGER (RESCOPE-1). | `core-js/constants/urls.ts:12` is a `NEXT_PUBLIC_NODE_ENV` ternary with no per-URL `NEXT_PUBLIC_STUDIO_URL` override (unlike `ACADEMY_URL`); the only knob (a global dev-flip) sends Studio to the wrong port `:9000` (demo studio-desk is `:39000`) AND breaks `WEB_APP_URL`/`HIRING_APP_URL`. Confirmed live (demo-3: `NEXT_PUBLIC_NODE_ENV=[]`, prod host baked). See iter-02/decisions.md D1. | 2026-06-25 |

## RESCOPE-1: the manager Studio left-nav escape is platform-bound — 2026-06-25

**Trigger:** coverage-protocol.md §"Re-scope trigger (the zero-edit line)" + the milestone overview's re-scope
clause — a 100%-(d)-blocking escape (escapes=139, ALL `studio.anthropos.work`) whose only clean fix is a
platform-repo edit (forbidden this release).

**The escape:** the baked `studio.anthropos.work` "Studio" left-nav outbound link the manager/enterprise nav
renders on every authenticated page (employee nav omits it → the employee gate had 0 escapes). This is the
exact prod-eject the user flagged in the live-demo review ("if I click Studio it brings me to production").

**Why it's platform-bound (the falsified rewrite hypothesis):** next-web's `STUDIO_URL`
(`packages/core-js/src/constants/urls.ts:12-15`) is a `NEXT_PUBLIC_NODE_ENV === 'development'` ternary
(`http://localhost:9000` | `https://studio.anthropos.work`) with **no `NEXT_PUBLIC_STUDIO_URL` per-URL
override** — unlike `ACADEMY_URL` (line 16-17, `process.env.NEXT_PUBLIC_ACADEMY_URL || …`), which is precisely
the override ant-academy's demo uses to rewrite its own Studio link. The demo next-web build bakes only 3
URL build-args + a gitignored `.env.local` pk overlay; it leaves `NEXT_PUBLIC_NODE_ENV` unset → the prod
branch (confirmed live: container `NEXT_PUBLIC_NODE_ENV=[]`, bundle carries `studio.anthropos.work`). The only
available knob — `NEXT_PUBLIC_NODE_ENV=development` in `.env.local` — is broad-and-wrong: it points Studio at
`:9000` (the demo studio-desk is on the OFFSET `:39000`) AND flips `WEB_APP_URL`/`HIRING_APP_URL`/pagination,
introducing NEW wrong-port links across other manager surfaces. No HTML-rewriting proxy exists in the demo
stack. So the sole clean fix is a 1-line platform edit to `urls.ts` — forbidden.

**The user's decision (one of):**
- **(a) Carve-out + disclose (lowest cost; the natural fit).** Treat the Studio link as a documented, disclosed
  external link — add a manager-vantage exception so the harness records it as a presenter-note ("do not click
  Studio live — it routes to production") instead of a gate escape. The protocol already has the
  presenter-notes mechanism (gate (d): "legitimate external links allowed but disclosed"); this widens it to
  cover a known, single, platform-bound nav link. The gate's (d) clause then reads "0 *closable* escapes +
  disclosed Studio". This addresses the user's concern (the presenter knows not to click it) without a
  platform PR.
- **(b) Upstream platform PR (the complete fix).** Add `process.env.NEXT_PUBLIC_STUDIO_URL ||` to `urls.ts:12`
  (a 1-line change mirroring `ACADEMY_URL`), out-of-band as a real next-web PR; then the demo injection sets
  `NEXT_PUBLIC_STUDIO_URL=http://localhost:39000` in the `.env.local` overlay (zero-edit, like the existing pk
  overlay) and the link becomes demo-local. This is the only path to a TRUE 0-escape manager gate.
- **(c) Pivot** — any other strategy the user prefers.

**Status:** awaiting user decision. The rest of TOK-01 (lines 2-4 — manifest route reconciliation + dashboard
populate + sample rules + calibration) is INDEPENDENT of this escape and remains valid next work; a future
`build-mstone-iters` call resumes there under TOK-01 once the re-scope is resolved.

## TOK-01: manager-coverage — reconcile-route + clear-escape + populate-dashboard + exhaust-frontier — 2026-06-25

**Tok type:** bootstrap (iter-01)

**Initial strategy:** Drive the M42e semantic harness as `dan-manager` against demo-3 and close the manager
gate (same (a)-(d) bar) in four leverage-ordered fix lines, all in rext (zero platform edits):
1. **Studio-link escape (highest leverage — clears 139 in one fix).** The baked `studio.anthropos.work`
   left-nav link renders on every authenticated manager page (the enterprise nav; employee nav omits it).
   Diagnose the rendered DOM: is the link host a `NEXT_PUBLIC_*_URL` the demo injection can rewrite to the
   local studio-desk offset port (`:39000`), or hardcoded in next-web? If env-configurable → rewrite in the
   demo injection (`gen_injected_override.py` / `up-injected.sh` build-args). If HARDCODED → **re-scope
   trigger** (escalate, record, do NOT edit the platform).
2. **Reconcile the manager manifest to the real route model + populate the dashboard.** The seeded manager
   `jump_to` is `/enterprise/workforce?tab=skills-verification` (NOT the manifest's `/workforce/*` guesses —
   confirmed in `stories.seed.yaml` + `test_cockpit.py` + the live sweep's `/enterprise/*` route prefix).
   Re-author `MANAGER_PAGES` to the real `/enterprise/workforce` tab-query route model (the dashboard is ONE
   tabbed route, not 5 sub-routes), then verify each M36 surface (verification funnel / teams / role-readiness
   / succession / mobility) renders REAL seeded data. The 6 M36 seeders already write the data + CORS is
   already wired (`CORS_EXTRA_ORIGINS`), so seed/serve-grant only what the live render proves empty.
3. **Sample rules + raise the cap for the TWO manager fan-outs** so the frontier EXHAUSTS honestly: the
   team-roster `/user/<id>`(+`/skills`+`/activities`) AND the per-activity
   `/enterprise/activity-dashboard/{ai-simulations,skill-paths,interviews}/<uuid>` drill-downs. This is a
   PRECONDITION (the cap-250 baseline sweep timed out at 25 min without them) — not just polish.
4. **Calibrate the manager manifest** floors/selectors against the live render (`calibrated:false → true`).

**Rationale:** the manager gate reuses the proven M42e harness + the proven employee fix surfaces
(seeding / injection / sample-rules / manifest-calibration), so the right opening move is leverage-ordered
fix lines, not a new mechanism. The escape line is a single high-value fix (139→0 if env-configurable). The
route reconciliation turns the unmeasured notReached=5 into a measurable content gate. The sample rules are a
hard precondition the baseline sweep proved. Persona already PASSES for dan-manager (identity machinery
generalizes for free — iter-23), so no persona work is expected.

**Strategy class:** new-direction (bootstrap — no prior strategy to compare).

**Distance-to-gate context:** the M42e iter-23 manager smoke-sweep: `escapes=139 notReached=5
frontier=CAPPED(+79) failingSections=0(over reached set only) personaFailures=0`. Gate = `(0,0,0,0) +
frontier EXHAUSTED`. The baseline (this iter, cap 250) confirmed the `/enterprise/` route prefix + a SECOND
fan-out (`/enterprise/activity-dashboard/.../<uuid>`) the smoke-sweep didn't reach; it timed out before
writing a report, so the iter-23 numbers stand as the baseline.

**Next-tik direction (iter-02):** start with the **Studio-link escape** (line 1). Log in as dan-manager on
demo-3, capture the rendered left-nav "Studio" `<a href>` host + how it's constructed (env var vs hardcoded).
If env-configurable, land the host-rewrite in the demo injection (rext), re-build the demo frontend, re-sweep
(scoped, with the new sample rules so it can finish) to confirm escapes→0. If hardcoded → re-scope trigger.
