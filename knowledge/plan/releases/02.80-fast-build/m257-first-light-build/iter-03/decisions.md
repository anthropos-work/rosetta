# iter-03 — intra-iter decisions

## D1 — `public.job_simulation_sessions` is written as the SAME ROW, not a mapped copy

It is the *jobsim-in-app* table (`20260722104506.sql`) and is **column-identical to
`jobsimulation.sessions`**, so the correct re-point writes the same row at the same id — which is what
the migration itself did. Treating it as a renamed mirror would have forced **fabricating** `sim_type`
and `token`, and a fabricated `sim_type` is the **G14 inserted-but-invisible** class: the row lands, no
error is raised, and the reader filters it out. **An error would have been the better failure.** Both
are instead derived with the identical expressions the sibling seeder uses for the same key.

## D2 — the larger `repos.yml` blocker was routed, not solved

iter-03's own escalation condition said *"a seeder re-point that needs a schema decision beyond
re-pointing → route forward rather than guess."* This is that case, one level up: `platform`'s
`repos.yml` now declares `app` the **sole** migration owner, so a fresh stack never creates the
`jobsimulation` schema that ~15 rext writes target.

The two fix shapes — **pin `platform` to a pre-drift ref** (reproducible now, knowingly stale, drifting
further weekly) versus **follow the new model** (correct and durable, but a second B1 of unknown size,
where B1 alone was 34 sites across 20 files) — differ in **release scope**, not in implementation
detail. That is a user decision, and it is the exit reason for this call.

## D3 — two known-imperfect things were left alone deliberately

- **`feedback.go`'s score approximation** (`>=55` band vs the sibling's `passThreshold` nudge) is
  pre-existing and **out of B1's remit** — B1 re-points tables, it does not change seeded values. But
  the risk profile changed: a divergence that was harmless between a mirror and its source is **not**
  harmless between two tables claiming to be the same row. Routed with the exact fix named.
- **`app`'s `.dockerignore` not excluding `studio/.git`** — parity with `app`'s own CI was kept rather
  than diverging from it in tooling. Flagged, because a 2.2 MB directory entering the build context is
  on-topic for a release named *fast build*.
