# M203 Spec Notes

Technical notes accumulate here during the iter loop. The authoritative design lives in [`overview.md`](overview.md)
+ the consolidated capability spec
[`knowledge/plan/spec-drafts/playthroughs/spec.md`](../../../spec-drafts/playthroughs/spec.md) (v0.3). M203 is
**tooling + docs only — zero platform-repo edits**; an un-drivable surface escalates via
`unimplementable-without-platform-edit`, it never edits the platform.

## Scope (see overview.md for the authoritative gate)
The declared **manager-vantage** use cases — Dan's core journeys:
- **Workforce funnel + member roster** — the mapped→verified funnel + the member list.
- **Member drill-down** — the per-member activity-dashboard.
- **Succession / at-risk** — the Growth tab signals.

Each becomes one Playthrough (one use case ↔ one Playthrough), played as Dan via the M201 foundation.

## Reuse paths (the M201 foundation + the shared e2e foundation)
- The **M201 page-object layer** (the per-surface locator/landmark registry) — extend it with the manager
  surfaces; the registry is the **additive merge surface** shared with M202 (employee). Re-pin is O(surfaces).
- The **M201 dedicated decoupled seed** + the reset-to-seed serial-default runner.
- `stack-demo/rosetta-extensions/stack-verify/e2e/lib/cockpit-login.ts` — Dan's login (the M37 handshake).
- `stack-demo/rosetta-extensions/stack-seeding/` — the seeding machinery the reset rides.

## Per-vantage assertion notes (the manager activity-dashboard)
- The member drill-down (activity-dashboard) reads replayed content via the M40 Directus serve-grant deep-fetch
  closure (the M46 lesson) — assert the **outcome state** (the dashboard hydrated with real per-member content),
  not pre-seeded specifics (P2: assert outcomes, not generated member names / counts that vary across captures).
- Succession/at-risk signals are computed projections — assert their **presence / structure**, not exact values
  that vary with the seed.

## Tag / two-repo state
TODO (iter loop): per-iter rext authoring commits + tags; consumption-clone checkouts; the corpus m203 branch.

## Open questions (resolve in the iter loop; record in decisions.md)
- Which manager surfaces need a landmark anchor vs play on semantic locators (per-iter against the false-fail
  rate) — the additive registry merge with M202 must not collide.
- The member-drill-down outcome assertion against the activity-dashboard's federation deep-fetch (the M46
  serve-grant closure must be present in the demo's replay).
- The succession/at-risk Growth-tab signal's concrete user-observable outcome marker.
