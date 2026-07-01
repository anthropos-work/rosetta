---
iter: iter-02
milestone: M51
iteration_type: tik
status: planned
created: 2026-06-30
---

# iter-02 — tik (strand 1: the 3rd AI-Readiness story + the org-enablement gate writer)

## Type
tik — first tik of M51, runs under the active strategy **TOK-01** (active-cycle signals-true additive-to-stories
seed). Per coverage-protocol.md Phase A–E; this iter lands the first seeder slice + takes the **baseline sweep**.

## Step 0 — Re-survey before targeting
Re-ran the baseline check: the 3rd org does NOT exist yet (the stories preset has 2 stories; `organization_settings`
on demo-1 has 0 rows — nothing writes it). TOK-01's `Next-tik direction` names exactly this target (strand 1) and it
is still untouched + meaningful. No substitution.

## Active strategy reference
**TOK-01** (`../decisions.md`) — active-cycle, signals-true, additive-to-stories. Strand 1 of its 4-strand plan.

## Cluster / target identified
The 3rd-org IDENTITY + ENABLEMENT gate is the prerequisite for everything else (config, funnel, cockpit, sweep):
without the org row + the `organization_settings` `ai_readiness` gate, the dashboard route returns
`aiReadinessEnabled=false` and there is nothing for the manager-vantage sweep to log into. TOK-01 strand 1.

## Hypothesis
Appending a 3rd `stories[]` entry (org "AI Readiness", size 200, hero trio) to `stories.seed.yaml` + adding a net-new
`OrgSettingsSeeder` (one `organization_settings` row per AI-readiness story, `setting='ai_readiness', is_enabled=true`)
+ re-seeding demo-1 yields: a 3rd org with a 200-member population + the manager hero seat + the enablement gate ON.
Then the baseline manager-vantage sweep establishes the starting `(failing, escapes)` for the 3rd org.

## Expected lift
This is the BUILD-from-zero iter; the "lift" is going from "no 3rd org / no sweep possible" to "3rd org renders +
the first authoritative baseline sweep number exists." The dashboard will likely render ENABLED-but-funnel-less
(config + funnel are strands 2/3) — that's the expected baseline, not a failure.

## Phase plan
- Phase A (sweep / measure): N/A pre-build — the org doesn't exist. After the fix, the baseline sweep IS the measure.
- Phase B (triage): n/a (build iter; no failures to triage yet).
- Phase C (fix): (1) append the 3rd story to `presets/stories.seed.yaml`; (2) author `seeders/org_settings.go`
  (`OrgSettingsSeeder`, surface `org-settings`, depends-on `org`, PerStackIsolated); (3) register it; (4) unit test;
  (5) `go build ./... && go vet && gofmt && go test ./stack-seeding/...`; (6) re-seed demo-1 in place against the
  authoring-copy tooling (TOK-01 in-place re-seed model).
- Phase D (re-sweep / re-measure): run the manager-vantage coverage sweep on demo-1 as the new 3rd-org manager hero
  → baseline `(failing, escapes)`.
- Phase E (close): grade `closed-fixed` if the 3rd org renders + gate ON + a baseline sweep number exists; route the
  config/funnel/cockpit strands forward to iter-03+.

## Escalation conditions
- If enabling the dashboard needs a platform edit → re-scope-trigger (escalate, never edit platform). (Not expected —
  the gate is a pure `organization_settings` INSERT; the contract confirms it.)
- If re-seeding demo-1 from the authoring copy can't reach the live DB → user-blocker.

## Acceptable close-no-lift outcomes
If the 3rd org seeds but the baseline sweep can't run (e.g. the manager hero's seat doesn't resolve in Clerkenstein
without a roster re-export beyond this iter's scope), close-fixed-partial: the org + gate landed; the baseline sweep
routes to iter-03. (Falsification: the seat-wiring is a bigger surface than strand-1 scope.)
