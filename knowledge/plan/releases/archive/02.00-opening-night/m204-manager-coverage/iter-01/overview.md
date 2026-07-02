---
iter: 01
milestone: M204
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
created: 2026-07-02
---

# M204 iter-01 — bootstrap tok (author the manager-coverage strategy)

**Type:** tok (bootstrap)

## Purpose
The unconditional iter-01 bootstrap tok: author the FIRST manager-vantage coverage strategy (TOK-01) that the
first batch of tiks follows. No prior iters, no stalled strategy to revise — this authors the opening move, from
the M204 `overview.md` (scope + exit gate), the M201 manifest corpus (the declared manager use cases), the
M202/M203 foundation (`corpus/ops/demo/playthroughs.md` + the built rext `playthroughs/` section), and the Phase 0b
KB-fidelity audit (YELLOW — known-context carried below).

## Inputs consumed
- **Gate** (overview.md): every declared manager-vantage use case has a PASSING Playthrough on a COLD
  reset-to-seed demo, 0 false-fails over 5 consecutive reset runs. The declared manager journeys:
  1. **Workforce funnel + member roster** — the mapped→verified funnel + the member list.
  2. **Member drill-down** — the per-member activity-dashboard.
  3. **Succession / at-risk** — the Growth-tab / succession signals.
- **M201 corpus** (`manifest-draft.yaml`) — the prose use-case declarations these map to:
  - Journey 1 → `workforce-intelligence.skills-funnel.UC1` (`/enterprise/workforce` funnel + org-scale gap)
    **+** `workforce-intelligence.roster.UC1` (`/enterprise/members` roster).
  - Journey 2 → `assignment-monitoring.assign-and-track.UC2` (`/enterprise/activity-dashboard` + per-member
    drill-down; the M201 corpus maps the drill-down to this UC, NOT a WI UC).
  - Journey 3 → `workforce-intelligence.talent-pool.UC1` (succession/at-risk/mobility at
    `/enterprise/workforce/(new)/succession` — a ROUTE, not the SPA Talent-Pool tab).
- **Foundation** (M202/M203): the manifest + validator (`ptvalidate`), the per-surface page-object layer
  (`e2e/lib/*-page.ts` extending `PageObject`), the pt-world dedicated seed + `seed-worlds.yaml`, the
  reset-to-seed serial runner (`run-playthroughs.sh N --reset`, with the M203 Sentinel-Reload), the four-state
  `ptreport` map (gate = `NoRegressions()`).
- **Phase 0b audit (YELLOW)**: manager hero = **Morgan Reyes** (`pt-manager`, admin, Meridian Labs / Org A —
  confirmed wired in demo-1's roster export); routes documented consistently; pt Org A size:40 under the perf
  wall; the seed-scale question (does the base stories model render the M36 org-dashboard surfaces at Org A?) is
  the load-bearing per-iter unknown.

## Initial strategy → recorded as TOK-01 in ../decisions.md
See milestone-root `decisions.md` → **TOK-01: manager-surface-per-iter (declare → page-object → play → diagnose
→ re-measure)**. In brief: the 3 manager journeys map to 4 declared UCs; the strategy is one manager surface per
tik, in the M203 measure→triage→fix loop, on the shared page-object layer (additive, non-colliding with M203's
employee surfaces). Order by evidence-of-drivability + the read-only nature of the surfaces (all 4 are monitoring
READ flows — no mutation — so reset-to-seed is a determinism formality, not a per-run necessity, but the 5-run
gate still runs it). The pt-manager (Morgan) actor logs in via the existing `hero-login` seat-switch. Each UC's
observable outcome = the org-scale aggregate surface renders REAL data (P2 — presence/structure, never
placeholder, never exact seeded counts).

## Next-tik direction (iter-02, the first tik)
**iter-02 target:** the Workforce funnel + member roster (journey 1 → `workforce-intelligence.skills-funnel.UC1`
+ `workforce-intelligence.roster.UC1`) — the manager's landing surface (`jump_to: /enterprise/workforce`), the
lowest-risk drivable READ surface + the natural first login proof for the manager vantage. Author `workforce.yaml`
(the manager product manifest), a `workforce-page.ts` page object, and the first manager spec(s); run
`--reset --grep` against demo-1; reconcile with `ptreport`. First confirm the base pt-world seed renders the
funnel + roster at Org A size:40 (the key known-context) before concluding a capability gap.

## Close
See `progress.md`.
