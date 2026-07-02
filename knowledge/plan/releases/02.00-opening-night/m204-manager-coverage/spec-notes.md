# M204 Spec Notes

Technical notes accumulate here during the iter loop. The authoritative design lives in [`overview.md`](overview.md)
+ the consolidated capability spec
[`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) (v0.3). M204 is
**tooling + docs only — zero platform-repo edits**; an un-drivable surface escalates via
`unimplementable-without-platform-edit`, it never edits the platform.

## Scope (see overview.md for the authoritative gate)
The declared **manager-vantage** use cases — Dan's core journeys (declared in the **M201 manifest corpus**):
- **Workforce funnel + member roster** — the mapped→verified funnel + the member list.
- **Member drill-down** — the per-member activity-dashboard.
- **Succession / at-risk** — the Growth tab signals.

Each becomes one Playthrough (one use case ↔ one Playthrough), played as Dan via the M202 foundation.

## Reuse paths (the M202 foundation + the shared e2e foundation)
- The **M202 page-object layer** (the per-surface locator/landmark registry) — extend it with the manager
  surfaces; the registry is the **additive merge surface** shared with M203 (employee). Re-pin is O(surfaces).
- The **M202 dedicated decoupled seed** + the reset-to-seed serial-default runner.
- `stack-demo/rosetta-extensions/stack-verify/e2e/lib/cockpit-login.ts` — Dan's login (the M37 handshake).
- `stack-demo/rosetta-extensions/stack-seeding/` — the seeding machinery the reset rides.

## Per-vantage assertion notes (the manager activity-dashboard)
- The member drill-down (activity-dashboard) reads replayed content via the M40 Directus serve-grant deep-fetch
  closure (the M46 lesson) — assert the **outcome state** (the dashboard hydrated with real per-member content),
  not pre-seeded specifics (P2: assert outcomes, not generated member names / counts that vary across captures).
- Succession/at-risk signals are computed projections — assert their **presence / structure**, not exact values
  that vary with the seed.

## Tag / two-repo state
TODO (iter loop): per-iter rext authoring commits + tags; consumption-clone checkouts; the corpus m204 branch.

## Open questions (resolve in the iter loop; record in decisions.md)
- Which manager surfaces need a landmark anchor vs play on semantic locators (per-iter against the false-fail
  rate) — the additive registry merge with M203 must not collide.
- The member-drill-down outcome assertion against the activity-dashboard's federation deep-fetch (the M46
  serve-grant closure must be present in the demo's replay).
- The succession/at-risk Growth-tab signal's concrete user-observable outcome marker.

## Pre-flight audits — iter-01
**KB-fidelity audit (Phase 0b, iter-01 bootstrap):** verdict **YELLOW** — docs aligned + code-faithful; gaps are
known-context for the strategy, no RED/blind areas. Report: run via general-purpose sub-agent (2026-07-02).
Key findings carried into TOK-01:
- Foundation docs (`playthroughs.md`) match the built rext `playthroughs/` section exactly; the M203 iter-05
  Sentinel-Reload is real in `run-playthroughs.sh` (POST `AuthorizationService/Reload` post-reseed).
- Manager routes documented consistently across M201 corpus / `stories-spec.md` / `coverage-protocol.md`:
  `/enterprise/workforce` (ONE tabbed SPA), `/enterprise/members` (roster read), `/enterprise/activity-dashboard`
  (+ per-member drill-down), succession/at-risk at `/enterprise/workforce/(new)/succession` (a ROUTE, not the SPA
  Talent-Pool tab).
- **Manager hero = Morgan Reyes** (`pt-manager`, org_role=admin, Meridian Labs / Org A). "Dan" in the milestone
  prompt is the demo-showcase precedent (Dan Rossi, Cervato); the pt-world dedicated seed's manager is Morgan
  (test data ≠ demo data, spec §5.4). Manager UCs name `actor.hero: pt-manager`, `actor.entitlement: enterprise`.
- **Perf wall**: pt Org A = size:40, well under the ~200-member wall (M46/M51/M53) — no injection demo-patches
  expected; grids render fast.
- **Seed-scale is the key known-context**: confirm the base stories model runs the M36 org-dashboard seeders at
  Org A (funnel / roster / succession / activity-dashboard). Org A's only end-user hero (Pat) is `thriving` — the
  at-risk signal may need a struggling hero in Org A (Sam is in Org B). Resolve per-iter against real render.
- `workforce-intelligence.ai-readiness-monitoring.UC1` is OUT of M204's declared 3 (showcase-only, Northwind).
