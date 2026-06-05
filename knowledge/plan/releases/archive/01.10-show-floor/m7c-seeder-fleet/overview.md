---
milestone: M7c
slug: seeder-fleet
version: v1.1 "show floor"
milestone_shape: iterative
status: done
created: 2026-06-04
completed: 2026-06-05
close_kind: gate-met-over-reachable + waiver (taxonomy + content, Re-scope trigger, user-confirmed)
last_updated: 2026-06-05
exit_gate: "A `stack.seed.yaml` for a 1k-user org with N months of activity yields a stack where (a) the demo identity logs in → 200, (b) data-DNA coverage ≥ 90% across the catalogued surfaces (critical 100%), (c) seeding completes < 2 min wall-clock, (d) the seeding audit log shows ZERO shared/prod writes."
iteration_protocol_ref: corpus/architecture/alignment_testing.md
delivers: rosetta-extensions/stack-seeding/seeders/ (the fleet + backdated activity generator) + the seed presets
---

# M7c — The seeder fleet, to a coverage gate

## Goal
Implement the **fleet of modular seeders** — one per surface in M7b's data-DNA catalog — each measured against
its conformance gene, until a `stack.seed.yaml` produces a **believable, fully-populated, provably-safe** demo
world: a 1k-user org with months of backdated activity, the demo identity logging in to a populated dashboard,
**no prod pollution**, in **under two minutes**.

## Why iterative (not section)
The seeders' *acceptance* is fixed (the data-DNA gene per surface — M7b), but **which** surfaces carry the
believable-demo weight, **how** scenarios compose, and **where** schema reality fights the plan (backdating
ent-Immutable timestamps, FAILED-session chains, bulk-COPY vs side-effecting primitives) only emerge from
building seeder-by-seeder and re-measuring. A fixed up-front checklist of *how each seeder works* would be
speculative. The **coverage + believability + speed + safety gate** is the commitment; the path to it is open.
This mirrors M1 (drive a measured score to a gate) — here the score is **data-DNA coverage** instead of Clerk
alignment.

## Exit gate (the commitment)
A `stack.seed.yaml` for "1k-user org, N months activity" yields a stack where:
- **(a)** the real `user_clerkenstein` identity logs in via browser → **200** on authorized routes (lands in a
  populated org, not an empty shell);
- **(b)** **data-DNA coverage ≥ 90%** across the catalogued surfaces, **critical surfaces 100%** (every seeded
  row is schema-conformant — FK-valid, constraints satisfied);
- **(c)** seeding completes in **< 2 min** wall-clock at 1k scale;
- **(d)** the **seeding audit log shows ZERO `shared`/`external` writes** — the isolation guard held end-to-end.

## Iteration protocol
`corpus/architecture/alignment_testing.md` — the M0 measure → fix-the-diverging-genes → re-measure loop, applied
to **data** via M7b's data-DNA: each iter implements/deepens a seeder, runs the conformance + coverage measure,
and drives coverage toward the gate. (M7b delivers the data-specific extension this loop measures against.)

## Scope (the surfaces — from M7b's catalog; the fleet, not a fixed how)
- **Users** (bulk `COPY`, deterministic identities) · **Orgs / Memberships / Casbin** (the plural/singular gotcha
  honored from M7a) · **Features / Tiers** (Sentinel grants) · **Taxonomy** (consume the pre-embedded skiller
  snapshot) · **Content / Library** (snapshot-replay, never live Directus writes — M7a guard) · **SkillPath
  sessions** · **JobSim sessions + results** (passed AND failed chains) · **Assignments** · **the backdated
  activity generator** (time-distributed real-sim sessions across the span, `seed-verified-skill` pattern).
- **Scenario composition**: the curator annotates `stack.seed.yaml` per-seeder blocks to reach a target scenario
  (e.g. "acme, 1k users, alice→sim-X, months 1–3 heavy then taper").
- **2–3 seed presets** (small/medium/large, e.g. 200/500/1k) shipped as ready blueprints. *(If presets prove to
  belong with the M8 recipe/discoverability work, route there — confirm at build.)*

### Out
- The framework / DAG / safety guard (M7a) and the data-DNA harness (M7b) — M7c consumes both.
- AI-generated content (transcripts / AI-scored narratives / fresh embeddings) — the hard line, v1.2 stretch.
- Authoring taxonomy/jobsim content — consumed from the pre-embedded snapshot.
- Recipe corpus + discoverability skills — M8 (M7c delivers the seeders; M8 curates the use-cases).

## Depends on
**M7a** (the framework + the isolation guard + the perf path + the `user_clerkenstein` proof) + **M7b** (the
data-DNA catalog that lists the seeders + the coverage gate). **Parallel with:** none (the last build milestone
before M8 curates its output).

## Estimated complexity
**large** (multi-iter) — the heaviest build surface in v1.1: ~8–10 seeders, 1k-scale bulk performance, backdating
fidelity, and scenario composition, each gated on conformance.

## Re-scope trigger
If consecutive strategy iters (toks) can't make a surface conformance-pass (e.g. a store whose schema can't be
seeded structurally without a platform change, or backdating a hard-Immutable timestamp), **waive that gene with
justification** in the divergence report (lower the believable-demo bar consciously) or escalate to the user —
don't chase an unreachable 100%. The believable-demo subset of surfaces is the real target, not every entity type.

## Open questions (resolve during iters)
- The minimum surface subset that makes a demo "believable" (the must-cover genes vs nice-to-have).
- Backdating: per-surface, which timestamps direct-SQL-settable vs ent-Immutable/DB-default `now()`.
- Whether presets ship here or in M8.

## KB dependencies (read as contract)
- `corpus/architecture/alignment_testing.md` (the iteration protocol + M7b's data dimension)
- M7a's `corpus/ops/seeding-spec.md` (the framework + isolation guard + blueprint) + M7b's `dna/` (the catalog + gate)
- `corpus/services/{backend,skiller,jobsimulation,skillpath,cms}.md` (the per-surface primitives + schemas)
- `corpus/ops/staging_from_dump.md` (the pre-embedded skiller snapshot + Directus content replay)

## Delivers → `rosetta-extensions/stack-seeding/seeders/` + presets; updates `corpus/ops/seeding-spec.md`
The seeder fleet + the backdated activity generator + the seed presets; extends the spec with the per-seeder
catalog (each seeder's surface, primitive, isolation class, conformance gene).

## Exit (iterative)
Close on a **Gate Outcome Ledger** when the exit gate fires (coverage ≥ 90% / critical 100%, login 200, < 2 min,
zero shared writes) — OR a user pragmatic-close mandate, with `carry-forward.md` routing any un-covered surfaces.
