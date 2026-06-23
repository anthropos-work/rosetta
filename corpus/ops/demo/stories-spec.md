# The Verified-Skill Chain — Stories & Heroes (spec)

**The reference for the seeded *verified skill* — the spine of a believable demo world.** A demo or dev
stack's headline surfaces (a person's **skill profile** + Skill Spotlight chart, the org **Workforce
dashboard**) are driven by *verified* skills: skills a person proved by passing AI simulations. This doc is the
canonical description of how `rosetta-extensions/stack-seeding` materializes one — the **7-table chain** the
`PersonaSeeder` writes — plus the constraints that make a seeded verified skill actually *render* (and the ones
that silently hide it if you get them wrong).

> **Scope (v1.9 "storytelling").** This doc covers the **verified-skill chain** delivered in **M34** — the
> `PersonaSeeder`, the `TaxonomyRefs` resolver, the `jobsim_sessions.go` G14 fix, the `users.go`
> name/avatar/email patch, and the **seed-side closure gene**. The full declarative **Stories & Heroes** model
> (`stack.stories.yaml`, multi-org, the thriving/struggling/manager hero **trio**, the presenter cockpit) is
> M35–M38; this doc is the foundation those build on. It graduates the adversarially-verified analysis (the
> gitignored `.agentspace/seeding_gaps.md`) into the corpus. The code lives in the gitignored
> `rosetta-extensions` monorepo (its own git; authored + tagged in `.agentspace/rosetta-extensions/`, consumed
> per-stack at a pinned tag) — **no platform repo is modified.** Pairs with
> [`seeding-spec.md`](../seeding-spec.md) (the framework) + [`safety.md`](../safety.md) (the isolation
> boundary).

## For PMs — what it does and why it matters

A demo only sells if the buyer sees the product's *point*: a person whose skills are **verified by the
platform**, with the gap between what they **claim** and what they've **proven** laid bare. Before M34 the
seeder built a populated-but-hollow world — users named "User 1", **zero** verified skills (the core profile
screen was empty), and session rows that were silently broken so nothing rolled up to the dashboard. M34 makes
**one hero (Maya Chen)** real, end-to-end: she has a real name and face, a profile of verified skills that
renders, a Skill Spotlight chart that plots her trajectory, and a visible **claimed-vs-verified gap** — the
single most convincing moment in the demo. It does this by writing the exact database rows the platform itself
would have produced after Maya passed a batch of simulations, **without running a single real simulation** and
**without touching production**.

## For engineers — the mechanism

### A verified skill is a 7-table fan-out, not one row

The platform produces a verified skill asynchronously: a passed `jobsimulation.session` fires a Redis-Streams
event that the `app` service's `JobsimulationSessionEndedHandler` turns into validation rows and the public
profile mirrors. The `PersonaSeeder` **short-circuits** that pipeline and writes the terminal rows directly —
per (hero × verified skill), across **three Postgres schemas**:

```
            schema: jobsimulation                         │  schema: public (app service)
 sessions (sim_type ASSESSMENT|HIRING, status='ended')  ┐ │
   └─ validation_attempt_results (evaluation_status='passed') │
        └─ validation_attempt_skill_results ◄───────────┘ │   (skill=NodeID, competency_level_score>0)
             └─ validation_criterion_results (3/session)  │
 ───────────────────────────────────────────────────────  │
                                                          ├─ local_jobsimulation_sessions (the app mirror)
                                                          ├─ user_skills (is_verified=true, job_simulation_id)
                                                          └─ user_skill_evidences (UPSERT: levels+counts+verified)
 skiller.skills.node_id ── supplies the skill_id string (a loose ref, NOT a FK) ──┘
```

The seeder writes **3 passed sessions per verified skill** (the Spotlight chart hides below 2 datapoints), each
through the full chain, then one `user_skills` row per session and one UPSERTed `user_skill_evidences` per
skill. For Maya with 8 verified skills that's ~200 rows — a complete, queryable profile.

### The constraint landmines (verified against the live schema)

A raw COPY/INSERT lets you write rows the *running platform never would* — so the seeder has to honour every
constraint the app layer normally enforces. Two classes:

**DB-enforced (a violating INSERT is *rejected*):**
- `user_skills` CHECK `user_skills_check_foreign_keys` — at least one provenance edge must be non-NULL ⇒ the
  seeder sets **`job_simulation_id`** (the SIMULATION/Directus template UUID, *not* the session UUID).
- `user_skills` partial UNIQUE `idx_unique_job_simulation (skill_id, job_simulation_id, user_skill_user)` ⇒ a
  **distinct real sim template per verified row** (the seeder draws a distinct sim per session).
- `user_skill_evidences` UNIQUE `(skill_id, user_id)` ⇒ the evidence is an **UPSERT** (`INSERT … ON CONFLICT
  (skill_id, user_id) DO UPDATE`), never a blind insert — the fleet's `id`-keyed `CopyRowsIdempotent` can't
  dedup a composite UNIQUE, so this one is a per-row `Exec` (#M34-D3).
- `validation_attempt_skill_results.validation_attempt_result_id` is **NOT NULL** ⇒ the FK to its attempt
  result must be set. *(This one bit during M34 build — the seeder omitted it and the chain failed to insert;
  the integration test caught it. The unit test now asserts it.)*

**Logical (the row INSERTs fine but the UI hides or breaks it — the "inserted-but-invisible" class):**
- The session's `status` / `completion_status` / `result_status` / `sim_type` columns are **free-text
  `varchar`** (no PG enum, no CHECK). A wrong value is **filtered out of every dashboard query** at read time,
  not rejected at write. This was the **G14** bug: the pre-M34 session seeder wrote `status='completed'` (no
  such value — valid is `'ended'`), `completion_status='completed'` (valid: `passed|failed|…`),
  `result_status='passed'/'failed'` (valid: `'completed'`), and bare `'assessment'`/`'interview'`/`'training'`
  for `sim_type` (the real values are the **full `SIMULATION_TYPE_*` proto strings**, and only
  `SIMULATION_TYPE_ASSESSMENT|_HIRING` roll up to verified skills). All four are now written correctly,
  confirmed against the live value distribution.
- Levels are stored **0–100** and divided to the org scale at query (`level × maxLevel / 100`, `maxLevel`
  defaults to 5 — no settings seed needed). Store ~80 to read as 4/5. *(Two distinct scales: the
  `user_skill_evidences.*_level` columns + validation scores are 0–100; `user_skills.level` is a separate
  ~1–5 convention with no DB bound. Don't conflate.)*
- The misspelled column **`local_jobsimulation_sessions.completition_status`** (sic) — copy it exactly.
- The session `token` is a free-text varchar but the app-layer validator bounds it to `^[a-z0-9]{5,10}$` ⇒ the
  seeder writes a 7-char hex token, not a 36-char UUID.

### The claimed-vs-verified gap — set `user_level` (the reference omission)

The demo's headline widget diffs **`user_skill_evidences.user_level`** (what the person *claims*) against
**`anthropos_level`** (what the simulations *verified*). The reference `seed.sql` the chain was ported from
sets `level` and `anthropos_level` but **never `user_level`** — porting it verbatim leaves `user_level` NULL,
and the widget **excludes every row**. The `PersonaSeeder` therefore sets `user_level` explicitly per the
hero's **`self_eval_bias`**: `under` (a modest under-claimer — `user_level < anthropos_level`, the thriving
arc), `over` (an inflated over-claimer — the stark gap, the struggling arc), or `calibrated` (≈ equal). At
least one skill per hero must show a real gap or the widget is empty.

### Closure — real skill node-ids, never fabricated, and *measured*

A skill ref (`user_skills.skill_id`, `user_skill_evidences.skill_id`,
`validation_attempt_skill_results.skill`) is a loose string (`K-XXXXXX-XXXX`), **not a DB FK**. A fabricated
node-id passes every field-regex but **dangles** — it resolves to no skill in `skiller`, so the profile
federates a blank name/category and the chart has a hole (the skill-side analog of the M23 empty-picker class).
Two pieces guarantee closure:

- **The `TaxonomyRefs` resolver** draws every skill node-id from the **real replayed public `skiller`
  taxonomy** — role-coherent where the hero's role resolves (`skillsByRole`: `job_roles ⋈ job_role_skills ⋈
  skills`, public-only, is_core-first), falling back to a flat public-skill pool otherwise. If **no** taxonomy
  has been replayed (an empty pool), the seeder **skips** the hero — it **never fabricates** a node-id.
- **The seed-side closure gene** (`datadna measure-closure --stack demo-N`) then *proves* it: it counts the
  distinct seeded skill node-ids (across all three ref surfaces) that don't resolve in the replayed
  `skiller.skills` — must be **0**, naming a sample on failure. This mirrors the M23
  `snapshot-cross-surface-closure` gene (the content side); together they are the closure family (see
  [`../../architecture/`](../../architecture/) and the `dna/README.md` in the extension). "Believable" is
  **measured, not assumed.**

### The supporting fixes (population believability)

- **`jobsim_sessions.go` (G14)** — beyond the valid enums, the session score is now a **continuous, mid-skewed
  0–100 distribution with a per-user upward growth arc** (replacing the binary 85/35), and `sim_type` is
  weighted ~75% toward the verification-feeding ASSESSMENT/HIRING types. The arc gives the dashboard's
  Growth/Biggest-Improvers a real trend to narrate (the "company mid AI-transformation" story).
- **`users.go`** — real first/last names from an in-code name bank, a deterministic avatar URL, and a
  `first.last@<org-domain>` email (no more "User N" / no picture / `@{stack}.local`). A blueprint **hero**'s
  real name + email land at the population index the `PersonaSeeder` verifies her chain against, so the named
  row and the verified skills are one user — both seeders derive that index from the same shared bridge
  (`personaUserIndex` / `personaIndexMap`), so they cannot drift (#M34-D2).

## The blueprint — declaring a hero

A hero is a `personas` entry in the stack blueprint (`stack.seed.yaml`). The M34 vertical-slice shape:

```yaml
personas:
  - id: maya-thriving       # stable key (seeds her deterministic user index)
    name: Maya Chen
    role: Backend Developer  # a public job_role with role-skills → role-coherent verified skills
    verified: 8              # number of distinct skills to verify (the chain runs once per skill)
    self_eval_bias: under    # under | over | calibrated — drives the claimed-vs-verified gap
```

The shipped `presets/stories-maya.seed.yaml` is the runnable vertical-slice world (Cervato Systems + Maya).
The `personas` list is **optional** and **additive** — heroes ride on top of the generic population, so a
blueprint without personas seeds exactly as before. M35 graduates this to the full `stack.stories.yaml`
(multi-org, the hero trio, vantage/trajectory, the cockpit).

## Running it

```bash
# 1. Replay the public reference library FIRST (a --local-content demo does this automatically, cache-first):
stacksnap replay --surface taxonomy --stack demo-N   # the skill node-ids the chain draws from
stacksnap replay --surface directus --stack demo-N   # the sim templates the sessions link to

# 2. Seed the world (the personas drive the verified-skill chain):
stackseed --stack demo-N --seed presets/stories-maya.seed.yaml

# 3. Prove closure (every seeded skill ref resolves in the replayed taxonomy):
datadna measure-closure --stack demo-N               # exit 1 on any dangling ref
```

Without a taxonomy replay the `PersonaSeeder` skips the hero (no fabrication) and her profile is empty — the
taxonomy snapshot is an explicit prerequisite for the verified-skill chain, surfaced by an empty pool, not a
silent failure. A `--local-content` demo stack replays both by default, so a bare `/demo-up N` already lands a
seeded, set-dressed world.

## Safety

Every table the chain writes — `jobsimulation.sessions`, `validation_attempt_*`,
`validation_criterion_results`, `public.local_jobsimulation_sessions`, `public.user_skills`,
`public.user_skill_evidences` — is **`PerStackIsolated`** (the stack's own offset-port Postgres container), so
the chain inherits the existing zero-pollution posture: it cannot touch production or another stack, and the
seeding run's audit log proves it. See [`../safety.md`](../safety.md) §2.1. The taxonomy it reads is the
**public** reference data the snapshot firewall already guaranteed public-only at capture.
