---
iter: 23
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-01
---

# iter-23 — the batchA residual, with every CORRECTION re-derived against platform source

**Active strategy reference:** `TOK-01` ("instrument first, then follow") — step 4, *"then the corpus"*.
Clause 5 is the corpus half of the gate and the only one of the three open clauses that is pure documentation
work; clause 2 (Playthroughs) is the substantial remainder and needs a serving stack.

## Step 0 — re-survey (mandatory, done before targeting)

Ran at open, against platform origin HEAD `2adcf71` (re-fetched; **unchanged** — re-scope trigger stays at
occurrence 1 of 2) and the `app` clone at `5ba17044`.

1. **All 18 handed-forward blockers re-anchored by quote.** Line numbers had shifted in
   `external_services.md` (iter-22 edited it); every quote still resolves.
2. **Every CORRECTION re-derived against platform source before use** — the §5 rule this milestone added at
   iter-22, after two inherited corrections turned out to be false. Results below.

| # | handed correction | re-derivation verdict |
|---|---|---|
| 1–3 | the `local_*` mirror is dropped; resolver reads the canonical entity; write-set collapses to one row | **CONFIRMED** — `app/terraform/migrations/20260729133514.sql:58,62` drops it; `intelligence.go:1700` `m.ent.JobSimulationSession.Query()`; `20260722104506.sql:2` creates `job_simulation_sessions` and `:79` `DROP TABLE "sessions"`; `atlas.hcl:8` pins `search_path=public`; the only `CREATE SCHEMA` in the whole migration set is `auth` |
| 4–5 | one subgraph; no `graphql` compose service | **CONFIRMED** — no `graphql:` service in `platform/docker-compose.yml`; both frontends point at `:8082/graphql/query` (`:318`, `:334`, `:352`, `:361`) |
| 6 | the `validation_*` trio is in `public`, and the middle one is `validation_attempt_skill_results` | **CONFIRMED** — `20260722081626_jobsim_data_model.sql:336/355/376` |
| 7–11 | the Directus env lands on `cms`, not `backend`; seven services in the default profile; frontends target `backend` | **CONFIRMED** — `backend`'s env block (`:43-77`) has no `DIRECTUS_*`; `cms`'s (`:164-165`) has both; profile `graphql` = sentinel · backend · jobsimulation · cms · storage · roadrunner · gotenberg = **7** |
| 12–14 | the ai-readiness tables are in `public` | **CONFIRMED** — `…jobsim_data_model.sql:47/244/261`; `how_we_measure.go:285-287` reads `FROM public.interactions i JOIN public.job_simulation_sessions s` |
| 15 | colony is split `v0.35.2` (app+messenger) / `v0.34.3` (sentinel+storage) | **CORRECTION ITSELF INCOMPLETE.** Six live services carry **three** pins, not two: `app`+`messenger` `v0.35.2`, **`cms`+`jobsimulation` `v0.35.1`** (both are live compose services at origin HEAD — the handed correction omitted them entirely), `sentinel`+`storage` `v0.34.3`. Applying it verbatim would have shipped a second incomplete claim |
| 16 | `app` is on colony `v0.35.2`, not `v0.34.3` | **CONFIRMED** (`app/go.mod`) |
| 17 | `ai` is `v1.40.2` | **CONFIRMED** — and it is `v1.40.2` in `cms` and `jobsimulation` too |
| 18 | studio-room's root is `app/studio/` | **CONFIRMED** — `app/studio/` holds `gen.py`, `requirements.txt`, `agents/`, `services/`; `studioManager.go:119` invokes `studio/gen.py`, venv at `:92-94` |

**Two blockers the handed list did not carry, found by the same re-derivation** (so the target grows, it does
not shrink):

- `hiring.md:142-144` claims rext's `PersonaSeeder` writes **both** `jobsimulation.sessions` and
  `public.local_jobsimulation_sessions` via `sessionCols()`/`localSessionCols()`. rext re-pointed this in M257:
  `persona_write.go:91` writes `{"public", "job_simulation_sessions", sessionCols(), …}` and `localSessionCols`
  **no longer exists**. The corpus describes our own tooling as it was three releases ago.
- `hiring.md:1-15` — the doc's own **headline and framing** ("the score is a MIRROR table in `app`", *"the ONE
  table that actually feeds the score"*) is the false claim, not just the rows below it. Correcting the table
  rows while leaving the thesis intact would leave the doc arguing for the wrong table.

## Cluster / target identified

`DOC-M257x-iter22-batchA-residual` — the 18 enumerated blockers + the 2 found at re-derivation, across 6 files:
`hiring.md` (the dominant one, and this milestone's own class sitting in the corpus — it names a **dropped**
table as the score source, then publishes a write-set whose targets no longer exist), `external_services.md`,
`ai-readiness.md`, `shared_libraries.md`, `clerkenstein.md`, `studio-room.md`.

If budget allows, `DOC-M257x-iter22-ops-guides-5050` (13 dead `:5050` references in `corpus/ops/**`, incl. the
onboarding guides) follows in the same iter as a second planned line.

## Hypothesis

These are false-at-HEAD claims with confirmed replacements; an enumerated exactly-once-anchor sweep closes them
without re-derivation risk, because the re-derivation already happened above.

## Expected lift

Clause 5's residual blocker count falls by ~20. Clause 5 does **not** close this iter — closing it requires the
full 40-file re-read (iter-22's lesson: fixing the handed list measures the hand-off, not the corpus), which is
the next iter's work.

## Phase plan

Two planned lines (declared, so the scope-creep tripwire counts against this shape):

1. Sweep the 20 batchA sites with an exactly-once anchor harness; run both corpus guards.
2. The `:5050` batch in `corpus/ops/**`, same harness.

## Escalation conditions

- A platform commit landing mid-iter → re-scope trigger occurrence 2 → STOP.
- An anchor matching 0 or 2+ times → fails loudly; re-derive rather than force.

## Acceptable close-no-lift outcomes

A correction that re-derivation refutes is a **finding**, not a failure — record the refutation and leave the
corpus line alone (this is what saved iter-22 from shipping two false statements).
