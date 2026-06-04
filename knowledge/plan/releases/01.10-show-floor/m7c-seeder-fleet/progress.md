# M7c — progress (iterative running ledger)

**Milestone:** M7c — The seeder fleet, to a coverage gate · **Shape:** iterative · **Status:** planned

**Exit gate:** a `stack.seed.yaml` (1k-user org, N months activity) → demo identity logs in **200** · data-DNA
**coverage ≥ 90% / critical 100%** · seeding **< 2 min** · seeding audit **zero shared/prod writes**.

**Iteration protocol:** `corpus/architecture/alignment_testing.md` (measure → implement/deepen a seeder →
re-measure data-DNA coverage), via M7b's data dimension. Built by `/developer-kit:build-mstone-iters`.

## Running ledger

| Iter | Type | Surface(s) | Coverage | Notes |
|------|------|-----------|----------|-------|
| 01 | tok-bootstrap | — | 40% (4/10, baseline) | TOK-01 strategy; reachability survey; corrected 3 guessed catalog table names. Believability core (activity surfaces) reachable w/o Directus. |
| 02 | tik | jobsim-sessions | **50% (5/10)** | The believability core: `JobsimSessionsSeeder` — 1980 backdated, pass/fail-weighted sessions (deterministic, time-distributed over the activity span). Promoted planned→seeded; `measure` [PASS] 100%/Critical 100% live on demo-1; isolation clean. Caught + fixed a harness bug: `introspect` couldn't load a freshly-promoted surface (Validate-before-populate). seeders 14 tests. |

## Routes carried forward
- **taxonomy** + **content** are snapshot-blocked (skiller snapshot / shared Directus) → tail iter or waiver per the Re-scope trigger (confirm at close).
