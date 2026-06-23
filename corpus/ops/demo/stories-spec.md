# The Verified-Skill Chain — Stories & Heroes (spec)

**The reference for the seeded *verified skill* — the spine of a believable demo world.** A demo or dev
stack's headline surfaces (a person's **skill profile** + Skill Spotlight chart, the org **Workforce
dashboard**) are driven by *verified* skills: skills a person proved by passing AI simulations. This doc is the
canonical description of how `rosetta-extensions/stack-seeding` materializes one — the **7-table chain** the
`PersonaSeeder` writes — plus the constraints that make a seeded verified skill actually *render* (and the ones
that silently hide it if you get them wrong) — **and** (v1.9 M35) the declarative **Stories & Heroes** model
that lifts it into a multi-org, thriving/struggling/manager-trio demo world.

> **Scope (v1.9 "storytelling").** This doc covers the **verified-skill chain** delivered in **M34** — the
> `PersonaSeeder`, the `TaxonomyRefs` resolver, the `jobsim_sessions.go` G14 fix, the `users.go`
> name/avatar/email patch, and the **seed-side closure gene** — plus the **Stories & Heroes model + multi-org**
> delivered in **M35** (the `stack.stories.yaml` blueprint, per-story `OrgID`, the
> thriving/struggling/manager hero **trio**, the vantage/trajectory axes, supporting-population fidelity — see
> [§ The Stories & Heroes model (M35)](#the-stories--heroes-model-m35) below) — plus the **Workforce dashboard
> surfaces** delivered in **M36** (the mapped→verified funnel, teams/tags + the mentor tag, target-roles
> gap+mobility, the succession interview feeders, ~2:1 feedback, the assignment status-mix fix, and the
> org-scale claimed-vs-verified distribution — see
> [§ The Workforce dashboard surfaces (M36)](#the-workforce-dashboard-surfaces-m36) below). The presenter cockpit
> + the Clerkenstein multi-identity seat-switch are M37–M38; this doc is the foundation those build on. It graduates
> the adversarially-verified analysis (the gitignored `.agentspace/seeding_gaps.md`) into the corpus. The code
> lives in the gitignored `rosetta-extensions` monorepo (its own git; authored + tagged in
> `.agentspace/rosetta-extensions/`, consumed per-stack at a pinned tag) — **no platform repo is modified.**
> Pairs with [`seeding-spec.md`](../seeding-spec.md) (the framework) + [`safety.md`](../safety.md) (the
> isolation boundary).

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
blueprint without personas seeds exactly as before. The full multi-org model is below.

## The Stories & Heroes model (M35)

M35 graduates the single-hero `personas` shape into a declarative **`stories[]`** model: **one
`stack.stories.yaml` seeds multiple orgs**, each with a **thriving / struggling / manager hero trio** at
vantage-appropriate fidelity. It **supersedes** the org-centric single-org blueprint for a believable demo
world; the size-tier presets (`dev-min`, `small-200`, …) stay for light structural dev seeds.

### The shape

```yaml
stack: demo-1
stories:
  - id: ai-transformation               # seeds this story's deterministic OrgID
    name: "AI Transformation & Reskilling"
    org: { name: Cervato Systems, slug: cervato-systems, industry: software, narrative: ai-transformation, size: 220, activity: { months: 6, pass_rate: 0.7 } }
    heroes:
      - id: maya-thriving
        name: Maya Chen
        role: Backend Developer          # a public job_role WITH role-skills (role-coherent — see O6)
        vantage: end-user                # the SEAT: end-user (individual surfaces) | manager (org dashboard)
        trajectory: thriving             # the ARC: thriving | struggling (end-users only)
        skills: { verified: 8, mapped: 22, category_breadth: 4, self_eval_bias: under, arc: rising }
      - { id: tom-struggling, name: Tom Becker, role: Backend Developer, vantage: end-user, trajectory: struggling, skills: { verified: 2, self_eval_bias: over, arc: flat } }
      - { id: dan-manager,    name: Dan Rossi,  role: Engineering Manager, vantage: manager }   # rides the aggregates
  - id: sales-ramp
    org: { name: Solvantis, slug: solvantis, industry: saas-sales, narrative: onboarding-ramp, size: 120 }
    heroes: [ … Sara·thriving / Nick·struggling / Leah·manager … ]
```

The shipped `presets/stories.seed.yaml` is the runnable locked **2-stories × 3-heroes** roster.

### Vantage & trajectory (the two hero axes)

- **`vantage`** (`end-user | manager`) — the **seat** into the product. An end-user hero demos the *individual*
  surfaces (profile, Spotlight, my-growth); a manager hero demos the *org-intelligence* surfaces (Workforce
  dashboard, team gaps, succession). **A manager seeds NO verified-skill chain of her own** — she reads the org
  aggregates her employee heroes populate (the coherence property: the two employees ARE her standout high/low
  rows). `""` defaults to end-user.
- **`trajectory`** (`thriving | struggling`) — the **arc** of an end-user. It drives the verified-skill
  fidelity AND (when `self_eval_bias` is unset) the claimed-vs-verified bias:
  - **thriving** → dense (more verified skills) · high level band (~L4) · rising growth arc · **under-claims**.
  - **struggling** → sparse (few) · low band (~L1–L2) · flat arc · **over-claims** (the stark gap the demo
    turns on).

Each hero's `skills:` block (`verified` / `mapped` / `category_breadth` / `self_eval_bias` / `arc`) **overrides**
the trajectory-derived defaults when present; absent, the trajectory drives everything.

### Multi-org — per-story `OrgID`, the platform's real multi-tenancy

Each story is **one org**. The seeder derives a **deterministic per-story `OrgID`** and threads it through
**every** seeder (`org` / `users` / `identity` / `jobsim-sessions` / `skillpath-sessions` / `assignments` /
`activity` / `personas`), so all orgs live in one stack's per-stack Postgres scoped by `organization_id` —
mirroring the platform's real multi-tenancy (so it's *more* realistic, not a hack). Two key rules:

- **The FIRST story keeps the Clerkenstein default org** (`22222222-…` = `DefaultDemoUser().OrgEid`, clerk id
  `org_clerkenstein`) so a **single-identity demo login** (the only mode until M37) lands in it. Every later
  story gets its own deterministic id. The demo identity (`IdentitySeeder`) is seeded **only** for the first
  story — the per-story selectable identity registry is M37.
- **Per-story id namespacing.** Two stories' populations never collide: the first story uses the bare stack
  key, later stories prefix with `:story:<id>`. (Keeping the first story bare is what preserves the M34
  single-org ids byte-for-byte — the legacy single-org path is unchanged.)

### Hero placement — collision-free declaration-order slots

A hero rides on a population user row (D2): the `UsersSeeder` writes her real name at the slot the
`PersonaSeeder` verifies against. M35 assigns heroes the **first `len(heroes)` population slots in declaration
order** (collision-free by construction; the blueprint's `len(heroes) <= size` validation guarantees they fit).
*(M34 hashed heroes into the population, which collided ~10% of the time for a trio in a 30-person org — sharing
a manager's row with an employee's name. The declaration-order fix makes the trio's rows correct, not just
visible; a non-fatal warning still fires for the residual Size<heroes clamp case.)*

### Supporting-population fidelity

Every member (not just heroes) gets a **real replayed job role** — `memberships.job_role_id` is set to a real
public `job_roles.node_id` (the `J-XXXXXX-XXXX` form, NOT a uuid) drawn at run time from the **replayed** stack's
roles that *have* role-skills (`jobroleref.go`, mirroring the skill-side `TaxonomyRefs`; no taxonomy replayed →
left NULL, never fabricated) — plus a **ramped `joined_at`** (`memberships.created_at` spread across the activity
span, so the org reads as one that grew over time). The trio sits in a believable org.

### Closure across all orgs

The seed-side closure gene is **org-agnostic** (it counts every dangling seeded skill ref vs the replayed
taxonomy, with no org filter), so `datadna measure-closure` proves **0 dangling refs across all orgs** — the
multi-org world is as closed as the single-org one. As of M36 the gene spans **four** seeded skill-ref surfaces
(it now also covers `public.membership_skills.skill_id`, the dashboard mapped surface — see below).

## The Workforce dashboard surfaces (M36)

M34/M35 made the **individual** profile believable (Must #1). M36 makes the **org Workforce-Intelligence
dashboard** (REST `/api/workforce/*`, the org-admin view) believable for the seeded story (Must #2): every
aggregate renders **non-empty and distributed**, not binary-or-zero. A manager hero (e.g. Dan Rossi) logs in and
sees her two employee heroes as the standout high/low rows of an org that reads real around them. Six new
seeders + two fixes land the **spine** (not every widget — the hard scope line):

| Surface | Seeder (`stack-seeding/seeders/`) | What it feeds on the dashboard |
|---|---|---|
| **Mapped skills** | `membership_skills.go` | The mapped→verified **verification funnel**. Every member is mapped to a role-coherent set of real public skills (the `skill_name` is set — every dashboard query filters it NOT NULL); since mapped covers ~all members but only a subset verify, **mapped outnumbers verified per skill** → the believable drop-off. Also feeds the **AI-readiness** scan (an AI-narrative org biases a share of members toward AI-named skills). |
| **Teams / tags** | `tags.go` | The universal **slice dimension**: a dozen business-unit tags (front-loaded so the Teams tab is non-uniform) + a cross-cutting **`mentor`** tag (the Growth-tab Mentors KPI counts members tagged `mentor`). Each member is on exactly one business unit. |
| **Target roles** | `target_roles.go` | The **gap + two-way internal mobility**: `organization_target_roles` (an admin-set development target = the gap) + `user_target_roles` (a self-set aspiration = mobility-ready), each a real public role node-id chosen different from the member's current role. |
| **Succession feeders** | `succession.go` | `interview_extraction_results` for >20% of members (with the `summary` jsonb the succession query reads) to lift the **Succession tab** past the coverage gate (`too_sparse` → `full`). Trajectory-aware: a struggling hero reads at-risk (low wellbeing + negative sentiment), a thriving one reads positive. (The other feeder, `validation_attempt_*`, already lands via the M34 chain.) |
| **Feedback** | `feedback.go` | `job_simulation_feedbacks` at **~2:1 positive** (the Italgas anchor), `is_positive` matched to the option's polarity — the "people liked it" signal. |
| **Org-scale gap** | `population_evidence.go` | The **claimed-vs-verified gap at org scale** (the headline "aha", §3c): a ~55% share of the *supporting* population (not just heroes) gets verified-skill evidence rows with both `user_level` (claim) and `anthropos_level` (verified) set and **diverging** — a population mix of over- and under-claimers. (`user_skill_evidences` has no FK on `jobsimulation_session_id`, so a population evidence row is a clean write without the full hero chain; the heroes' full 7-table chain is the PersonaSeeder's.) |

Two **fixes** (not new seeders) round out the spine:

- **Assignments status mix** (`assignments.go`) — the pre-M36 seeder wrote every assignment as bare
  `active` / no `due_date` / no session, so the dashboard bucketed **all** of them as `not_started`. M36 gives
  each a deterministic lifecycle bucket (~35% completed / ~15% overdue / ~35% in-progress / ~15% not-started)
  realized via `status` + a past/future `due_date` + (for completed/in-progress) an
  `organization_assignment_sessions` row carrying progress. That session FKs a `local_skill_path_sessions` row
  the seeder also writes — the dashboard reads the **app mirror** `local_skill_path_sessions`, NOT
  `skillpath.skill_path_sessions`, and the population has no `local_jobsimulation_sessions` mirror, so the
  session takes the skill-path arm of the table's check constraint.
- **Skillpath completed share** (`skillpath_sessions.go`) — the learning-session seeder marked `completed`
  only on an exact `progress=100` (~1%), starving the learning/my-growth surfaces. M36 makes ~30% complete.

**Distribution coherence.** The growth arc (Growth/Biggest-Improvers) and the over/under-claimer split are
*correlated with the heroes*: the thriving employee is the dashboard's top-performer / succession candidate; the
struggling employee is the at-risk / needs-attention row — the **same** coherence property M35 established,
surfaced now at org scale.

**Closure stays measured.** Every new skill-ref surface draws its node-ids from the **same** taxonomy resolver
the chain uses, so the seed-side closure gene (extended to `membership_skills`) still proves **0 dangling refs**
— the mapped funnel is as closed as the verified chain.

## Running it

```bash
# 1. Replay the public reference library FIRST (a --local-content demo does this automatically, cache-first):
stacksnap replay --surface taxonomy --stack demo-N   # the skill node-ids the chain draws from
stacksnap replay --surface directus --stack demo-N   # the sim templates the sessions link to

# 2. Seed the world. The M34 vertical slice (one hero):
stackseed --stack demo-N --seed presets/stories-maya.seed.yaml
#    …or the M35 full multi-org roster (2 orgs × the thriving/struggling/manager trio):
stackseed --stack demo-N --seed presets/stories.seed.yaml

# 3. Prove closure (every seeded skill ref resolves in the replayed taxonomy — all orgs):
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
