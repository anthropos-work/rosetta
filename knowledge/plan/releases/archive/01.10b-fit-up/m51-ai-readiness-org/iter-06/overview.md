---
iter: iter-06
milestone: M51
iteration_type: tik
status: closed-no-lift
created: 2026-06-30
---

# iter-06 — re-sweep with the corrected harness/manifest in effect; drive the residual workforce perf-wall grids

## Type
tik — under TOK-01 (active-cycle signals-true). The coverage-drive strand. iter-05 landed the
content-presence `warmHeavyGrids` + the org-agnostic manifest correction + the AI-readiness manifest
assertion, but its 16:47 GATED report was generated against the OLD harness/manifest (it still shows the
`cervato-systems.com` false-fails) — so the true residual under the corrected harness is UNMEASURED. iter-06's
job is to re-sweep with the committed iter-05 harness in effect and read the true `(failing, escapes)`, then
drive the residual.

## Step 0 — Re-survey before targeting
Re-confirmed against the iter-05 close + the live state:
- demo-1 is UP at fit-up-m50 (all 3 perf demo-patches baked) + the AI-readiness showcase org seeded (Northwind
  200, ENABLED, 78.4% all-3, heroes pinned). No re-up / re-seed needed — the iter-05 changes are HARNESS-ONLY
  (warmHeavyGrids + manifest), which per coverage-protocol.md ("The demo must be live + at the consumed tag",
  harness-only clause) run from the AUTHORING copy directly against the live offset ports.
- The iter-05 16:47 report's 5 failures: 2 are the `cervato-systems.com` false-fails (CORRECTED by the
  committed org-agnostic manifest fix — not yet in effect in that report); 3 are the genuine org-scale
  cold-grid perf-wall (verification-funnel, talent-languages, activity-table).
Target (the corrected-harness re-sweep + residual drive) is current + meaningful. No substitution.

## Active strategy reference
**TOK-01** (`../decisions.md`) — active-cycle signals-true. This iter is the coverage-drive strand (TOK-01
step 4). The fix surface for the residual is the HARNESS (the content-presence warm + per-section heavy-grid
poll) and — if a genuinely-slow-but-correct heavy section persists — the coverage-protocol "Org-scale grid
perf wall" disclosed-presenter-note / render-budget rule, per the protocol's disclosed-allow.

## Cluster / target identified
ONE re-sweep establishes the true residual; then the highest-leverage cluster is whichever of the 3
perf-wall grids still false-fails AFTER the content-presence warm primes the cold query. They share the
backing-query family, so a single further lever (a widened per-section heavy-grid re-assert poll for the
org-scale grids, OR a disclosed-presenter-note for a genuinely-slow-but-correct section) is expected to clear
the cluster. Plus a REAL FIX: bake the Sentinel-reload-after-seed step into the seeding flow so the casbin
grants take effect without a manual reload (the stale-Sentinel-policy root cause cleared in iter-05).

## Hypothesis
With the committed `warmHeavyGrids` priming the cold ~11.6s members query during warm + the org-agnostic
manifest correction in effect: the 2 cervato false-fails vanish AND the content-presence warm hydrates the
shared backing-query family so the 3 residual perf-wall grids read real rows at the authoritative visit →
`(failing → 0, escapes = 0)`. If a residual grid is genuinely slow-but-correct beyond the warm ceiling, a
widened per-section heavy-grid poll (or a disclosed-presenter-note per the protocol) closes it without
loosening the gate.

## Expected lift
failingSections 5 → 0 (the 2 cervato by the manifest correction + the 3 perf-wall by the content-presence
warm), escapes 0, persona GREEN → gate-met. Conservative floor: 5 → ≤2 (the cervato pair + at least one
warm-cleared grid), with a characterized residual + a named next lever.

## Phase plan
- Phase A (re-sweep): GATED manager-vantage sweep on demo-1 with the committed iter-05 harness in effect →
  the true residual `(failing, escapes)`.
- Phase B (triage): for each residual failing section — screenshot + DOM/network/latency probe per the
  protocol's empty-page diagnostic; classify perf-wall-cold vs genuine-content vs harness-budget.
- Phase C (fix): the routed real fix — bake Sentinel-reload-after-seed into the seeding flow; AND, for any
  residual perf-wall grid, the protocol-sanctioned lever (widened per-section heavy-grid poll, OR a
  disclosed-presenter-note for a slow-but-correct heavy section). Authored in the authoring copy.
- Phase D (re-sweep): GATED manager-vantage sweep → record `(failing, escapes)` delta.
- Phase E (close): grade on the residual delta; route any remaining residual forward with a named handler.

## Escalation conditions
- If a residual grid is genuinely empty (not slow) → re-triage the fix surface (seed vs snapshot serve-grant)
  per the protocol routing table; if a seed gap, that's a stack-seeding fix not a harness fix.
- If the harness change regresses the employee-vantage sweep (the warm is shared — but warmHeavyGrids is
  manager-vantage-gated, so this is unlikely) → user-blocker.
- If clearing the cluster needs a platform edit → re-scope-trigger (`unimplementable-without-platform-edit`).

## Acceptable close-no-lift outcomes
If the re-sweep shows the corrected harness cleared the cervato pair but a residual perf gap keeps `(failing)`
above 0 with a documented characterization of exactly which grids resist warming + the named next lever →
closed-fixed-partial. A complete characterization that sets up the next iter's lever is a valid outcome.
