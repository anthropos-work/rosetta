---
milestone: M267
title: "The entitlement unlock"
milestone_shape: section
status: planned
release: "02.100-content-consolidation"
depends_on: "none"
parallel_with: "M266, M270"
complexity: small
last_updated: "2026-08-23"
---

# M267: The entitlement unlock

**Goal:** A seeded hero can actually START a simulation — no plan ceiling, no silent entitlement gate.

Serves annotation request **B1** ("please remove any limitation to access and start content in the demo
env") and the **"cannot start" half of B2**. **The fix is ONE seeder INSERT.**

## Scope

**In:**

  - **Seed one `p6` row per seeded org.** The whole milestone. Shape and helper already exist — reuse
    `casbinGrantSQL()` (`identity.go:227-238`) and write, per seeded org,
    `INSERT p6 (<org_eid>, FEATURE_JOB_SIMULATIONS, 1000000)` — the row the DEV path already writes at
    `dev-stack/dev-identity.sh:205-208`. **The measured chain this closes, carried verbatim:**
      - The UI string is `errors.simulationLimitReached`, rendered by
        `next-web-app/packages/ui/src/AISimulation/AISimulationStartWithoutSession.tsx:209` on
        `ERROR_JOB_SIMULATION_LIMIT_REACHED`.
      - That error exists in ONE place:
        `app/internal/jobsimulation/simulator/manager/errors.go:7`, returned at `manager.go:386` when
        `CanPerformFeatureAction(FEATURE_JOB_SIMULATIONS)` is false.
      - → `user/directory.go:121` → `subscriptions.go:398` → `sentinel/manager.go:65`, a Casbin enforce
        on **matcher 6**:
        ```
        casbin.go:29   p6 = org, feat, max
        casbin.go:45   m6 = ( g3(p6.org, r6.sub) ) && r6.feat == p6.feat && r6.count <= parseFloat(p6.max)
        ```
      - ⚠️ **UNLIKE m2/m3/m5, m6 HAS NO `default` ESCAPE — the `p6` row must name the org id.**
      - A demo has the `g3` row (`stack-seeding/seeders/identity.go:253-283`
        `seedCasbinFeatureGrants`) and **NO `p6` row**. Grepping all of `rosetta-extensions` for `p6`
        returns only `dev-stack/dev-identity.sh:205-208` — the **DEV** path. The platform's own
        `sentinel/init_policy.sql` seeds **ZERO** `p6` rows.
      - In real operation the row is written by the Clerk licensing path
        (`app/internal/organization/licensing.go:171` → `SetOrganizationFeatureCredits`, default
        **1,000,000** at `manager.go:1566`) — which **NEVER RUNS on a demo**, because the seeder writes
        orgs straight into Postgres.

  - **Confirm the safety argument holds after the insert.** `r6.count` is `Used` from
    `public.organization_features` (`ent/repository/organizations.go:1652`); the seeder writes no rows
    there (`identity.go:125` says so), so **`Used` = 0 regardless of backdated sessions**. Verify no
    *other* seeder in the release (M268's seeded truth, the programs seeding) starts writing that table
    — see Open questions.

  - **No new plumbing — confirm the existing reload path carries the row.** Casbin caches in memory, but
    the reload is already wired: `demo-stack/up-injected.sh:2749` publishes `sentinel:policy:invalidate`
    over Redis post-seed, and `playthroughs/e2e/run-playthroughs.sh:184` does the same. **Observe it,
    do not build it.**

  - **Prove it live**: a seeded hero (employee vantage and manager vantage) starts a simulation on a
    cold reset-to-seed stack and gets a session, not `errors.simulationLimitReached`. Browser-verified —
    a curl hits an access gate, and a bundle-grep proves presence, not that the gate opened.

  - **RECORD the three limiters that are already inert, so they are not chased:**
      - `MaxConcurrentSessions = 200` (`manager.go:328`) — org-wide **LIVE** sessions, **not a plan**.
      - `OrganizationSetting.SimulationsMaxAttempts` — applied only when `isHiring`;
        `org_settings.go` does not seed it.
      - Org-**LESS** users would hit matcher `m` with `TIER_FREE = 2` (`init_policy.sql:23`) — **no demo
        hero is org-less.**
      - ⚠️ **DIAGNOSTIC TRAP:** on any OTHER failure the demo shows a **generic** message, because the
        raw-error path is gated on `posthog.isFeatureEnabled('anthropos_internal_team')`
        (`AISimulationStartWithoutSession.tsx:214-217`) and **a demo has no PostHog**. Expect
        `errors.simulationCannotStartGeneric` to hide the real cause.

  - **Deliver the doc**: the `p6` grant becomes part of the documented seed contract in
    [`corpus/ops/seeding-spec.md`](../../../../../corpus/ops/seeding-spec.md), together with the
    `Used = 0` safety argument and the invalidate-publish dependency.

**Out:**

  - **The Stripe tier gates themselves.** The M205-residual reservation is about **PROVING they work**;
    this milestone **DISARMS them in demo** — same surface, opposite intent, and **this milestone
    consumes neither**.
  - **The 4-modality Playthrough coverage that is the other half of B2** (chat / voice / code /
    document). The sibling milestone `m269-modality-playthroughs` is named for it. *Inferred from the
    release layout, not stated in this milestone's brief — see Open questions.*

## Depends on

**none.** Starts cold.

## Parallel with

**M266, M270.**

## Open questions

  - **Is `FEATURE_JOB_SIMULATIONS` the only feature the `p6` matcher gates on a demo?** The measured
    chain covers the start-simulation path only. Whether any other content surface (skill paths, AI
    Labs, academy, the M268 programs) routes through `CanPerformFeatureAction` with a **different**
    feature constant is **not measured**. If it does, "remove any limitation" (B1) is wider than one row
    per org.
  - **Does the fix reach the Playthrough world?** `pt-world` is a separate, decoupled seed. Whether the
    `p6` insert lands on its three private orgs through the same seeder path, or needs a second insert
    site, is not established here.
  - **Do the DEV and DEMO paths converge or stay two sites?** `dev-identity.sh:205-208` already writes
    the row as inline SQL; the seeder would write it through `casbinGrantSQL()`. One shape, two
    implementations — worth a decision (`D-M267-N`), not an accident.
  - **Does any other v2.10 seeder write `public.organization_features`?** The `Used = 0` safety argument
    rests on `identity.go:125`. It is a statement about the **identity seeder**, not about the seeding
    module as a whole, and M268 has not been written yet.
  - **How do we read the real error during verification?** The raw-error path is PostHog-gated
    (`AISimulationStartWithoutSession.tsx:214-217`) and a demo has no PostHog. If a residual failure
    survives the `p6` row, the generic message will hide it and diagnosis may have to come from
    `backend` logs rather than the browser — **or**, if the surface genuinely needs changing, from a
    sha-pinned demopatch (see below). Unresolved.

## KB dependencies

- [`corpus/ops/seeding-spec.md`](../../../../../corpus/ops/seeding-spec.md) — the seed contract this
  milestone extends
- [`corpus/services/sentinel.md`](../../../../../corpus/services/sentinel.md) — the Casbin PDP, the
  matcher set, and the in-memory-cache/invalidate model
- [`corpus/ops/safety.md`](../../../../../corpus/ops/safety.md) — a demo is an authz-**weakened** build
  by design (§2.3, §3); this milestone weakens one more gate deliberately and must say so there

## Delivers →

[`corpus/ops/seeding-spec.md`](../../../../../corpus/ops/seeding-spec.md) — **the `p6` grant becomes
part of the documented seed contract** (the row, the `Used = 0` argument, and the
`sentinel:policy:invalidate` dependency).

## Demo-patch?

**No demopatch needed.** This is a **rext-only DB write** — a seeder insert into `sentinel.casbin_rules`
on a stack's own database. No platform file is read, patched or reverted.

If verification later shows the fix needs a change to platform **source** (the most likely candidate is
the PostHog-gated raw-error path above), that change goes through the sha-pinned demopatch mechanism —
[`corpus/ops/demo/demopatch-spec.md`](../../../../../corpus/ops/demo/demopatch-spec.md) — and **never**
as an edit to an `anthropos-work` repo.

## Constraints (release-wide, non-negotiable)

- **Zero platform-repo edits.** A need that can only be met by a platform edit goes through demopatch or
  **escalates**; it does not edit.
- All stack tooling lives in `rosetta-extensions`, built + tested in the authoring copy, tagged,
  **pushed to origin**, then consumed per-stack at a pinned tag. *Tagging is not publishing.*
- Secrets handled values-blind.
