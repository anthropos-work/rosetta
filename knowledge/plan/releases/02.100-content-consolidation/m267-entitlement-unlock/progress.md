# M267 — Progress

**Status: PAUSED AT PHASE 1, section 1 of 6.** Checklist mirrors the `In:` list in
[`overview.md`](overview.md).

## ▶ RESUME HERE

**Paused by the user 2026-08-24** to take priority work on `main`. Re-enter with
`/developer-kit:build-milestone` on branch `m267/entitlement-unlock` (rosetta) — the rext branch of the same
name exists at `v2.9.23-rext` with **no commits**, which is correct: no code work had started.

**What is DONE, and must not be redone:**

- **Phase 0b KB-fidelity: GREEN.** Recorded in [`spec-notes.md`](spec-notes.md) with the three audit-reuse
  conditions; report in [`kb-fidelity-audit.md`](kb-fidelity-audit.md). Check the conditions, then **skip
  straight to Phase 1** if they hold.
- **One doc fix landed** — `corpus/services/sentinel.md` at `:201` and `:93`. The unqualified claim that
  *"'default' org policies apply to all organizations"* is true for `m2`/`m3`/`m5` and **false for `m6`**
  (`casbin.go:45`, zero occurrences). **Read that before writing the insert**: there is no one-row way to
  cover every org, and getting it wrong fails silently behind the PostHog-gated error path.

**What is NOT done:** every checkbox below. No seeder code was written, no test, no stack was touched.

**Where to start:** section 1, and its first real decision is open question 1 in
[`overview.md`](overview.md) — *is `FEATURE_JOB_SIMULATIONS` the only feature `p6` gates on a demo?* That
answer decides whether B1's "remove any limitation" is one row per org or several, so it comes before the
insert, not after.

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
