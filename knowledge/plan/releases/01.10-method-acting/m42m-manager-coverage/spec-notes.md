# M42m — Spec notes

Iteration-protocol-specific technical notes (the manager-vantage coverage sweep, the harness reuse, the
manager-reachable page set, the Workforce-dashboard surfaces, the in-demo/escape boundary). Accumulates across
iters.

## Pre-flight audits — iter-01
`/developer-kit:audit-kb-fidelity` (manual, via sub-agent, 2026-06-25): **YELLOW** — docs accurate where they
make claims (coverage-protocol.md harness description matches the code; all 6 M36 seeders exist + write what
stories-spec claims, 3 spot-checked clean). YELLOW only because M42m's 3 core unknowns are undocumented (the
milestone's designed discovery work, not a fidelity failure). **Two load-bearing facts into iter-01:**
1. **The real manager route is `/enterprise/workforce?tab=skills-verification`** (confirmed in the seeded
   `stories.seed.yaml` cockpit `jump_to` + `test_cockpit.py`), NOT the manifest's guessed `/workforce/*`
   sub-routes. The `MANAGER_PAGES` paths are `calibrated:false` best-guesses → **reconcile the manifest route
   model FIRST** (tab-query, not sub-route). The notReached=5 is very likely a wrong-path manifest guess, not
   purely a nav/seed gap.
2. **`CORS_EXTRA_ORIGINS` on `/api/workforce/*`** is already wired by the demo injection
   (`gen_injected_override.py` per `frontend-tier.md`) — verify before assuming an empty dashboard is a seed
   gap.
Blind areas (expected, = discovery targets): (a) the next-web client page path for the workforce dashboard
(API `/api/workforce/*` documented; client path NOT) — likely `/enterprise/workforce?tab=…`; (b) what gates
the Workforce nav-link visibility (g3-for-sim-start is the closest analog precedent; manager JWT carries
`org_role=admin`); (c) where the baked `studio.anthropos.work` Studio left-nav link is defined + whether it's a
`NEXT_PUBLIC_*_URL` (rewritable) or hardcoded (re-scope trigger if hardcoded).

## Coverage harness + protocol (reused from M42e)
The M42e Playwright harness (`.agentspace/rosetta-extensions/stack-verify/e2e/`) is vantage-generic: the
`MANAGER_MANIFEST` (`coverage-manifest.ts`, seat `dan-manager`) + `run-coverage.sh <N> manager dan-manager`
drive it against the manager hero. The gate composition is identical to employee
(`0 failingSections + 0 personaFailures + 0 escapes + 0 notReached + frontier EXHAUSTED`). A "failing page" =
a manifest section below the content/cardinality bar; an "escape" = an off-demo prod-eject `<a href>` the
crawl's escape-classifier didn't clear via the allow-rule. Persona (role↔skills, avatar menu==profile,
org name+logo) already PASSES for dan-manager (the M42e identity machinery generalizes — iter-23 smoke-sweep).

## Manager-roster hero (demo-3)
Canonical manager hero: **Dan Rossi** (cockpit seat `dan-manager`, the first story's manager — story
`ai-transformation`, Cervato Systems, org `22222222-…`). The sweep logs in as `dan-manager` via the cockpit
handshake. Persona PASSES (the M42e identity machinery generalizes — role↔skills, avatar menu==profile, org
name+logo). The gate is per-hero; Dan is the canonical manager. (Leah Donovan = the second story's manager,
Solvantis — out of the demo-3 first-org default-login scope; Dan is the sweep hero.)

## The real /enterprise/* route model (iter-04 — the reconciliation)
The manifest's `/workforce/*` sub-route guesses were WRONG. Diagnosed live (dan-manager, demo-3): the manager
nav links the **`/enterprise/*`** surface, and the Workforce dashboard is **ONE tabbed SPA route**, not 5
sub-routes:
- **`/enterprise/workforce`** (`/enterprise` redirects here) — the **Workforce-Intelligence dashboard**. The
  M36 sections render as **in-page TABS** (Growth / Skills & Verification / Talent Pool / Assignments / Activity
  Log), NOT `?tab=` query routes (the tab query is ignored client-side). Live render: 221 members, 493 skills
  mapped across 159 roles, 262 verified / **53.1% coverage**, 40 skill paths / 480 sims completed (69.6% pass),
  **19 ant-card + ~67 chart SVGs**. This single page carries the **verification funnel + the org-scale gap**.
- **`/enterprise/members`** — the org member table (221 members: name / email / role / business-unit tags).
- **`/enterprise/assignments`** (+ `/ai-simulations` `/ai-interviews` `/skill-paths` sub-tabs) — the
  assign-AI-sims roster; **fans out to `/user/<uuid>`** (the team-roster member profiles).
- **`/enterprise/activity-dashboard`** (+ `/ai-simulations` `/skill-paths` `/interviews` sub-tabs) — the
  per-sim activity table (success / avg-score / attempts / time / dates); **fans out to
  `/enterprise/activity-dashboard/.../<uuid>`** drill-downs (the leaf is dashboard chrome, textLen ~70).
- **`/enterprise/organization-feedback`** — the ~2:1 feedback distribution (see the fix below).
- **`/enterprise/settings`** — org branding / AI-sim toggles / usage — **terse by design (a documented
  exception**, mirroring the employee `/settings`).

## M36 Workforce-Intelligence dashboard surfaces — coverage status
All M36 surfaces render REAL, substantial seeded data on demo-3 (the M36 seeders already populate them; the
notReached=5 was purely the route-model error, not a content gap):
- **mapped→verified funnel** — 493 mapped / 262 verified / 53.1% coverage (on `/enterprise/workforce`). ✓
- **org-scale claimed-vs-verified gap** — the dashboard stat cards + tabs. ✓
- **teams / role gap+mobility / succession** — the in-page tabs of `/enterprise/workforce` (Growth / Talent
  Pool / etc.) + `/enterprise/members` tags. ✓
- **feedback** — `/enterprise/organization-feedback` — was the ONE empty surface; fixed (below). ✓

## The two manager FAN-OUTS (the frontier-exhaustion preconditions)
1. **`/user/<uuid>`** (+`/skills`+`/activities`) — the team-roster member profiles (200+ members; each a real
   profile: name / role / location / bio / skills / completed sims). Template-identical.
2. **`/enterprise/activity-dashboard/{ai-simulations,skill-paths,interviews}/<uuid>`** (+ nested `/<uuid>`) —
   the per-activity drill-downs (hundreds; the leaf renders only the dashboard tab chrome).
Both handled by manager **SAMPLE_RULES** (`coverage.spec.ts` `MANAGER_SAMPLE_RULES`, vantage-aware, sample 8
each) so the frontier where escapes/failures live EXHAUSTS (`cappedAtFrontier=false`) while bounding the
explosion — the M42e `/sim/<slug>` sample-rule pattern, applied to the manager families.

## Admin pages (manager vantage)
All manager-only pages are the `/enterprise/*` set above — each in-scope for POPULATE (asserted by the manager
manifest), except `/enterprise/settings` (a documented terse exception, in-demo but not cardinality-floored).

## rosetta-extensions fixes (manager vantage)
- **`stack-seeding/seeders/feedback.go` (the org-feedback mirror fix, iter-04).** `/enterprise/organization-feedback`
  rendered "No data" / 0 sessions despite 103 seeded `job_simulation_feedbacks` rows. ROOT CAUSE (read-only
  platform diagnosis — `app/internal/data/ent/repository/jobsimulation.go GetOrganizationFeedback`): the page
  JOINs feedback to the app mirror `public.local_jobsimulation_sessions` on
  `feedback.session_id = mirror.jobsimulation_session_id` and scopes by the **MIRROR's** `organization_id`. The
  population sessions the feedback links live ONLY in `jobsimulation.sessions` (the `JobsimSessionsSeeder` writes
  NO mirror — only the `PersonaSeeder` mirrors, for heroes), so the join found nothing. FIX: the `FeedbackSeeder`
  now ALSO writes a `local_jobsimulation_sessions` mirror per feedback session (reconstructing the population
  session's coherent values — score / start-end / status — from the SAME deterministic key). Reproducible,
  idempotent, `PerStackIsolated`. Live result: feedback rows joinable 0→103, page "No data" → "103 sessions /
  70% positive / 59% pass / 63% avg" + 21 real rows. Zero platform edits.
- **`stack-verify/e2e/lib/coverage-manifest.ts` (MANAGER route reconcile + calibrate, iter-04).** `MANAGER_PAGES`
  re-authored from the wrong `/workforce/*` guesses to the real `/enterprise/*` routes, with per-section asserts
  calibrated against the live render; `calibrated:true`. `MANAGER_MANIFEST.seedPaths` updated.
- **`stack-verify/e2e/tests/coverage.spec.ts` (manager sample rules, iter-04).** Added `MANAGER_SAMPLE_RULES`
  (the two fan-outs) + vantage-aware selection so the manager frontier exhausts.
