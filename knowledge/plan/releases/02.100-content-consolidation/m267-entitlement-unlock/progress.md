# M267 — Progress

**Status: PLANNED.** Not started. Checklist mirrors the `In:` list in [`overview.md`](overview.md).

## 1. Seed one `p6` row per seeded org

- [ ] Reuse `casbinGrantSQL()` (`identity.go:227-238`); insert
      `p6 (<org_eid>, FEATURE_JOB_SIMULATIONS, 1000000)` once per seeded org
- [ ] Idempotent (the `WHERE NOT EXISTS` guard the helper already carries)
- [ ] Covers every seeded org, not just the hero orgs

_Not started._

## 2. Confirm the `Used = 0` safety argument holds

- [ ] `r6.count` is `Used` from `public.organization_features`
      (`ent/repository/organizations.go:1652`)
- [ ] No seeder writes that table (`identity.go:125`) — re-checked against the rest of v2.10
- [ ] Backdated sessions do not move `Used`

_Not started._

## 3. No new plumbing — confirm the reload path carries the row

- [ ] `demo-stack/up-injected.sh:2749` publishes `sentinel:policy:invalidate` post-seed — observed
- [ ] `playthroughs/e2e/run-playthroughs.sh:184` does the same — observed
- [ ] Nothing new built

_Not started._

## 4. Prove it live

- [ ] Employee vantage starts a simulation on a cold reset-to-seed stack
- [ ] Manager vantage starts a simulation
- [ ] Browser-verified, not curl/bundle-grep
- [ ] `errors.simulationLimitReached` is gone

_Not started._

## 5. Record the inert limiters + the diagnostic trap

- [ ] `MaxConcurrentSessions = 200` (`manager.go:328`) — live sessions, not a plan
- [ ] `OrganizationSetting.SimulationsMaxAttempts` — `isHiring` only; `org_settings.go` does not seed it
- [ ] `TIER_FREE = 2` (`init_policy.sql:23`) — org-less only; no demo hero is org-less
- [ ] The PostHog trap (`AISimulationStartWithoutSession.tsx:214-217`) written down

_Not started._

## 6. Deliver the doc

- [ ] `corpus/ops/seeding-spec.md` — the `p6` grant as part of the seed contract
- [ ] The `Used = 0` argument recorded there
- [ ] The `sentinel:policy:invalidate` dependency recorded there

_Not started._

## Findings

_None yet._
