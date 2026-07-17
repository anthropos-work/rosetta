# The Verified-Skill Chain — Stories & Heroes (spec)

**The reference for the seeded *verified skill* — the spine of a believable demo world.** A demo or dev
stack's headline surfaces (a person's **skill profile** + Skill Spotlight chart, the org **Workforce
dashboard**) are driven by *verified* skills: skills a person proved by passing AI simulations. This doc is the
canonical description of how `rosetta-extensions/stack-seeding` materializes one — the **7-table chain** the
`PersonaSeeder` writes — plus the constraints that make a seeded verified skill actually *render* (and the ones
that silently hide it if you get them wrong) — **and** (v1.9 M35) the declarative **Stories & Heroes** model
that lifts it into a multi-org, thriving/struggling/manager-trio demo world.

> **The demo-patch mechanism is specified in [`demopatch-spec.md`](demopatch-spec.md).** It is the sanctioned **zero-platform-edit escape hatch**: patch the demo's own ephemeral clone before the image build, revert after — the canonical repos are never touched. Read it before adding or re-pinning a patch. Since M217 the gate is **self-healing**: the *anchor* is the contract, the whole-file sha is only a baseline.

> **Scope (v1.9 "storytelling").** This doc covers the **verified-skill chain** delivered in **M34** — the
> `PersonaSeeder`, the `TaxonomyRefs` resolver, the `jobsim_sessions.go` G14 fix, the `users.go`
> name/avatar/email patch, and the **seed-side closure gene** — plus the **Stories & Heroes model + multi-org**
> delivered in **M35** (the `stack.stories.yaml` blueprint, per-story `OrgID`, the
> thriving/struggling/manager hero **trio**, the vantage/trajectory axes, supporting-population fidelity — see
> [§ The Stories & Heroes model (M35)](#the-stories--heroes-model-m35) below) — plus the **Workforce dashboard
> surfaces** delivered in **M36** (the mapped→verified funnel, teams/tags + the mentor tag, target-roles
> gap+mobility, the succession interview feeders, ~2:1 feedback, the assignment status-mix fix, and the
> org-scale claimed-vs-verified distribution — see
> [§ The Workforce dashboard surfaces (M36)](#the-workforce-dashboard-surfaces-m36) below) — plus the
> **Clerkenstein multi-identity seat-switch** delivered in **M37** (a demo can present as any seeded hero; see
> [`clerkenstein/knowledge/architecture.md` § Multi-identity]) — plus the **presenter cockpit** delivered in
> **M38** (a standalone served panel that lists each story → its hero trio with a **[Log in as]** action, so a
> demo-giver picks a hero and lands on the right screen to present a flow live — UX-specced standalone in
> [`cockpit-spec.md`](cockpit-spec.md) since v1.10 M38→M43, see [§ The presenter cockpit (M38)](#the-presenter-cockpit-m38)
> below) — plus the **profile-identity layer**
> delivered in **v1.10 "method acting" M39** (the roster org-name thread → the real company on the top bar, the
> `user_basic_info` role backfill → a real role+title on the /profile header, and the offline real-face avatar —
> see [§ The profile-identity layer (v1.10 M39)](#the-profile-identity-layer-v110-method-acting-m39) below) — plus the
> **profile-depth layer** delivered in **v1.10 M41** (the new `ProfileSeeder`: a believable work-history +
> education timeline, a verified-skill depth bump `8 → ~30`, and a ~60-skill claimed-but-unverified tail that
> widens the visible claimed-vs-verified gap — see
> [§ The profile-depth layer (v1.10 M41)](#the-profile-depth-layer-v110-method-acting-m41) below). It graduates
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
 public.skills.node_id ── supplies the skill_id string (a loose ref, NOT a FK) ──┘
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

> **Trajectory-aware self-rating (v1.10 M44 §A).** A struggling hero now models the **"hasn't done the
> initial self-assessment"** state: `PersonaSeeder` writes `user_level` **only** for a **self-rated** hero
> (the new `Persona.EffectiveSelfRated()` — struggling = false → `user_level` **NULL**; thriving / calibrated /
> manager = true). A thriving hero shows a **completed** self-assessment; a struggling hero's verified skills
> still render (the chart reads `anthropos_level`, which is untouched), but the claimed side is **absent** — the
> profile reads as a sparse/incomplete self-rating rather than the M35 over-claim. The org-scale population
> (`PopulationEvidenceSeeder`) is unaffected — every population member still self-rates (the over/under-claim
> mix is the headline org "aha"). Full rubric:
> [`profile-completeness-spec.md`](profile-completeness-spec.md).

### Closure — real skill node-ids, never fabricated, and *measured*

A skill ref (`user_skills.skill_id`, `user_skill_evidences.skill_id`,
`validation_attempt_skill_results.skill`) is a loose string (`K-XXXXXX-XXXX`), **not a DB FK**. A fabricated
node-id passes every field-regex but **dangles** — it resolves to no skill in the `public` taxonomy, so the profile
federates a blank name/category and the chart has a hole (the skill-side analog of the M23 empty-picker class).
Two pieces guarantee closure:

- **The `TaxonomyRefs` resolver** draws every skill node-id from the **real replayed public
  taxonomy** (the `public`-schema skills/roles catalog) — role-coherent where the hero's role resolves (`skillsByRole`: `job_roles ⋈ job_role_skills ⋈
  skills`, public-only, is_core-first), falling back to a flat public-skill pool otherwise. If **no** taxonomy
  has been replayed (an empty pool), the seeder **skips** the hero — it **never fabricates** a node-id.
- **The seed-side closure gene** (`datadna measure-closure --stack demo-N`) then *proves* it: it counts the
  distinct seeded skill node-ids (across all three ref surfaces) that don't resolve in the replayed
  `public.skills` — must be **0**, naming a sample on failure. This mirrors the M23
  `snapshot-cross-surface-closure` gene (the content side); together they are the closure family (see
  [`../../architecture/`](../../architecture/) and the `dna/README.md` in the extension). "Believable" is
  **measured, not assumed.**

### The supporting fixes (population believability)

- **`jobsim_sessions.go` (G14)** — beyond the valid enums, the session score is now a **continuous, mid-skewed
  0–100 distribution with a per-user upward growth arc** (replacing the binary 85/35), and `sim_type` is
  weighted ~75% toward the verification-feeding ASSESSMENT/HIRING types. The arc gives the dashboard's
  Growth/Biggest-Improvers a real trend to narrate (the "company mid AI-transformation" story).
- **`users.go`** — real first/last names from an in-code name bank, a deterministic **real-face avatar**
  (v1.10 M39 — see below), and a `first.last@<org-domain>` email (no more "User N" / no picture /
  `@{stack}.local`). A blueprint **hero**'s real name + email land at the population index the `PersonaSeeder`
  verifies her chain against, so the named row and the verified skills are one user — both seeders derive that
  index from the same shared bridge (`personaUserIndex` / `personaIndexMap`), so they cannot drift (#M34-D2).
  v1.10 M39 also has `users.go` **backfill `public.user_basic_info`** (the /profile header source) — see the
  profile-identity layer below.

### The profile-identity layer (v1.10 "method acting" M39)

v1.9 seeded the verified-skill **spine**, but a logged-in hero's **profile page** still read thin: the top bar
showed a generic org, the header showed "no role", and the avatar was a 2-letter initials disc. M39 lights the
three highest-leverage, lowest-effort identity fixes — **tooling + docs only, zero platform-repo edits** — so a
hero (Maya on demo-3) shows the **right company, a real role+title, and a real face**.

- **Org name (G1) — the roster carries the story org.** The FAPI org resource used to hardcode "Clerkenstein
  Demo Org". Now `roster.go` threads each hero's `st.Org.Name`/slug into the roster JSON (`org_name`/`org_slug`),
  Clerkenstein renders it on the org resource, and the **top bar reads the real company** (e.g. "Cervato
  Systems"). It's a `DisallowUnknownFields` **paired change** across the two structs + a no-roster
  `"Clerkenstein Demo Org"` fallback — the full mechanism is in
  [`../../services/clerkenstein.md` § Roster org-name threading](../../services/clerkenstein.md#multi-identity).

- **Role backfill (G2) — `user_basic_info`, the table the header actually reads.** The /profile header reads
  `profile.info.jobRole` → `infoResolver.JobRole` ← **`public.user_basic_info.job_role_id`** — but the seeder
  wrote the role only to `public.memberships` (the wrong table for the header), so the header showed "no role".
  M39 has `users.go` **backfill `user_basic_info`** (`job_role_id` + `job_title` + a deterministic believable
  `summary` + `location`) from the **same resolved role** it writes to the membership. The trigger-created row
  already exists, so it's an **idempotent UPDATE keyed by `id`** with an `IS DISTINCT FROM` guard (a re-seed of
  identical data matches 0 rows — the M17 re-run contract — #M39-D4/D5). **One UPDATE lights two surfaces**: the header
  role/title **and** the role-gap radar / role-readiness widgets (`jobRoleMatch` keys off the same field). The
  no-fabrication rule holds: `job_role_id` is NULL with no replayed taxonomy, and a hero keeps her **declared**
  role label as the title (the same split `users.go` applies to `memberships.job_role_name`). Backfilled for
  **every** member (not heroes only) so any profile a presenter clicks into reads coherent. Real schema (no
  `job_role_title` column — the header uses `job_role_id` → resolved label + `job_title`; `email` is NOT NULL
  UNIQUE, which is why it must be an UPDATE not an INSERT).

- **Real-face avatars (G4 → M42e P4) — offline, deterministic, license-clean, now PHOTOREALISTIC.** `users.picture`
  was a DiceBear *initials* SVG fetched from `api.dicebear.com` — a 2-letter disc **and** an online fetch a sealed
  demo box can't reach. M39 replaced it with a self-authored **parametric SVG face generator** (an illustrated
  cartoon face). **v1.10 "method acting" M42e P4** (user decision: SYNTHETIC photorealistic faces of non-existent
  people) replaces the cartoon with a **real photorealistic synthetic face**: `avatar.go`'s `photoAvatarDataURI`
  picks one of **12 bundled StyleGAN2 / "this-person-does-not-exist"-class portraits** (`stack-seeding/assets/avatars/`,
  `go:embed`-ed, 160×160 JPEG ~5–7 KB) deterministically by `hash(uuid) % 12` and emits a **`data:image/jpeg;base64,…`
  URI** into `users.picture`. The faces depict **NO real person** (synthetic ⇒ no consent/privacy; machine-generated
  ⇒ no copyright — see `assets/avatars/LICENSE.md`), so they stay **license-clean** while being a real photo. Still
  **offline-safe** (the photo is bundled + embedded, zero fetch) and **deterministic** (same user → same face,
  reruns byte-identical). The illustrated-SVG generator is retained as the dependency-free fallback. **Menu ==
  profile (M42e P4):** the SAME data URI threads to BOTH `users.picture` (the /profile avatar) AND the Clerkenstein
  roster `Picture` → `DemoUser.Picture` → FAPI `userRes.image_url`/`has_image` (the top-menu avatar) — proven
  byte-identical (a re-seeded hero's `users.picture` SHA256 == her roster `picture` SHA256). The **org logo** rides
  the same path: the seeded `organizations.logo_url` monogram threads through the roster `OrgLogo` → `orgRes.image_url`
  so the top-menu org glyph renders the seeded mark.

The live acceptance: re-seed demo-3 + log in as Maya → the top bar shows "Cervato Systems", the profile shows a
role + title + summary + location, and every person carries a real face. Code-of-record: `rosetta-extensions`
@ tag `method-acting-m39`. Out of M39 scope (later milestones): work/education history + skill depth (M41); the
library + activity-feed serve-grant (M40).

### The profile-depth layer (v1.10 "method acting" M41)

M39 gave the logged-in hero the right **identity** (company, role+title, face); M41 gives her the **depth** behind
it: a believable **work history + education timeline** and a **deep, role-aligned skill set with a wide, obvious
claimed-vs-verified gap**. Before M41 the `/profile` timeline was empty (`public.user_experiences` /
`public.user_educations` were **0 rows DB-wide** — written by no seeder) and the skill set was shallow (preset
`verified: 8` → 8 distinct verified skills). M41 adds a new **`ProfileSeeder`** (surface `"profiles"`) and bumps
the depth — **tooling + docs only, zero platform-repo edits**; the `/profile` timeline reads
`ent.UserExperience` / `ent.UserEducation` via `TimelineGrouped(userID)` unchanged — M41 only supplies the rows.

> **M44 §C update:** a **manager** is no longer skipped — she now gets a **modest** personal profile (a flat
> 3-8 verified skills + a manager-track timeline + a small claimed tail) so her OWN `/profile` is populated.
> The "skipped" claims below describe the M41 baseline; see
> [`profile-completeness-spec.md`](profile-completeness-spec.md) § per-vantage for the manager + bulk-member fills.

- **Work history + education (G3).** Per **end-user** hero (M41 baseline — managers got a personal timeline in
  M44 §C; see the note above), the
  `ProfileSeeder` writes a believable **3-job role progression** (`user_experiences`) + a **degree**
  (`user_educations`), all deterministic + backdated within/just-before the story's activity span so the history
  corroborates the verified skills. The titles reuse the **resolved `jobRoleRefs`** (the same role node-id the
  membership carries), the per-entry `skills` json is a role-coherent slice of **real public skill names** (from
  `resolveNamedSkillRefs`), and the current role is open-ended (`to` NULL). **Live-schema landmines** (the
  overview's column guesses were wrong — these are verified against demo-3): `user_experiences.company` is
  `uuid NOT NULL` with an FK → `companies(id)` (the GraphQL `Company` resolver does `QueryCompany().Only(ctx)`,
  so a NULL/dangling company errors the whole timeline) ⇒ the seeder writes a real **`companies`** row per
  distinct employer (#M41-D2); `from`/`to` are **DATE** (not timestamptz) with a `from<=to OR to IS NULL` CHECK;
  `location_type` is the **lowercase** ent enum `inoffice|hybrid|fullremote` (a wrong-case value inserts but the
  GraphQL `LocationType` enum can't map it); and the `skills` column is **`json`** — an array of skill names.

- **Skill depth + the claimed-but-unverified tail (G5).** The preset `verified:` knob is bumped **8 → ~30** for
  the thriving heroes (`stories.seed.yaml` + `stories-maya.seed.yaml`), so the verified chain writes **~30
  distinct verified skills × `verifiedSessionsPerSkill` (3) ⇒ ~90 `user_skills` + ~30 evidences** on the
  verified side. **On top**, the `ProfileSeeder` seeds a **~60-skill claimed-but-unverified tail**:
  `user_skills` with `is_verified=false`, **no `job_simulation_id`**, and `user_skill_evidences` with
  **`anthropos_level` NULL, `user_level` set** — so the profile "overall" reads **≈ 90 = ~30 verified + ~60
  claimed**, **widening** the visible claimed-vs-verified gap (the demo's headline aha). The **DB landmine**:
  `user_skills_check_foreign_keys` requires ≥1 provenance edge non-NULL — since the tail has no
  `job_simulation_id`, it ties to the seeded **work history** via `user_skill_experience` /
  `user_skill_education` (#M41-D3) (which *also* makes the claimed skills render **under** each work experience —
  the `workExperience.Skills` resolver reads `userskill.HasExperienceWith`). The tail draws skills **distinct**
  from the verified set (it offsets past the first `EffectiveVerified()` of the same role-coherent-then-flat
  combined pool (#M41-D6)), so the two counts don't overlap. Both arcs read coherent: a thriving
  **under-claimer**'s deep profile and a struggling **over-claimer**'s stark gap (few verified, many claimed —
  the tail applies to both arcs (#M41-D5)) are each believable.

The **gap mechanic is intact** — `user_level` (claimed) vs `anthropos_level` (verified) is still the widget's
spine; the unverified tail leaves `anthropos_level` NULL so the gap renders, and the verified evidence UPSERT is
never clobbered (the claimed UPSERT's `ON CONFLICT … WHERE is_verified = false` guard keeps the verified side
winning (#M41-D4)). **Closure stays measured** — every skill node-id/name comes from the same replayed taxonomy resolvers
the verified chain uses; no replayed taxonomy → the timeline still writes (blank skills/role — never fabricated)
and the tail is skipped, so the closure gene stays green. Every table the seeder writes (`companies`,
`user_experiences`, `user_educations`, `user_skills`, `user_skill_evidences`) is **`PerStackIsolated`**.

The live acceptance: re-seed demo-3 + log in as Maya → the `/profile` Work Experience + Education sections
populate with a believable career, and her skill set reads deep (~30 verified + ~60 claimed) with a wide,
obvious claimed-vs-verified gap. Code-of-record: `rosetta-extensions` @ tag `method-acting-m41`. Out of M41
scope (later milestones): the employee/manager 100%-coverage Playwright sweeps (M42e/M42m).

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

The shipped `presets/stories.seed.yaml` is the runnable **3-stories × 3-heroes** roster (the M35 two stories +
the M51 AI-readiness showcase org — see [The AI-readiness showcase org](#the-ai-readiness-showcase-org--the-3rd-story-v110b-fit-up-m51)).

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
  - **Second-world-on-a-shared-stack landmine `(#M202-D4)`.** The default-org slot is **single-tenant** — the
    first story of *any* seed run claims it. So seeding a **second** stories world onto an **already-seeded**
    stack (e.g. the Playthroughs `pt-world` on top of a seeded demo-1) collides: its first story merges into the
    prior world's default org and duplicate-keys on that org's pre-existing rows (a *different* unique constraint
    than the idempotent-on-id COPY). The zero-platform-edit, zero-fork fix is a leading **anchor story** (size 0,
    no heroes) that harmlessly re-declares the default org, pushing the real orgs to story index ≥1 so they get
    their own deterministic `StoryOrgID`s. See [`playthroughs.md`](playthroughs.md) (the Playthroughs dedicated
    seed) for the worked instance.
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
| **Mapped skills** | `membership_skills.go` | The mapped→verified **verification funnel**. Every member is mapped to a role-coherent set of real public skills (the `skill_name` is set — every dashboard query filters it NOT NULL); since mapped covers ~all members but only a subset verify, **mapped outnumbers verified per skill** → the believable drop-off. The funnel joins the mapped side to the verified side **on the skill _name_, not the node-id** — so `membership_skills.skill_name` must equal the verified skills' `public.skills.name`; the seeder's `skillref_named.go` resolver draws names from the same replayed taxonomy the verified chain uses, so they line up by construction (#M36-D1). Also feeds the **AI-readiness** scan (an AI-narrative org biases a share of members toward AI-named skills). |
| **Teams / tags** | `tags.go` | The universal **slice dimension**: a dozen business-unit tags (front-loaded so the Teams tab is non-uniform) + a cross-cutting **`mentor`** tag (the Growth-tab Mentors KPI counts members tagged `mentor`). Each member is on exactly one business unit. |
| **Target roles** | `target_roles.go` | The **gap + two-way internal mobility**: `organization_target_roles` (an admin-set development target = the gap) + `user_target_roles` (a self-set aspiration = mobility-ready), each a real public role node-id chosen different from the member's current role. |
| **Succession feeders** | `succession.go` | `interview_extraction_results` for >20% of members (with the `summary` jsonb the succession query reads) to lift the **Succession tab** past the coverage gate (`too_sparse` → `full`). Trajectory-aware: a struggling hero reads at-risk (low wellbeing + negative sentiment), a thriving one reads positive. (The other feeder, `validation_attempt_*`, already lands via the M34 chain.) |
| **Feedback** | `feedback.go` | `job_simulation_feedbacks` at **~2:1 positive** (the Italgas anchor), `is_positive` matched to the option's polarity — the "people liked it" signal. **v1.10 M42m — the org-feedback mirror fix.** The `/enterprise/organization-feedback` page (`GetOrganizationFeedback`, `app/.../repository/jobsimulation.go`) does NOT read `job_simulation_feedbacks` directly — it **JOINs** feedback to the app mirror `public.local_jobsimulation_sessions` on `feedback.session_id = mirror.jobsimulation_session_id` and scopes by the **mirror's** `organization_id`. The population sessions the feedback links have no mirror (only the `PersonaSeeder` mirrors, for heroes — the M36-D2 "the dashboard reads the app mirror" rule at org scale), so feedback was **inserted-but-invisible** (the page showed "No data" on a fully-seeded org — the org-feedback analog of G14). The `FeedbackSeeder` now **also writes a `local_jobsimulation_sessions` mirror per feedback session** (reconstructing the population session's coherent values from the same deterministic key), so the JOIN resolves and the page renders the real ~2:1 distribution. |
| **Org-scale gap** | `population_evidence.go` | The **claimed-vs-verified gap at org scale** (the headline "aha", §3c): a ~55% share of the *supporting* population (not just heroes) gets verified-skill evidence rows with both `user_level` (claim) and `anthropos_level` (verified) set and **diverging** — a population mix of over- and under-claimers. (`user_skill_evidences` has no FK on `jobsimulation_session_id`, so a population evidence row is a clean write without the full hero chain; the heroes' full 7-table chain is the PersonaSeeder's.) |

Two **fixes** (not new seeders) round out the spine:

- **Assignments status mix** (`assignments.go`) — the pre-M36 seeder wrote every assignment as bare
  `active` / no `due_date` / no session, so the dashboard bucketed **all** of them as `not_started`. M36 gives
  each a deterministic lifecycle bucket (~35% completed / ~15% overdue / ~35% in-progress / ~15% not-started)
  realized via `status` + a past/future `due_date` + (for completed/in-progress) an
  `organization_assignment_sessions` row carrying progress. That session FKs a `local_skill_path_sessions` row
  the seeder also writes — the dashboard reads the **app mirror** `local_skill_path_sessions`, NOT
  `skillpath.skill_path_sessions` (#M36-D2), and the population has no `local_jobsimulation_sessions` mirror, so
  the session takes the skill-path arm of the table's check constraint (#M36-D3).
- **Skillpath completed share** (`skillpath_sessions.go`) — the learning-session seeder marked `completed`
  only on an exact `progress=100` (~1%), starving the learning/my-growth surfaces. M36 makes ~30% complete.

**Distribution coherence.** The growth arc (Growth/Biggest-Improvers) and the over/under-claimer split are
*correlated with the heroes*: the thriving employee is the dashboard's top-performer / succession candidate; the
struggling employee is the at-risk / needs-attention row — the **same** coherence property M35 established,
surfaced now at org scale.

**Closure stays measured.** Every new skill-ref surface draws its node-ids from the **same** taxonomy resolver
the chain uses, so the seed-side closure gene (extended to `membership_skills`) still proves **0 dangling refs**
— the mapped funnel is as closed as the verified chain.

## The AI-readiness showcase org — the 3rd story (v1.10b "fit-up" M51)

M51 adds a **3rd story** curated to demonstrate the **member-AI-readiness flow** end-to-end on the manager
dashboard — the very feature the M201 verify reported as a false-negative (shipped in prod, invisible to the stale
clones; confirmed present at M47). It is a full peer of the M35 two stories, seeded through the same declarative
`stories[]` model and the same closure gate, plus three **net-new AI-readiness seeders** and one **app read-path
demo-patch**.

### The story

A 3rd `stories[]` entry in `presets/stories.seed.yaml` (so the roster is now **3 stories × 3 heroes**, not the
former 2×3): org **Northwind Aviation** (200 members, `narrative: ai-readiness`), with a hero trio —

| Hero | Vantage / trajectory | AI-readiness state |
|---|---|---|
| **Aria** | end-user, thriving | **COMPLETED** — stage 3 (all 3 steps of the onboarding/evaluation done) |
| **Ben** | end-user, struggling | **STARTED** — stage 1 (mid-cycle, in-progress on the onboarding element) |
| **Dana** | manager | views the org **AI-readiness dashboard** (rides the aggregates; seeds no chain of her own) |

The `narrative: ai-readiness` biases a share of members toward AI-named skills via the existing
`isAISkillName`/`filterAISkills` path (`membership_skills.go`) — the same substring semantics the dashboard's
`matchAISkill` uses.

> **✅ CORRECTED (M219, v2.3 "cue to cue") — the `jump_to` targets above were the LEGACY page.** Dana's deep-link
> pointed at `/enterprise/workforce/ai-readiness`, the pre-v3.0 **unlinked orphan**; Aria's and Ben's pointed at
> the generic `/profile`, **which shows no readiness at all**. They now land on **`/ai-readiness`** (Dana — the
> current dashboard) and **`/home`** (Aria + Ben — the member readiness surface has **no route of its own**; it is
> embedded there). A legacy target is now a **hard failure** at seed time (`cockpit.go` `LegacyReadinessPaths` /
> `ValidateCockpitManifest`). See [`../../services/ai-readiness.md` § Surfaces](../../services/ai-readiness.md).

### The three net-new seeders (the AI-readiness chain)

Nothing wrote the `organization_settings` or `ai_readiness_*` tables before M51; three seeders do now, DAG-ordered
`config → funnel` (the funnel's Step-scored signals reference the config's cycle/skills/sims rows — a wrong
`DependsOn` would order them backwards; pinned by `TestAIReadinessSeeders_RegistrationContract`):

| Seeder (`stack-seeding/seeders/`) | Writes | What it lands |
|---|---|---|
| **`OrgSettingsSeeder`** (`org_settings.go`) | `organization_settings` (`setting='ai_readiness', is_enabled=true`, one row per org) | The **enablement gate** the dashboard keys on. (Enablement — this **gate 1**/data layer — is an **org setting**, resolved from the M48 contract, *not* a PostHog flag. The next-web UI *additionally* checks a **gate 2** PostHog flag `flag_ai_readiness`, a separate layer the seeder doesn't write — see [`ai-readiness.md`](../../services/ai-readiness.md#org-enablement-the-gate) for both gates + how the demo satisfies the FE flag.) |
| **`AIReadinessConfigSeeder`** (`ai_readiness_config.go`) | `ai_readiness_cycles` (**×2 since M219 — one `closed` + one `active`**), `ai_readiness_skills` (~core weight-1.0 + enabling 0.5, **real replayed-taxonomy node-ids** via the resolver — never fabricated), `ai_readiness_sims` (×2, **from the content pool's RESERVED TAIL** since M219), `ai_readiness_steps` (×3) | The **cycles + 3-step definition** the funnel scores against. |
| **`AIReadinessFunnelSeeder`** (`ai_readiness_funnel.go`) | **199 frozen `ai_readiness_snapshots`** (one per stage≥1 member, platform-model-scored) + `ai_readiness_user_step_progresses` + the Step-1 `user_skill_evidences` + the Step-2/3 sessions + (**M219**) each interview's `jobsimulation.actors` + `jobsimulation.interactions` turns | The **200-member funnel at 78.4% all-3-complete** (stage-3 = 156, stage-2 = 21, stage-1 = 22), Aria pinned stage 3 / Ben stage 1 / Dana excluded. |

> **M219 rewrote three of this seeder's contracts** — the cycle state (both cycles, because the two vantages need
> opposite ones), the per-member skill count (one mapped skill scored the COMPLETED "Champion" hero **5/30**), and
> the sim-ref reservation (a generic activity session could draw the readiness sim and silently score a member
> against a step they never took). The full contract, with the arithmetic and the RED-proven fences, is in
> [`../../services/ai-readiness.md` § The FILLED-ness contract](../../services/ai-readiness.md).

> **⚠️ A HERO'S ROLE MUST CLASSIFY — or her skills are the taxonomy's alphabetical head (M219).** Both
> AI-readiness heroes shipped with junk profiles for four releases, and the first employee sweep ever run on a
> Northwind seat found it immediately. Two causes, one symptom:
>
> 1. **The curated tier only covered the orgs that existed when it was built.** M42e added the curated
>    skill-name allow-lists precisely to keep the flat pool's `ORDER BY node_id` head (`15Five`, `3dcart`) out of
>    a profile — but it curated exactly **`software`** and **`sales`**. M51 then added Northwind with **"Data
>    Analyst"** and **"Operations Specialist"**, which match *neither*, so both heroes fell through to the flat
>    pool: **Aria — a Data Analyst — claimed "24-hour dietary recall", "2D Animation Software" and "3D
>    Bioprinting in Dentistry".** The classifier's own comment blessed the fall-through as *"no regression for an
>    unclassified role"*. A silent fall-through to the flat pool **is** the regression.
> 2. **"Operations Specialist" is not a public `job_role` at all** (the preset comment claiming it "resolves" was
>    false — the taxonomy has Operations *Analyst*/*Manager*/*Engineer*, never *Specialist*). So Ben had **no
>    role**: no title on `/profile/skills`, no role-core skills, and therefore **even his VERIFIED skills came
>    off the flat head — he was "verified" in `15Five` and `17Track`.** He is now an **Operations Analyst**,
>    which resolves; a non-resolving name must **drop, never be invented**.
>
> M219 adds `data` + `operations` curated categories (hand-picked, taxonomy-verified — the ops family also
> contains real junk for this persona: *"Lean NOx Traps"*, *"Scheduling irrigations"*). **The fence, which is the
> point:** `TestShippedPresets_EveryHeroRoleClassifies` asserts **every role any shipped preset actually seeds
> classifies to a real curated category**, read from the real presets, never a fixture. **Add an org with a new
> role family and it fails at `go test` — not four releases later, in front of a customer.**
>
> ---
>
> ### ⚠️ …AND THAT FENCE PROVED THE WRONG PROPOSITION. THE JUNK SHIPPED ANYWAY (M219 R-8).
>
> The fence above passed. **Aria Holt still shipped claiming `15Five`, `17Track` and `24-hour dietary recall`**,
> and so did **eight ordinary Northwind members** (Ava Park, Amara Fischer, Zara Costa, Hannah Petrov, Theo
> Ferrari, Tom Okafor, Arjun Andersen, Sven Okafor — all claiming *"24-hour dietary recall"*). It was never
> hero-only.
>
> **"The role classifies" is not "the family is big enough."** The draw is filled to
> `want = trajectoryVerifiedCount + claimedTailCount`. Aria's `want` is **28**. Her `data` family shipped **28
> names — of which only 23 RESOLVED** against the live taxonomy (5 were dead: *ETL*, *Exploratory Data Analysis*,
> *KPI Development*, *Data Governance*, *Forecasting*), and **~8 of the survivors deduped** against Data Analyst's
> 10 role-skills. **25 usable tokens for a want of 28** → the last 3 came off the flat pool's `ORDER BY node_id`
> head. **Ben was clean only because his `want` (16) was small enough that his family still covered it** — and
> that asymmetry is the *proof* the defect was pool **SIZE**, not pool **resolution**. (`operations` shipped 5
> dead names too: *Process Optimization Techniques*, *Capacity Planning*, *Root Cause Analysis*, *Vendor
> Management*, *Quality Assurance*. Nobody had ever checked.)
>
> **THE FIX — the flat tier is DELETED, not demoted.** The ladder is now:
>
> ```
>   role's job_role_skills  →  the role's CURATED family  →  the CURATED GENERAL family  →  STOP
> ```
>
> `curatedGeneralSkills` (33 transferable-professional names — *Communication*, *Written Communication*,
> *Problem Solving*, *Microsoft Excel*, *Prompt Engineering* …, every one verified to resolve) is the coherent
> last tier: it fills an unclassified role's draw **without lying**, so the flat pool could be removed outright
> rather than merely pushed further down. **If all three tiers are exhausted the seeder draws FEWER skills.**
> Honest degradation is the contract; padding is not. `data` grew 28 → **50** and `operations` 30 → **45**, all
> names re-verified against the live taxonomy.
>
> **THE TWO FENCES** (both RED-proven against the pre-R-8 ladder, green after —
> `seeders/curated_ladder_m219r8_test.go`):
> 1. **`TestSeededMembers_NeverDrawFromFlatPool`** — the structural DoD. The flat pool is **poisoned** with
>    sentinel tokens and every draw path (`combinedNamedPool`, `resolveHeroSkills`, `skillsForRole`) is run at
>    every shipped persona's real `want`. One poison token anywhere is a failure. It models the two forces that
>    actually shrank the pool — **1-in-6 name attrition** and **role/curated node-id overlap** — because a first
>    cut *without* them **passed against the broken ladder**, which is the same theatre it was written to end.
> 2. **`TestCuratedLadder_CoversLargestWant`** — the SIZE property the old fence never proved: every family must
>    carry `want + curatedAttritionMargin` (15). RED-proving it also surfaced a **second** under-margin family
>    nobody had noticed: **`sales`** (33 vs the 41 Sara needs).
>
> **The rule this encodes, stated once:** *a fence that proves the wrong proposition is the bug.* "The role
> classifies" ≠ "the pool is big enough". "It resolves" ≠ "it has skills" (see the role-less-hero box below).
> "It serves" ≠ "it renders".
>
> ### ⚠️ A HERO'S ROLE MUST *CARRY ROLE-SKILLS*, NOT MERELY EXIST (M219 R-8 — the role-less hero)
>
> Ben rendered with **no role title at all** on `/profile/skills`, while Aria's rendered fine. The repoint to
> **"Operations Analyst"** (above) was *not* the fix it looked like: that job_role **exists** (`J-OPERAT-3566`)
> and carries **ZERO `job_role_skills`**.
>
> The seeder's own resolver has *always* required role-skills — `readJobRolePool` / `readJobRoleByName` both
> demand `EXISTS(job_role_skills)` (*"a role with none isn't a believable hero role"*). So the resolver
> **rejected** the role, `job_role_id` landed **NULL**, and `public.user_basic_info.job_role_id` — which is what
> `/profile` + `/profile/skills` read the role title from, and what the `jobRoleMatch` role-gap widgets key off —
> had nothing to read. **The seeder then took its silent-degradation path and seeded happily.** A preset comment
> was even written asserting *"Operations Analyst DOES resolve"*. It resolves as a row; it does not resolve as a
> **hero role**.
>
> **Fixed two ways:** Ben is now a **Business Operations Analyst** (`J-BUSOPE-38C4`, 10 role-skills, still
> classifies to the `operations` family — persona unchanged), in `stories.seed.yaml`,
> `seed-generation-manifest.yaml` **and** `playthroughs/seed/pt-world.seed.yaml`; and
> **`assertHeroRolesResolve`** (wired into `UsersSeeder`) now makes a hero role the resolver rejects a **HARD
> SEED FAILURE** naming the hero, the role, and what the user would have seen. The no-fabrication path is
> preserved exactly where it is legitimate: with **no replayed taxonomy** there is nothing to resolve against, so
> the fence is a no-op and the hero keeps her declared label.
>
> ### And the allow-list alone fixes NOTHING — the FALLBACK LADDER is the load-bearing part
>
> The curated `operations` pool shipped **fully populated and completely unused**. On the first cold reseed the
> Data Analyst hero came out clean while the Operations Analyst hero was *still* "verified" in `15Five`.
>
> `skillsForRole`'s ladder was **role → FLAT**. A public `job_role` can **exist and carry zero
> `job_role_skills`** — `Operations Analyst` is exactly that — so `byRole` is empty and the function returned the
> flat `ORDER BY node_id` head. And `combinedNamedPool` draws its **role tier from that same function**, so
> **tier 1 was already the junk** and filled the whole quota before the curated tier was ever consulted.
>
> **R-5b made the ladder `role → CURATED → flat`** — flat genuinely *last*, rather than first. Both twins
> (`namedSkillRefs`, `taxonomyRefs`) carried the bug — which is why even the hero's **verified** chain certified
> him in a junk skill.
>
> > ⚠️ **SUPERSEDED — and the supersession is the point. That ladder STILL SHIPPED THE JUNK.** Demoting flat was
> > not enough: it **fired whenever the curated family ran DRY before `want`**. Aria's `want` was 28; her `data`
> > family shipped 28 names of which only **23 resolved** (5 dead) and **~8 deduped** against her role's 10
> > role-skills → **25 usable** → the last 3 came off the flat head. Ben was clean *only* because his `want` (16)
> > was covered — **that asymmetry is the proof the defect was pool SIZE, not pool ORDER.**
> >
> > **R-8 DELETES the flat tier.** The ladder is `role → curated → **general** → **STOP**` (see the R-8 section
> > above). An exhausted ladder yields **FEWER** skills, never padded ones — honest degradation is the contract;
> > padding is not. *"It classifies" ≠ "it is big enough."*
>
> ⚠️ **Two unit tests were green throughout**: they proved the curated pool *resolved*. Neither proved anything
> ever *read* it. **A test that proves a thing exists is not a test that proves it is used** — only a cold
> reset-to-seed, and looking at what actually came out of the database, exposed this.

### Why closed-cycle + frozen snapshots (the strategy M51 shipped)

The M48 contract offers two seed strategies (see [`../../services/ai-readiness.md`](../../services/ai-readiness.md)):
an **active** cycle (the dashboard live-recomputes from signals) or a **closed** cycle (the dashboard reads
pre-computed `ai_readiness_snapshots` directly). M51's iters 03→06 built and then **falsified** the active-signals
path — the live-recompute never completes in the coverage harness's budget (a per-skill federated translation N+1,
the M46 per-object-RPC class). The milestone **shipped the closed-cycle / frozen-snapshot** strategy: the cycle is
seeded `closed`, one frozen snapshot per member carries the platform-model score, so the dashboard reads finished
data. (`ai_readiness_*` were also added to `stackseed --reset` + a baked `--reload-sentinel` after seed, so the
showcase re-seeds cleanly.)

### The `app-aireadiness-snapshot-loadmembers` demo-patch (the frozen-read perf bound)

Freezing the *scores* was not enough: the frozen read path (`app buildResponseFromSnapshots`) still calls
`loadMembers(orgID, "")` — an **unbounded whole-org member hydration** to re-join current name/role/tags onto each
snapshot — so even a direct `?cycle=<closed>` GET timed out at 180 s at 200 members (**"frozen" froze the SCORES,
not the RESPONSE** — iter-08's root cause). The fix is a **pure perf** app injection demo-patch
(`patches/app-aireadiness-snapshot-loadmembers.yaml` + `apply-app-aireadiness-loadmembers.sh`, wired non-fatally
into `up-injected.sh`, following the M46 `app-targetrole-authz-skip` precedent): it swaps the unbounded
`loadMembers` for the existing bounded sibling `loadMembersByUserIDs` over the ~199 snapshot user-ids. **Data-
identical** (the members map is keyed and looked up by snapshot `UserID`; members with no matching snapshot were
loaded-but-never-used), scoped to the snapshot read path only. The frozen GET went **180 s-timeout → 19 ms**; the
dashboard renders the full funnel. This is a disclosed demo-perf relaxation, not a prod fix — **prod's frozen read
still hydrates the whole org and would need `loadMembers` bounded / a `frozen_tags` column (M314b)**, documented in
[`coverage-protocol.md`](coverage-protocol.md).

### The gate

Proven by the **M42 semantic coverage gate, manager vantage**, on Northwind: `(failingSections, escapes) = (0, 0)`,
persona green, frontier-exhausted, on a fresh `/demo-up` (rext `fit-up-m51`). Closure stays measured — the
AI-readiness skills resolve through the same taxonomy resolver, so the seed-side closure gene proves 0 dangling
refs across all 3 orgs.

## The hiring org — the 4th story (v2.4 "casting call" M222 gate / M223 seed)

The 4th story is a **genuine recruiting org** (`is_hiring=true`) that demonstrates the **recruiter
candidate-comparison scoreboard** — `/enterprise/activity-dashboard → AI-Simulations → [simId]`, one row per
candidate who took a hiring simulation, ranked by a comparable score. Unlike the three Workforce orgs
(admin/member), it is **~5 `admin` + 45 `candidate`, no `member`**; its "positions" are **5 real captured
`SIMULATION_TYPE_HIRING` sims** (real content, zero synth — **no `directus.job_position` replay**, per M222 BA-6).

**M222 "read the room" lands the GATE only**: the blueprint's `StoryOrg.IsHiring` (yaml
`is_hiring`) field + the **`narrative: hiring`** discriminator (unified through `ResolvedStory.IsHiringOrg()`), the
`OrgSeeder` thread into `public.organizations.is_hiring` (was hardcoded `false`), and the reserved deterministic
`blueprint.HiringOrgID()` (`= StoryOrgID("hiring")`). **M223 "casting the ensemble" (this milestone)** adds the
4th story + the two seeders below. **M224 "the callback"** adds the cockpit hero seat (1 recruiter + 2
candidates) + the Clerkenstein `publicMetadata.isHiring` wiring (the browser re-skin's client-side half of the
dual-write).

> **The render-gate trap (why the read-model is documented before the seeder).** The scoreboard's score renders
> from the MIRROR table **`app.public.local_jobsimulation_sessions.score`**, **NOT** `jobsimulation.sessions.score`.
> Seeding only `jobsimulation.sessions` yields a scoreboard that renders its chrome with every score blank — the
> same **render-gate-bypasses-the-seed** class M219 hit with AI-readiness. Full read-model + write-set contract:
> [`../../services/hiring.md`](../../services/hiring.md).

### The M223 hiring chain — two seeders (`hiring-config` + `hiring-funnel`)

The 4th story is **heroless at M223** (`stories.seed.yaml` `id: hiring`, Meridian Talent, size 50, `role_mix
0.1 admin / 0.9 candidate` → 5 admins + 45 candidates); its cockpit trio materializes at M224. Two net-new
seeders build the recruiter surface (both DAG level 2, `depends=[org users content]`):

- **`HiringConfigSeeder`** — resolves the org's **5 shared positions** and writes them as
  `organization_sim_invitation_links` (one per (org, sim); the table's `UNIQUE(simulation_id,
  organization_id)` makes it a bounded 5/org, no balloon — folding in the "optional invitation links" scope).
  The positions are **5 REAL captured `SIMULATION_TYPE_HIRING` sims**, resolved by a **type-aware reader**
  `readHiringSimPool` (`directus.simulations WHERE type='SIMULATION_TYPE_HIRING'` — the generic `contentref`
  pool is type-blind; the `readAIReadinessSkillPool` precedent). The shared `hiringSimRefs` takes the first 5
  and is used by **both** seeders so positions and scored sessions **co-derive**. The 5 are **withheld from the
  generic sims pool** (`withoutIDs`, before the reserve tail) so no generic activity session can collide with a
  position (the M219 R-3 disjoint reservation). No `directus.job_position` replay — 0 rows captured, unread by
  the scoreboard (M222 D4).
- **`HiringFunnelSeeder`** — the candidate-assessment funnel. **CRITICAL: it writes the MIRROR pair like the
  `PersonaSeeder` does (`persona.go` — `jobsimulation.sessions` + the `public.local_jobsimulation_sessions`
  mirror), NOT the `AIReadinessFunnelSeeder` shape** (which deliberately skips the mirror and scores off frozen
  `ai_readiness_snapshots`). The recruiter scoreboard reads the score from the mirror, so skipping it renders an
  EMPTY comparison — a test **RED-proves** the fence. Per (candidate × position) it co-writes the pair,
  org-scoped, `sim_type='SIMULATION_TYPE_HIRING'`, G14-valid enums. **The funnel shape** (measured on the
  preset): **~90% of candidates ASSESSED on EXACTLY ONE position** — the role applied for (**M227 fix #3**; before
  M227 every candidate took all 5) — round-robined evenly across the 5 so each ranks **~8 candidates** (43 assessed
  of 45 → min 8 / max 9 per position), the rest **assigned-not-taken** (no scored sessions — the 2nd candidate
  hero's future state, M224). Each candidate has a base **aptitude ∈ [30,95]** + jitter, so within a position the ~8
  candidates are **RANKABLE** (NOT identical, the M219 anti-flat-arc lesson); `passed` when score ≥ 60 else
  `failed`. **The compare gate retuned `≥40 → ≥6`** (a small margin below the seeded min of ~8). Only `role=candidate` members audition; the funnel writes **0 skill refs** (closure green
  trivially). The 5 admins inherit `org:feature:insights` from the **global `p3` admin Casbin policy** via their
  standard `admin` g2 grant — no net-new grant (M223 D1). New reset surface:
  `organization_sim_invitation_links` (child-first); the session pair reuses the already-reset
  `jobsimulation.sessions` + `local_jobsimulation_sessions`.

## The presenter cockpit (M38)

> **The cockpit UX is now specced standalone in [`cockpit-spec.md`](cockpit-spec.md) (v1.10 "method acting"
> M43).** That doc is the canonical reference for the panel's UI surface + deep-link contract; the **v1.10 M43
> UX pass superseded this section's two-button model** — there is now **one** unified **[Log in as]** CTA per
> hero (it logs in *and* lands on her per-role `jump_to`), a light professional restyle, FontAwesome icons, a
> seed-manifest download, and a staged login-progress overlay. This section is kept as the M37/M38 *producer/
> consumer* origin (the roster-export + handshake seam below); read it with `cockpit-spec.md` for the current
> UX.

The seeded world (M34–M36) + the Clerkenstein multi-identity seat-switch (M37) make the *individual* surfaces
real; **M38** makes the whole Stories & Heroes engine **clickable**. The **presenter cockpit** is a standalone
served panel that lists each story → its hero trio and, per hero, an action that logs the demo-giver in as her
and lands her on the right screen — so a presenter picks a hero and presents that part of the story live. _(M38
shipped two actions per hero — `[Login as]` → app root + `[Jump to section]` → the hero's deep-link; v1.10 M43
unified them into one `[Log in as]` → `jump_to` — see [`cockpit-spec.md`](cockpit-spec.md).)_

### For PMs — the demo-driving surface

A demo flows: *show Maya's verified-skill profile → show Tom's stark claimed-vs-verified gap → log in as their
manager Dan and watch the same two people become the standout high/low rows of his Workforce dashboard.* The
cockpit is the remote control for that flow: one click logs in as the chosen hero and jumps to her screen. No
typing a login, no hunting for the right URL — the story is a menu.

### The shape

```
Presenter Cockpit — demo-3
  Story: AI Transformation & Reskilling   (Cervato Systems · 220 people)
    ▸ Maya Chen — Backend Developer · EMPLOYEE · THRIVING
        "8 verified skills, rising growth arc, mobility-ready"        [Log in as]   (→ her /profile)
    ▸ Tom Becker — Backend Developer · EMPLOYEE · STRUGGLING
        "Few/low verified skills, OVER-rates himself (stark gap)"     [Log in as]   (→ his /profile)
    ▸ Dan Rossi — Engineering Manager · MANAGER
        "Team gaps, role-readiness, succession (Maya), at-risk (Tom)" [Log in as]   (→ Workforce · Skills Verification)
  Story: SDR Onboarding & Ramp   (Solvantis · 120 people)
    ▸ Sara Whitfield · EMPLOYEE·THRIVING  /  Nick Alvarez · EMPLOYEE·STRUGGLING  /  Leah Donovan · MANAGER

(v1.10 M43: ONE [Log in as] CTA per hero, routed to her per-role jump_to; the M38 [Jump to section] button is gone.)
```

(Display-label note: `vantage: end-user` renders as **EMPLOYEE**, `manager` as **MANAGER** — the YAML attribute
value stays `end-user | manager`; "employee" is a label, not an enum value.)

### For engineers — how it works

**Standalone served panel, never an in-app overlay (D15).** The cockpit is a host-native HTTP server
(`rosetta-extensions/demo-stack/cockpit.py`, stdlib only) on an **offset port** (`7700 + N·10000`), brought up
with the stack and torn down with it — it is **never** an edit to next-web. This preserves the hard
zero-platform-repo-edit line: the cockpit reaches the platform only as a browser would (over the FAPI
handshake), never by modifying it.

**Single source — it reads the same `stack.stories.yaml` that seeded the data (D9).** The cockpit menu is a
**manifest** the seeder projects from the very file that seeded the heroes (`stackseed --cockpit-export`), so
the annotations describing a hero in the cockpit are the same ones that scoped her seed — the menu can never
drift from the data. (The demo tooling is stdlib-only Python, so the YAML is parsed once on the Go side and the
panel reads the derived JSON — single-source preserved.)

**The CTA = one FAPI handshake redirect.** The action points the browser at the multi-identity fake FAPI's
handshake with the hero's seat-switch key _(M38 rendered two — `[Login as]` + `[Jump to section]`; v1.10 M43
unified to one `[Log in as]` → `jump_to`, see [`cockpit-spec.md`](cockpit-spec.md))_:

```
https://<fapi-host>/v1/client/handshake?__clerk_identity=<hero-key>&redirect_url=<jump_to>
```

`<fapi-host>` is the per-stack fake-FAPI on its own **offset port** `127.0.0.1:<5400 + N·10000>` (e.g.
`127.0.0.1:35400` for `demo-3`), served over HTTPS (clerk-js requires it); `redirect_url` is an absolute
next-web URL on the app's offset port `<3000 + N·10000>`. The FAPI selects the chosen hero's seat from
`__clerk_identity` **then** establishes the session and redirects to `redirect_url` — so the hero is the
active identity *everywhere* (the client view, `/v1/me`, the minted token, the cookies) AND the browser lands
on her screen, in one move. The unified **[Log in as]** (v1.10 M43) lands the hero on her `jump_to` per-role
screen. The key is the hero's `stories.yaml` id — the **same** key the roster export gave
Clerkenstein's registry, so the seat always resolves. (The handshake + multi-identity selection are M37; see
[`clerkenstein/knowledge/architecture.md` § Multi-identity].)

**The roster-export producer — the M37 integration seam.** M37 shipped the *consumer* (the fake FAPI loads a
`[]DemoUser` roster from `FAKE_FAPI_ROSTER`); M38 ships the *producer*. `stackseed --roster-export` derives the
roster JSON — each hero's **exact** clerk claims (`auth_id`/`eid`/`email`/`org_*`/`org_role`) — **single-sourced
from the seeder's own id-derivation** (so "login as Maya" authenticates the real seeded user). Clerkenstein is a
separate Go module and never imports the seeder; the seeder/demo-tooling exports, Clerkenstein consumes. The
demo bring-up sets `FAKE_FAPI_ROSTER` on the `demo-N-fake-fapi` container so its FAPI is multi-identity. The
`org_role` claim is **vantage-faithful** — a hero's exported role follows her seat (manager → `admin`,
end-user → `member`), single-sourced through one `roleForHero` helper that the seeder also writes the
`membership` row + the casbin `g2` grant with, so the three writes agree per hero (an "employee" demo seat
reads as `member` in her JWT, not org-admin) (#M38-D8).

**The deep-link catalog (O9).** The cockpit ships an enumerated, stable set of next-web routes per vantage — the
*individual* surfaces an end-user hero demos (`/profile`, Skill Spotlight, my-growth, take-a-sim) and the
*org-intelligence* surfaces a manager hero demos (the Workforce dashboard tabs — verification / role-readiness /
succession / mobility — plus the talent pool). A hero's `jump_to` is matched against this catalog for its button
label; an unrecognized `jump_to` still works (it's a raw path) with a generic label.

### Bring it up

```bash
# A storytelling demo: DEMO_STORIES=1 seeds the 4-org world (3 Workforce/AI-readiness hero trios + the M223
# hiring org), wires the multi-identity fake-fapi, and
# serves the cockpit. Default-off keeps every existing demo byte-identical (structural seed, single-identity).
DEMO_STORIES=1 /demo-up 3
# → the cockpit serves on http://localhost:37700 (7700 + 3·10000). Pick a hero → [Log in as] → her per-role screen.
```

`DEMO_NO_COCKPIT=1` brings up the stories demo without the panel (e.g. an API-only run); `DEMO_STORIES_PRESET`
overrides the preset (default `presets/stories.seed.yaml`). The cockpit + roster + seed all pin `--stack demo-N`,
so the exported ids and the seeded rows are guaranteed to match.

## Running it

```bash
# 1. Replay the public reference library FIRST (a --local-content demo does this automatically, cache-first):
stacksnap replay --surface taxonomy --stack demo-N   # the skill node-ids the chain draws from
stacksnap replay --surface directus --stack demo-N   # the sim templates the sessions link to

# 2. Seed the world. The M34 vertical slice (one hero):
stackseed --stack demo-N --seed presets/stories-maya.seed.yaml
#    …or the full multi-org roster (4 orgs: M35's two Workforce trios + the M51 AI-readiness org + the M223 hiring org):
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
