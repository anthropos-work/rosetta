**Type:** tik (TOK-01 lines 2-4 — route reconcile + dashboard populate + fan-out sample rules + calibrate)

# iter-04 — reconcile the manager route + populate the dashboard + exhaust the frontier + calibrate

## Phase A — Sweep / diagnose (measure)
Three live diagnostic probes as `dan-manager` on demo-3 discovered the real route model + the per-surface
content state (the baseline is the iter-23 smoke + iter-01 cap-250 timeout: escapes=139[fixed],
notReached=5, frontier CAPPED):
- The **real route surface** is `/enterprise/*`, and the Workforce dashboard is **ONE tabbed SPA route**
  (`/enterprise/workforce`; the `?tab=` query is ignored). The manifest's `/workforce/*` guesses matched
  nothing → the notReached=5.
- **The M36 dashboard already renders rich, real data** (221 members, 493 mapped, 262 verified / 53.1%
  coverage, 19 cards, 67 chart SVGs) — the populate work narrowed to the ONE empty page:
  `/enterprise/organization-feedback` ("No data" despite 103 seeded feedback rows).
- The two manager fan-outs confirmed: `/user/<uuid>` (team roster) + `/enterprise/activity-dashboard/.../<uuid>`.

## Phase B — Triage
- The notReached=5 = a route-model error (D1) → reconcile the manifest.
- The org-feedback "No data" = an inserted-but-invisible seed bug — the page JOINs feedback to the app mirror
  `local_jobsimulation_sessions` (D3), which the population sessions lack → seed the mirror (`stack-seeding`).
- The frontier CAP = no manager fan-out sample rules + the manager ALSO links the Library (D4) → add 4
  sample rules (2 manager fan-outs + the 2 inherited library families).

## Phase C — Fix (rext only; zero CANONICAL platform edits)
- **`stack-seeding/seeders/feedback.go`** — the FeedbackSeeder now also writes a `local_jobsimulation_sessions`
  mirror per feedback session (reconstructed from the same deterministic key; score / start-end / status
  coherent with the source population session), so the org-feedback JOIN resolves. + a `feedback_test.go`
  joinable-mirror test. gofmt clean, full seeders suite green.
- **`stack-verify/e2e/lib/coverage-manifest.ts`** — `MANAGER_PAGES` re-authored from the wrong `/workforce/*`
  to the real `/enterprise/*` routes (6 pages), sections + floors calibrated against the live render,
  `calibrated:true`. `MANAGER_MANIFEST.seedPaths` updated.
- **`stack-verify/e2e/tests/coverage.spec.ts`** — `MANAGER_SAMPLE_RULES` (the 2 fan-outs + the 2 library
  families) + vantage-aware sample-rule selection.
- **`stack-verify/e2e/run-coverage.sh`** — the manager expected-org default = "Cervato Systems" (Dan's org).
- **Re-applied to demo-3**: built `stackseed` from the authoring copy, additive re-seed → the mirror landed
  (Cervato `local_jobsimulation_sessions` 42→145, +103; all 103 feedback rows joinable, was 0). LIVE re-probe:
  org-feedback "No data" → "103 sessions / 70% positive / 59% pass / 63% avg" + 21 real rows.

## Phase D — Re-sweep (the authoritative measure)
The authoritative manager sweep (cap-120, gated, dan-manager) on the re-seeded demo-3:
```
reachable=70/120 failingSections=0 personaFailures=0 escapes=0 notReached=0 frontier=EXHAUSTED → GATE MET
PERSONA: role-skills-coherence=ok avatar-consistency=ok org-identity=ok
```
- All 6 `/enterprise/*` manifest pages PASS every section (workforce funnel 5112 chars + 19 cards + 10 tabs;
  members 2352; assignments 1911; activity-dashboard 2409; **organization-feedback 3675 + 21 session rows**;
  settings 574 [exception]).
- 4 fan-out families sampled-out + disclosed (28 users / 64 activity-drilldowns / 72 sims / 22 skill-paths).
- 1 presenter note (the LinkedIn-import help citation, allowed). 1 documented exception (`/enterprise/settings`).
- Persona PASS (Dan Rossi role↔skills, real-photo avatar menu==profile, Cervato Systems + logo).

(The first authoritative sweep CONFIRMED the diagnosis: the OLD manager rules didn't sample the library, so
the manager crawl exploded into /sim+/skill-path and timed out — corrected by the 4-rule superset.)

## Close — 2026-06-26

**Outcome:** the MANAGER semantic gate is **MET** on the re-seeded demo-3 — route reconciled (notReached 5→0),
dashboard fully populated (the one empty page, org-feedback, fixed via the seeder mirror), frontier EXHAUSTED
(the 2 fan-outs + 2 library families sampled), manifest calibrated (`calibrated:true`). `(failingSections,
escapes, personaFailures, notReached) = (0,0,0,0)`, `cappedAtFrontier=false`.
**Type:** tik
**Status:** closed-fixed (every planned line of TOK-01 (2-4) landed + the gate fired on the authoritative sweep)
**Gate:** MET
**Phase 5 grading:** (1) gate-met: y — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-1 (gate-met)
**Decisions:** D1 (route model), D2 (dashboard already real), D3 (the org-feedback mirror fix), D4 (sample
rules + library inheritance), D5 (settings exception) — see iter-04/decisions.md.
**Side-deliverables:** none.
**Routes carried forward:** none (the gate is met; the authoritative-FRESH-demo-up reproduction is a close
step, not a routed-forward fix). The consumption clone tag bump + a FRESH-demo-up gate-reproduction is the
recommended next step (harden / close-milestone).
**Lessons:** (1) a manager vantage that links the Library INHERITS the employee's huge `/sim` + `/skill-path`
families — manager sample rules must be a SUPERSET (the 2 manager fan-outs + the 2 library families), or the
crawl explodes + times out (no report). (2) An "empty dashboard page" on a fully-seeded org is often an
inserted-but-invisible JOIN gap (the page reads an app-side MIRROR, not the base table) — diagnose the
resolver's join before assuming a missing seed (the org-feedback mirror, the M36-D2 pattern at org scale).
(3) The route-model reconciliation alone turned notReached=5 into 6 fully-asserted dashboard pages — the M36
dashboard was already real; the manifest just pointed at the wrong routes.
