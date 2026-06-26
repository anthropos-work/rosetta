**Type:** tok (bootstrap)

# iter-01 — bootstrap tok: author TOK-01 (the manager-coverage strategy)

## Phase plan (bootstrap tok)
Author the FIRST M42m strategy from the milestone overview + coverage-protocol + stories-spec M36 + the M42e
iter-23 manager smoke-sweep + the M42e design-plan, grounded by a fresh live baseline reading. No fix lands
this iter.

## Work done
- Read the gate, the protocol, stories-spec M36 (the Workforce dashboard surfaces + their 6 seeders), the
  M42e design-plan, and the M42e harness code (coverage-manifest.ts / crawl.ts / coverage.spec.ts).
- Ran the **Phase 0b KB-fidelity audit** (sub-agent): **YELLOW** — docs accurate; M42m's 3 unknowns
  undocumented (= the discovery target). Two load-bearing facts confirmed live:
  1. the real manager route is `/enterprise/workforce?tab=skills-verification` (NOT the manifest's
     `/workforce/*` guesses), confirmed in `stories.seed.yaml` + `test_cockpit.py`;
  2. `CORS_EXTRA_ORIGINS` for `/api/workforce/*` is already wired by the demo injection.
- Ran a live **baseline manager sweep** (demo-3, dan-manager, no-gate, cap 250). It crawled 165 pages then
  timed out at the 25-min budget before writing the report — because the manager frontier has **TWO**
  exploding template-identical fan-outs (`/user/<id>` team-roster AND `/enterprise/activity-dashboard/.../<uuid>`
  per-activity drill-downs). The streamed `[crawl]` log gave the load-bearing facts: the `/enterprise/` route
  prefix, the two fan-outs, and `eject=1` on nearly every page (the universal Studio left-nav link).
- Authored **TOK-01** (milestone-root decisions.md): 4 leverage-ordered fix lines — Studio-link escape →
  manifest route reconciliation + dashboard populate → sample rules + cap raise → manifest calibration.

## Close — 2026-06-25

**Outcome:** TOK-01 authored — the manager-coverage strategy (reconcile-route + clear-escape +
populate-dashboard + exhaust-frontier). Baseline grounded: the iter-23 smoke-sweep gate numbers (escapes=139,
notReached=5, frontier CAPPED) stand; the live baseline confirmed the `/enterprise/` route prefix + a SECOND
fan-out + the universal Studio eject. Sample rules are a precondition (the cap-250 sweep timed out without them).
**Type:** tok (bootstrap)
**Status:** closed-fixed (the bootstrap deliverable — TOK-01 + the baseline reading — landed)
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap toks NEVER exit) — (3) re-scope: n —
(4) user-blocker: n — (5) cap-reached: n (toks don't count) — (6) protocol-stop: n — Outcome: continue (loop
into iter-02, a tik under TOK-01)
**Decisions:** D1 (bootstrap-tok baseline finding) — see iter-01/decisions.md.
**Side-deliverables:** spec-notes.md §"Pre-flight audits — iter-01" (the audit verdict + the 2 load-bearing
facts) + §"Coverage harness + protocol".
**Routes carried forward (→ iter-02+, under TOK-01):**
- iter-02: the Studio-link escape (line 1) — diagnose env-configurable vs hardcoded; rewrite or re-scope.
- iter-03+: manifest route reconciliation to `/enterprise/workforce` + dashboard populate (line 2).
- iter-03+: sample rules for the two fan-outs + cap raise (line 3) — a precondition for the authoritative sweep.
- iter-N: manifest calibration (line 4).
**Lessons:** the manager vantage's frontier is materially larger than employee's — TWO template-identical
fan-outs, not one — so the authoritative gate sweep CANNOT run without the sample rules first (the cap-250
no-rule sweep times out). A manifest's uncalibrated route guesses (`/workforce/*`) that match nothing show up
as `notReached`, not as a section FAIL — so a notReached count can be a route-model error, not a seed gap;
reconcile the route model before assuming a content fix.
