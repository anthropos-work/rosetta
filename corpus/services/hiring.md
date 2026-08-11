# Hiring (recruiting org-type + the candidate-comparison read-model) — service documentation

> **Status:** documented 2026-07-15 (v2.4 "casting call" **M222 "read the room"** — the BLIND-AREA doc). Authored
> from a **live render-probe** on the v2.3 `billion` demo substrate: every read-path claim below was traced through
> the running dockerized `apps/web` + the `app`/`jobsimulation` code and **reproduced end-to-end**, not inferred.
> This doc is the contract the **M223** candidate seeder + **M224** Clerkenstein `publicMetadata` wiring build
> against. Before M222, hiring existed in the corpus only as a "distinct-frontend" line in
> [`next-web-app.md`](next-web-app.md) (Workforce `apps/web` vs Hiring `apps/hiring`) + the business KB — the
> **read-model that fills the recruiter comparison view had no code anchor**, which is exactly the gap M222 closed.

> **Why a blind area is the worst kind of gap.** The comparison surface renders from a table nobody had named. A
> seeder that wrote the "obvious" table (`jobsimulation.sessions`) would produce a page that renders its chrome and
> its columns with **every score blank** — a green coverage sweep over an empty scoreboard. This is the same
> **render-gate-bypasses-the-seed** class M219 hit with AI-readiness. The headline of this doc is the ONE table that
> actually feeds the score. Read § *The comparison read-model* before seeding anything hiring-shaped.

> **⚠️ RE-GROUNDED — v2.8 M257x iter-23, against platform origin `2adcf71` / `app` @ `5ba17044`.**
> *(The platform-side citations below were re-anchored again to platform `0dab54d` in the M257x sweep —
> `d11a403` had removed the cms / jobsimulation / roadrunner compose services and `repos.yml` entries.)*
> **`5ba17044` is the historical iter-23 re-grounding ref — NOT a governing pin over the anchors below.**
> The `app`-side anchors have been re-derived repeatedly since (M257x iter-49, -52, -98, -100, -102); every
> one re-derived at **iter-102** is measured at `app` **`ad9f3c49`** (= `origin/main` **and** the demo build
> pin `stack-demo/clones.pin.json`, 2026-08-06). They are **not** interchangeable:
> `internal/organization/manager.go:450` / `:453` / `:537` resolve at `ad9f3c49`, and **the offset back to
> `5ba17044` is NOT uniform — do not apply one delta to all three.** Measured line-for-line:
> `:450 → :448` (−2), `:453 → :451` (−2), but **`:537 → :487` (−50)**. This banner said *"each off by −2"*
> until M257x iter-108; on that rule `:537` would resolve to `:535` at `5ba17044`, which is unrelated code. **Read the ref that travels with the anchor; do not read this banner as a pin over the
> whole document.**
> **This doc named a table the platform has since DROPPED — which is the worst possible version of the warning
> directly above.** The score source was `public.local_jobsimulation_sessions`, a `Float32` MIRROR. `app`
> migration `20260729133514.sql:58-62` — *"5. Drop the mirrors."* — **re-points the *referencing* rows** (the
> assignment-session link ids) and then `DROP TABLE "local_jobsimulation_sessions"`. **There is no back-fill:**
> `SET "score"` has **0 hits across the entire migration set**, so no score was copied from the mirror to the
> canonical row — the canonical row already carried it. (This paragraph said "back-fills" until M257x iter-49.)
> Everything below is re-pointed; the three facts that changed:
> 1. **Score source → `public.job_simulation_sessions.score`**, read by
>    `app/internal/organization/intelligence.go:1700` (`m.ent.JobSimulationSession.Query()`). There is no
>    mirror/canonical **pair** any more, so the write-set is **one** session row, not two.
> 2. **Everything `app` writes is in `public`** — which is *not* the same claim as "the `jobsimulation` schema
>    is gone." `20260722104506.sql:79` is `DROP TABLE "sessions"` executed under `search_path=public`, so what
>    it dropped is **`public.sessions`**, replaced by `public.job_simulation_sessions` (`:2`). **No `app`
>    migration touches the `jobsimulation` schema at all**, and **in production** that schema **survives,
>    frozen, until M710** — while **no local dev/demo stack creates it at all** (at platform `0dab54d`
>    `platform/repos.yml` carried **6** entries — **4** since `838d907` — and `jobsimulation` is **not one of them** — `d11a403`
>    deleted it; `app` (`repos.yml:14-17`) is the only `migrations: true` entry. Qualified M257x iter-52,
>    re-anchored in the M257x sweep)
>    — as the twins [`service_taxonomy.md`](../architecture/service_taxonomy.md)'s Tier-1 **Database**
>    characteristic bullet (**named, not pinned** — it carried a line number until M257x iter-120, by then
>    that file's *Tier 1: Core Backend Services* **heading**, not the bullet; find it by its own words,
>    *"one schema, `public`, owned by `app`, which is the only repo with
>    migrations"*, the `cms` / `jobsimulation` / `skillpath` schemas *"legacy husks"*) and
>    [`dependency_map.md:78`](../architecture/dependency_map.md) (§ *2. Job Simulation* — session state goes
>    *"directly to the **`public`** schema (the legacy `jobsimulation` schema is non-authoritative)"*)
>    already said. (This carried a line pin into that doc until M257x iter-102 — by then the closing line of
>    the `:44-52` blockquote correcting the Studio-Desk → Backend → Studio-Room **generation edge**, which
>    says nothing about schemas, `jobsimulation`, or migrations. The bullet is named as well as numbered
>    because `service_taxonomy.md` is edited concurrently and `:62` can move.) `app/atlas.hcl:8` pins
>    `search_path=public`, and the only `CREATE SCHEMA` in the entire migration set is `auth`.
>    (This bullet read *"`jobsimulation.sessions` was dropped"* from iter-23 until M257x iter-49.)
> 3. **One subgraph.** There is no second subgraph for `Session!` to resolve from — no join key, no
>    NULL-bubble hazard.

## Role & Responsibility

**Hiring** is a **sold product** (`hiring.anthropos.work`, the `@anthropos/hiring-app` at `apps/hiring`) AND an
**org-type** that re-skins the **Workforce** app (`apps/web`) for a recruiting buyer. The demo cares about the
**org-type**, not the standalone Hiring app: an `is_hiring` organization runs its members through
**`SIMULATION_TYPE_HIRING`** job simulations and reads a **candidate-comparison scoreboard** — the recruiter's core
value: *"line up every candidate who took this hiring simulation, ranked by score."* That scoreboard renders **from
seedable data alone**, so the demo can show it **without a platform edit** (M222 D1, the release go/no-go) — but
**not in `apps/web`**. M222's dockerized-`apps/web` (`/enterprise/activity-dashboard`) premise was **falsified at
M224**: a user whose Clerk memberships are **all** hiring orgs is **ejected out of `apps/web`** to the standalone
Hiring product (`apps/web/src/context/UserStatusContext.tsx:141-173` — `userHasAllHiringOrgs` over
`publicMetadata.isHiring`, then `buildSwitchHandoffUrl({ targetProduct: 'hiring' })`), so *"the org genuinely reads
as hiring"* and *"the scoreboard is reachable in `apps/web`"* are **mutually exclusive**. The demo therefore runs the
untouched **`apps/hiring`** as a second UI container — still zero platform edits. See § *The render path*. (This
paragraph asserted the scoreboard *"lives in the dockerized `apps/web` … and survives the `is_hiring` flip"* from
M222 until M257x, while three later sections already said the opposite.)

There is **no `hiring` microservice**. The feature is a composition: an org-type flag on `app`, the
jobsimulation runtime (**itself now a domain inside `app`**, `app/internal/jobsimulation/` — no container
since `d11a403`), an `app`-side read-model (`IntelligenceManager`), the single `backend` GraphQL subgraph,
and an `apps/web` surface gated client-side on a Clerk org flag.

## The org-type gate — `is_hiring` is a DUAL-WRITE

`is_hiring` must be set in **TWO** places, because the platform derives the org's type differently on each side:

1. **Backend — `public.organizations.is_hiring boolean NOT NULL default false`.** The server-side org-type. The
   seeder writes it directly (M222 landed the gate — see § *The seeder-output contract*).
   **The `resolver_queries.go` insights path does NOT read it at all** — `InsightsJobSimulationByMemberships`
   (`:1034-1080`) gates on exactly two things, the `OrgFeatureInsights` Casbin permission (`:1035`) and
   membership status ∈ {active, invited} (`:1053`); `grep -in hiring` over that resolver and over
   `internal/organization/intelligence.go` returns only sim-TYPE filters and nothing, respectively (positive
   controls: `OrgFeatureInsights` ×8, `JobSimulationSession` ×44). **But "no read path reads it" would be a second false claim:** the CONTENT-LIBRARY read path does —
   `PrivateJobSimulations` branches its result set on `GetOrganizationIsHiring`
   (`resolver_cms_queries.go:95,210,258,295` — `isHiring` picks `hiringLibraryTypes()` over
   `workforceLibraryTypes()` at `:99-103`), as do `organization/manager.go:450` (a forced Clerk membership
   is created with role `candidate` instead of `member`, `:453`) and `:537` + `siminvitationlink.go:63` (both
   **hard-error `"organization is not hiring"`** — the latter is `CreateOrganizationSimInvitationLink`. Note
   the `HiringConfigSeeder` does **not** go through that RPC: it writes the 5 positions straight into
   `public.organization_sim_invitation_links` with `CopyRowsIdempotent` (`hiring_config.go:99`), so this
   hard-error never reaches it. This passage claimed the opposite until M257x iter-52). And the client re-skin is **not**
   driven by this column either: it is read from Clerk `publicMetadata.isHiring`
   (`useGetClerkOrganization.tsx:20`, quoted below). So the column gates the content library and the
   org-type surfaces; the *insights* scoreboard is indifferent to it.
2. **Client — Clerk `publicMetadata.isHiring = true`.** The **entire `apps/web` re-skin is derived client-side from
   Clerk, never from a GraphQL call:**

   ```ts
   // apps/web/src/hooks/useGetClerkOrganization.tsx:20-21
   const isHiringOrg    = Boolean(organization?.publicMetadata?.isHiring);
   const organizationId = organization?.publicMetadata?.eid as string;
   ```

   So a demo org whose DB row says `is_hiring=true` but whose Clerk metadata omits `isHiring` renders as a **normal
   Workforce org** in the browser — the nav never relabels, the "Results" framing never appears. **⇒ M224 must
   extend Clerkenstein's fake Clerk API to emit `publicMetadata.isHiring = true`** for the hiring org. This is a
   rext change (the mock), **NOT a platform edit**.

   > **The browser-visible emission is the FAPI, not the BAPI (M224 KB-fidelity correction).** Clerkenstein emits org
   > `public_metadata.{eid}` **independently on both sides**, but the one the client re-skin above reads (`@clerk/clerk-js`
   > `useOrganization().publicMetadata`) is the **fake FAPI**: `clerkenstein/clerk-frontend/resources.go` `orgMemberships()`
   > builds `PublicMetadata:{eid}`, fed by the `RosterEntry`→`DemoUser` roster thread (the M39 `org_name`/`org_slug`
   > precedent). **So `isHiring` slots into the FAPI roster+resource path** (`clerk-frontend/registry.go` `RosterEntry` +
   > `clerk-frontend/resources.go` `orgMemberships()`), which trips the **BLOCKING `/align-run`** clerk-frontend guard. The
   > server-side BAPI (`clerk-backend/resources.go` `organizationWithEid`) emits its own `{eid}` copy but is **not** what the
   > re-skin reads (the server derives hiring from the `public.organizations.is_hiring` DB column) — a BAPI `isHiring` extension
   > is optional, only if a server-side consumer reads `organization.publicMetadata.isHiring`.

> **Both, or the demo is half-lit — and the two halves fail in opposite directions.** Write each one down
> separately; a single "it doesn't re-skin" covers neither.
>
> - **DB-only** (column `true`, Clerk metadata absent) → **the client half is dead.** `isHiringOrg` is `false`,
>   so the nav keeps the "Activity dashboard" label (`useNavbarSections.tsx:476` @ `next-web-app` `8297c684`;
>   it stood sixteen lines earlier at `bb3313bc`) and the org is *not* filtered
>   out of the workforce list (`useGetClerkOrganization.tsx:16-18`). And the product-boundary hand-off — which
>   reads Clerk and **only** Clerk (`apps/web/src/context/UserStatusContext.tsx:144-149` computes
>   `userHasAllHiringOrgs` from `publicMetadata.isHiring`, then `:168-172` fires
>   `buildSwitchHandoffUrl({targetProduct:'hiring'})`) — **never fires**, so the recruiter is never handed to
>   `apps/hiring`; she sits in a Workforce-skinned `apps/web`. Point the cockpit at the hiring base anyway and
>   the *symmetric* guard bounces her straight back (`apps/hiring/src/context/UserStatusContext.tsx:125,144-145`
>   → `targetProduct:'workforce'`). This is exactly the `billion` spike that produced M222's false `apps/web`
>   premise (§ *The render path*). The server half, meanwhile, is entirely correct.
> - **Clerk-only** (metadata `true`, column `false`) → **the client half is fine and the server half is dead.**
>   The browser *does* re-skin (the re-skin reads Clerk, not the column) and the hand-off *does* route the
>   recruiter to `apps/hiring`. What breaks is server-side: the content library serves the **workforce**
>   type-set instead of the hiring one (`resolver_cms_queries.go:99-103`), and
>   `CreateOrganizationSimInvitationLink` hard-errors `"organization is not hiring"` (`siminvitationlink.go:63`,
>   guarded at `:62`)
>   for any caller that uses it. **The `HiringConfigSeeder` is not such a caller** — it writes the 5 positions
>   directly (`hiring_config.go:99`) and is unaffected. This bullet previously said it *"cannot write the 5
>   positions in the first place"*; that consequence is **refuted** (M257x iter-52).
>
> **Neither half, however, gates the insights scoreboard** — the text here used to say Clerk-only meant *"the
> insights read-path won't treat the cohort as hiring"*, and that sent every empty-scoreboard debug to the wrong
> place. What actually gates it: the `OrgFeatureInsights` Casbin grant, membership status ∈ {active, invited},
> and the presence of `public.job_simulation_sessions` rows. The seeder writes #1; the mock emits #2; M224 wires
> the pair.

## The `candidate` membership role

A hiring org's population is **admins + candidates**, not the Workforce **admin/member** shape. The blueprint's
`RoleMix` already carries a `Candidate` ratio (`blueprint/blueprint.go` `RoleMix{Admin, Member, Candidate,
AdminEmails}`), so no new role primitive is needed — M223's hiring story simply sets `role_mix ≈ 0.1 admin / 0.9
candidate` (no `member`). A candidate is a normal `public.memberships` row (`GetMemberships` requires status
`active`/`invited`) whose org is `is_hiring`; the comparison scoreboard joins sessions → memberships to hydrate each
candidate's name/role, so **every candidate the scoreboard shows must have an active membership**.

## Hiring simulations — `SIMULATION_TYPE_HIRING` and the (optional, absent) `job_position`

- **Hiring sims are `SIMULATION_TYPE_HIRING`-typed job simulations** — the same `jobsimulation` runtime the Workforce
  product uses, tagged as a hiring assessment. The captured public snapshot carries **87 real
  `SIMULATION_TYPE_HIRING` sims** (published + public) — a rich pool; M223 picks **5** as the org's "positions"
  (real content, zero synth — M222 BA-6).
- **`JobSimulation.jobPosition` is OPTIONAL and unread by the scoreboard.** The `directus.job_position` entity models
  a "role you're hiring for", but the comparison surface does **not** read it, and the captured snapshot has **0
  `job_position` rows** (the prod "443" was never captured). **⇒ the 5 "positions" ARE 5 real HIRING sims; there is
  no `job_position` replay** (M222 D4 → M223 Scope.In refinement). A candidate is comparable to another when they
  share the same **`sim_id`** — the sim IS the position for scoreboard purposes.
  ⚠️ **This said `jobsimulation_id` until M257x iter-102, and no such column exists on the live table.**
  `public.job_simulation_sessions` has 23 columns; the sim reference is **`sim_id`**
  (`20260722104506.sql:7`, `ent/schema/job_simulation_session.go:33`). `jobsimulation_id` was a column of the
  **dropped mirror** `local_jobsimulation_sessions` — created `20240527131926.sql:7`, dropped with the table
  at `20260729133514.sql:62`, the same drop this doc's re-grounding banner describes. A seeder or query
  written from the old name does not mis-rank; it **errors on an unknown column**. Measured at `app`
  `ad9f3c49` (= `origin/main` and the demo build pin, 2026-08-06).

## The comparison read-model (THE HEADLINE) — the score is `public.job_simulation_sessions.score`

The recruiter's scoreboard is `/enterprise/activity-dashboard → AI-Simulations → [simId]`: one row per candidate who
took the sim, ranked by a comparable **score**. That score comes from
**`app.public.job_simulation_sessions.score`** — the **canonical** session entity in the `app` service's `public`
schema, read directly by the resolver.

> **History, because a seeder built from the old shape writes to nothing.** Until `app` migration
> `20260729133514.sql` (2026-07-29) the score lived on `public.local_jobsimulation_sessions`, a `Float32`
> **MIRROR** that shadowed a `jobsimulation.sessions` row. That migration **re-pointed the referencing
> assignment-session link ids** and **dropped the mirror** (`:58-62`) — it did *not* back-fill (`SET "score"`
> = 0 hits set-wide). The earlier `20260722104506.sql:79` dropped **`public.sessions`** (a bare
> `DROP TABLE "sessions"` under `search_path=public`) in favour of `public.job_simulation_sessions` (`:2`);
> **`jobsimulation.sessions` itself was NOT dropped** — no `app` migration touches that schema, and in
> **production** it survives frozen until M710. (`askengine/registry.go:192` is cited for the M710 horizon
> only: it is an LLM-facing name-alias map whose `jobsimulation.*` names **resolve to the public tables** —
> it is not evidence that the schema is physically present. On a **local dev/demo stack it is not**:
> `jobsimulation` has had **no `repos.yml` entry at all** since `d11a403` (6 entries @ platform `0dab54d`,
> 4 since `838d907`), and `app` (`repos.yml:14-17`) is the only repo with migrations to run. Qualified M257x iter-52,
> re-anchored in the M257x sweep.) So what is gone is the **mirror half** of the old
> pair, not both halves; there is one row per (candidate × attempt) now, in `public`. Corrected M257x iter-49.

**The read-path, traced end-to-end (FE → GraphQL → resolver → Ent → table):**

| Step | Location | What it does |
|------|----------|--------------|
| 1 | `apps/{web,hiring}/.../simulationScoreColumn.tsx:54,95-97` | renders `row.score` (the visible number). The two copies are **byte-identical**, same line numbers — but a *genuine* hiring org renders the `apps/hiring` one (§ *The render path*) |
| 2 | `packages/graphql/src/query/insights.ts:31-82` | query `insightsJobSimulationByMemberships` |
| 3 | `app/.../resolver_queries.go:1034,1080` | resolver `InsightsJobSimulationByMemberships` (decl `:1034`) → `IntelligenceManager.InsightsJobSimulationByMemberships` (`:1080`) |
| 4 | `app/internal/organization/intelligence.go:1700` | reads `m.ent.JobSimulationSession` (the canonical entity; was `LocalJobsimulationSession` before the mirror drop) |
| 5 | `intelligence.go:1733` → `:2158-2160` | best-attempt: the call site (`:1733`) invokes `usersBestOrFirstJobSimulationSession` (declared `:2124`), whose window is `ROW_NUMBER() OVER (PARTITION BY sim_id, owner_id ORDER BY score DESC …)` at `:2158-2160`; the `RowNumber == 1` reduction back in the caller is `:1739-1744`. (This cited the caller's own lines until M257x iter-102 — that range holds the `onlyAssignments` branch and the call, not the `row_number()`.) |
| 6 | `intelligence.go:1820` | `Score` ← the session's own `score` column — **not a mirror's** (see row 7) |
| 7 | `app/internal/data/ent/schema/job_simulation_session.go:45` | Ent table `public.job_simulation_sessions`, `field.Float32("score").Default(0).Min(0).Max(100)` — **the score column, read at `intelligence.go:1820` and assigned at `:1846`. Not a mirror: `local_jobsimulation_session.go` no longer exists** |

**The best-attempt sort + the cohort** (`intelligence.go:1738-1764`): rows are grouped per `user_id`, reduced to
**ONE best-attempt row per candidate** (the highest `score`), then sorted `score DESC, completition_status ASC,
session_started_at DESC`. Candidates are **comparable** when they share the same **`sim_id`** +
`organization_id` — that pair defines **one comparable cohort** (one scoreboard). Both halves are the
resolver's own predicates: `jobsimulationsession.SimID(jobSimulationId)` (`intelligence.go:1702`) +
`jobsimulationsession.OrganizationID(organizationID)` (`:1709`); and the best-attempt window partitions on
that same sim column per candidate — `ROW_NUMBER() OVER (PARTITION BY sim_id, owner_id ORDER BY score DESC
…)` (`intelligence.go:2158-2160`, inside `usersBestOrFirstJobSimulationSession`, declared `:2124`).
(This said `jobsimulation_id` until M257x iter-102 — that column belonged to the **dropped**
`local_jobsimulation_sessions` mirror, never to `job_simulation_sessions`; see the `job_position` bullet
above. Measured at `app` `ad9f3c49`.)

**The silent-403 substrate:** the resolver gates on the **`OrgFeatureInsights` Casbin permission**
(`resolver_queries.go:1035` — the first statement of `InsightsJobSimulationByMemberships`; `:1089` is the
`GetMembership` call inside the *neighbouring* `InsightsJobSimulationBySessions`, whose own gate is `:1085`).
Without that permission the query returns a **silent 403** and the scoreboard is empty
regardless of data — so the seeder must replicate whatever grants the existing demo orgs the insights permission.

**BA-4 — the drill-down is a DIFFERENT set of tables (not the scoreboard).** Clicking a candidate opens the
per-session competency / Job-Fit detail (`[simId]/[userId]`), which reads
`public.validation_attempt_results` / `validation_attempt_skill_results` / `validation_criterion_results` — three
tables (all in **`public`**, `20260722081626_jobsim_data_model.sql:336/355/376`; note the middle one is
`validation_attempt_skill_results`, not `validation_skill_results`) the `PersonaSeeder` also writes (`rosetta-extensions/stack-seeding/seeders/persona_write.go:69-71,143-167` @ rext `63ce41a` — path-qualified and re-pinned M257x iter-126; a bare basename is ambiguous and `D-M257x-122-5` refuses to resolve one by proximity). These are needed **only for the
drill-down**, NOT for the comparison list. The anticheat badge is a **decorative icon only**, and it is **not a
column on the session row** — it is `summary` on the separate **`public.anticheat_results`** entity
(`ent/schema/anticheat_result.go:24`), whose `session_id` FK was re-pointed at `job_simulation_sessions` by
`20260722104506.sql:53`. So
the open BA-1 question — *"does the list score need a per-session `validation_*`/eval row?"* — is answered **NO**:
the scoreboard scores from the **single** `job_simulation_sessions` row (+ membership + the Casbin gate)
alone — the write-set used to be a PAIR and is now one row, since the mirrors were dropped.

## The seeder-output contract (the write-set M223/M224 build against)

**Minimal write-set per (candidate × sim):**

1. **`public.job_simulation_sessions`** — the **score source** + row generator, and the only session row there
   is. Non-null `status`, `started_at`, `ended_at`, `owner_id`, `sim_id`, `sim_type`, **`token`**, plus
   `score` (0–100),
   `completion_status` (a closed 5-value enum — **exactly** `pending` / `passed` / `failed` / `discarded` /
   `timedout`, `app/internal/data/ent/enum/jobsimulation.go:29-35`, `Values()` at `:37-43`; **no `SIMULATION…`
   member** — that prefix belongs to the adjacent `sim_type` column, which genuinely is `SIMULATION_TYPE_*`),
   `organization_id`, `tenant_id` (NULL or `=org`), `validation_version`.
   ⚠️ **`token` is the one column that makes the INSERT itself fail, and this contract omitted it until
   M257x iter-49.** It is `NOT NULL` (`20260722104506.sql:13`), `UNIQUE` (`:29`) and carries **no default** —
   one of **four** required-and-undefaulted columns (`owner_id` `:6`, `sim_id` `:7`, `sim_type` `:10`,
   `token` `:13`; every other `NOT NULL` column in the DDL carries a `DEFAULT`) — so an INSERT built from the write-set as it was
   written here does not render wrong, it **errors**. The shipped seeder has always written it
   (`rosetta-extensions/stack-seeding/seeders/persona_write.go:152-158` @ rext `63ce41a` — path-qualified and re-pinned M257x iter-126); the word `token` simply appeared nowhere in this document. Being UNIQUE, it
   must be generated per row, not reused. (iter-47 read this passage and booked it a MINOR; iter-48's seat
   escalated it after checking the DDL, the Ent schema and the seeder.)
   ⚠️ **Get `completion_status` wrong and NOTHING catches it — the row does not vanish, it renders wrong.**
   The column is a plain `varchar` with **no CHECK** (a rolled-back
   `UPDATE … SET completion_status='SIMULATION_COMPLETION_STATUS_PASSED'` is accepted); Ent's generated
   `assignValues` casts **unconditionally** and cannot error (`ent/jobsimulationsession.go:181-186`:
   `_m.CompletionStatus = enum.SessionCompletionStatus(value.String)`); the read-model re-casts just as blindly
   (`intelligence.go:1844`); and the gqlgen enum marshal is a bare `graphql.MarshalString(string(v))`
   passthrough with **no** membership check (`graphql/graph/graph.go:129546-129554`, and the proto-bound twin at
   `:129392-129400`) — even though the SDL declares only the five lowercase members
   (`graphql/graph/schemas/jobsimulations.graphqls:14` and `:128`). So a raw-SQL seeder writing a `SIMULATION_…`
   value INSERTs cleanly **and the value travels verbatim all the way to the browser**, where it is a status the
   UI cannot map and that no completion-status filter — nor the `completition_status` sort — will ever match.
   *Earlier revisions of this doc said the row "vanishes at Ent scan / GraphQL enum marshal". It does not:
   there is no rejection anywhere on the path, which is precisely why the mistake is expensive.*
   ⚠️ **Do NOT write `anticheat_summary` here — the column does not exist on this table.** It was a column of
   the **dropped** `local_jobsimulation_sessions` mirror (added `20250416091037.sql:5`, dropped with the table
   at `20260729133514.sql:62`); `job_simulation_session.go` declares no such field. An INSERT built from a
   contract listing it fails. Anticheat is its own optional row in `public.anticheat_results`.
   ⚠️ **The column is spelled `completion_status` — correctly** (`20260722104506.sql:12`,
   `ent/schema/job_simulation_session.go:39`). The `completition` misspelling exists only in the GraphQL
   sort-field enum (`enum.InsightsSortFieldCompletitionStatus`), its GraphQL member and a JSON tag;
   `insightsSortColumn` (`intelligence.go:885-886`) maps it back to `FieldCompletionStatus`, so it
   **never reaches SQL**.
2. **`public.memberships`** — the candidate must be **active** (`GetMemberships`; status `active`/`invited`).

> **The write-set used to be a PAIR and is now a single row** (M257x iter-23). Before the mirror drop it was
> `public.local_jobsimulation_sessions` (score) + a co-written `jobsimulation.sessions` twin (so the federated
> non-null `Session!` resolved from the other subgraph, else the list NULL-bubbled). The mirror table is
> dropped and there is no second subgraph, so **both halves collapsed into `public.job_simulation_sessions`.**
> (This said *"neither table exists"* until M257x iter-52 — a flat contradiction of the **History** blockquote
> (**`:196-209`**) above, which is the adjudicated form: the mirror was dropped, `jobsimulation.sessions` was
> not, and no local stack creates that schema either way. It cited an earlier range — the `job_position`
> bullet, now **`:176-187`**, a different construct entirely — until M257x iter-98. **Both pins in this
> sentence were themselves wrong until run 81** (one landed on the score sentence + a blank line; the other
> on the tail of the *preceding* bullet), because iter-102 re-derived them by **adding +23 and
> +16 to the old numbers instead of re-measuring** — arithmetic on a citation is not a citation. Both in-file line numbers were
> re-derived at M257x iter-102, when the repairs above shifted them by +23 and +16.) The old
> "393/393 rows on `billion` carry a matching pair" empiric described the pre-drop shape.

**Org prerequisites:** `public.organizations.is_hiring = true` (§ *the gate*) + Clerk `publicMetadata.isHiring =
true` (M224) + the **`OrgFeatureInsights` Casbin permission** substrate.

**The machinery already exists — M223 is NOT net-new.** The **`PersonaSeeder` already writes exactly this row** —
`rosetta-extensions/stack-seeding/seeders/persona_write.go:91` writes
`{"public", "job_simulation_sessions", sessionCols(), …}`. (Until M257 it wrote the **pair**, via a second col
builder `localSessionCols()`; that builder was deleted with the mirror and `sessionCols()` at `:152` now serves
the single canonical row.) M223's
candidate-assessment funnel is a **direct generalization** of the same fan-out — each candidate on the **one**
position they applied for (v2.4 "casting call" M227 fix #3; before M227 every candidate took all 5), round-robined
evenly across the 5 shared sims so ~8 candidates rank per position (the M51 `AIReadinessFunnelSeeder` shape, 2 shared
sims → 5) — with the M219 anti-junk discipline (a realistic non-degenerate score DISTRIBUTION, every skill/role ref
through the real resolvers, closure green, never fabricated), **not** a flat score grid.

## `isHiringOrg`, the `isEnterprise` divergence, and the `is_hiring` blast radius

`isHiringOrg` is **client-derived** (`useGetClerkOrganization.tsx:20-21`, above). What the flip changes:

- **The comparison surface SURVIVES** — it is only **RELABELED "Results"** (vs "Activity dashboard"):
  `packages/ui/src/NavBar/useNavbarSections.tsx:476`, inside `enterpriseInsightsMenuItem` (`:475-482`)
  (`label: isHiringOrg ? tNavbar('results') : tNavbar('activityDashboard')`) — measured at `next-web-app`
  `8297c684`; both anchors carried the earlier ref's line numbers until M257x iter-102, exact at that earlier `bb3313bc`.
  It stays in `enterpriseAdminNavbarMenuItems`; the route `/enterprise/activity-dashboard` has **no `is_hiring`
  guard**.
- **Two `isEnterprise` definitions DIVERGE — and that is not a bug:**
  - **Nav:** `isEnterprise = Boolean(organization)` (`template.tsx:90`) → stays **TRUE** for a hiring org, so the
    enterprise nav still renders.
  - **Billing:** `isEnterprise = !isHiringOrg && organizationId` (`FreeTrialContainer.tsx:29`) → flips **FALSE**, so
    hiring orgs are **excluded from the Workforce free-trial** container. Irrelevant to the comparison; enumerated
    here so a future reader doesn't "fix" the divergence.
- **Also under `is_hiring`:** the nav trims the Content Library to **AI-Simulations** alone
  (`packages/ui/src/NavBar/useNavbarSections.tsx:340-343` — `isHiringOrg` selects
  `[librarySimulationsMenuItem]`, label `tNavbar('aiSimulations')` at `:249-256`, instead of the
  simulations + skill-paths + academy + labs set), and hides the member Profile / Skills / Activities
  entries for non-admins (`:329-331`, each `!isHiringOrg || isAdmin`). Both clauses verify.
  ⚠️ **It does NOT gate Workforce Intelligence off — that clause is RETRACTED (M257x iter-102).** Nothing
  gates Workforce Intelligence on `isHiringOrg`. The entry is `enterpriseWorkforceMenuItem`
  (`tNavbar('workforceIntelligence')`, `:391-398`); it sits in the `intelligence` group, whose visibility
  comes from `orgSectionVisibility({ isAdmin, showStudio })` returning `intelligence: isAdmin`
  (`packages/ui/src/NavBar/orgGroups.ts:48-65`, the field at `:61`) — a function that **takes no
  `isHiringOrg` parameter at all** — and the item itself is gated on `showWorkforce`
  (`useNavbarSections.tsx:568`), which **defaults to `true`** (`useNavbarSections.tsx:161`) and is passed
  `false` in exactly **two** places, **both in `apps/hiring`**
  (`apps/hiring/src/app/(authenticated)/(verified)/template.tsx:167` and `:248`). So a recruiter loses
  Workforce Intelligence by being handed off to **`apps/hiring`** (§ *The render path*), not by the
  `is_hiring` flag inside `apps/web`; an `is_hiring` org's admin still browsing `apps/web` keeps it.
  Measured at `next-web-app` `8297c684`. None of these touch the comparison scoreboard.

## Interface

- **GraphQL** (the single `backend` subgraph, read by **both** `apps/web` and `apps/hiring` — the query lives in the shared `packages/graphql`): `insightsJobSimulationByMemberships` (`packages/graphql/src/query/insights.ts`)
  → `app` subgraph resolver `resolver_queries.go` → `IntelligenceManager.InsightsJobSimulationByMemberships`
  (`app/internal/organization/intelligence.go`). Gated on the `OrgFeatureInsights` Casbin permission.
- **The `Session!` field** resolves from the **same** subgraph and the **same** row — there is only one subgraph
  now, and only one session table. (Before the merges it resolved from `jobsimulation.sessions` in a *different*
  subgraph, joined on the mirror's `jobsimulation_session_id`, and a missing twin NULL-bubbled the row out of the
  list. **No join key, no NULL-bubble hazard, no twin to forget.**)
- **Surface:** route `/enterprise/activity-dashboard → AI-Simulations → [simId]` (list) + `.../[simId]/[userId]`
  (the per-candidate drill-down, reads the `public.validation_*` tables) — present in **both** `apps/web` and
  `apps/hiring`, but for a *genuine* hiring org only the `apps/hiring` copy is reachable (§ *The render path*).

## Local development

*(This section **previously said** `jobsimulation.sessions` still exists on a local stack; that is
**refuted** — M257x iter-52. There is no `jobsimulation` schema on a dev/demo stack at all.)*

To make a hiring org's comparison scoreboard render on a demo/dev stack: seed an org with `is_hiring=true`
(+ Clerkenstein `publicMetadata.isHiring=true`, M224), an active membership per candidate, and — per (candidate ×
sim) — **one** `public.job_simulation_sessions` row (the score lives on it; the old co-written
`jobsimulation.sessions` + `public.local_jobsimulation_sessions` **pair** is no longer written — the
`public.local_jobsimulation_sessions` **mirror** was dropped at `20260729133514.sql:58-62`, and **on a
dev/demo stack there is no `jobsimulation` schema at all** — `platform/repos.yml` has **no
`jobsimulation` entry** (6 entries @ `0dab54d`, 4 since `838d907`; `d11a403` deleted it), and app's only `CREATE SCHEMA` is `auth`
(`20230817154747_supabase_baseline.sql:2`), so `to_regclass('jobsimulation.sessions')`
is NULL. Do not seed into it. This passage **previously said** `jobsimulation.sessions` *"still exists,
frozen and unwritten, until M710"* unqualified — **refuted** for local stacks at M257x iter-52), plus the
`OrgFeatureInsights` Casbin grant. Pick 5 real `SIMULATION_TYPE_HIRING` sims from the captured
snapshot as the org's positions. The scoreboard then reads `insightsJobSimulationByMemberships`, one best-attempt
row per candidate. The drill-down additionally needs the `public.validation_attempt_results` /
`validation_attempt_skill_results` / `validation_criterion_results` rows (the PersonaSeeder pattern).

> **This is IMPLEMENTED as of v2.4 "casting call" M223** (`rosetta-extensions/stack-seeding`): the
> `stories.seed.yaml` 4th story (Meridian Talent, `narrative: hiring`, 5 admins + 45 candidates) + two seeders —
> **`HiringConfigSeeder`** (the 5 positions via the type-aware `readHiringSimPool`, written as
> `organization_sim_invitation_links`) and **`HiringFunnelSeeder`** (each candidate's scored `SIMULATION_TYPE_HIRING`
> session row — a MIRROR *pair* until the M257 re-point dropped the mirror — on the **one** position they applied for — round-robined evenly across the 5, ~8 per position (M227
> fix #3) — SOME assigned-only, a differentiated score spread). The
> `OrgFeatureInsights` substrate needs **no net-new grant** — the org's `admin` members inherit `org:feature:insights`
> from the global `p3` admin policy via their standard g2 grant. Seeder chain: [`../ops/demo/stories-spec.md`](../ops/demo/stories-spec.md#the-m223-hiring-chain--two-seeders-hiring-config--hiring-funnel)
> + [`../ops/seeding-spec.md`](../ops/seeding-spec.md#the-recruiter-vantage--the-hiring-org--candidate-comparison-funnel-v24-casting-call-m223).
> M223 does NOT ship the render proof or the cockpit heroes (M224); the per-candidate drill-down `validation_*`
> rows are also M224+ (the M223 scoreboard needs only the single `job_simulation_sessions` row —
> formerly a 2-table pair, until the mirrors were dropped).

## The render path (v2.4 "casting call" M224 — the two-app demo)

**M224 proved the render — and it does NOT land in `apps/web`.** M222 traced the comparison scoreboard rendering
in the dockerized `apps/web` `/enterprise` and concluded it was reachable by the recruiter. That held on the
`billion` spike **only because that org had no client `publicMetadata.isHiring`** — client-side it read as a
*workforce* org, so it never tripped the product-boundary guard. M224 **wired client `isHiring=true`** (D-DESIGN-1
— the org must *genuinely read as hiring*), and that flips the calculus:

- On the *unmodified* platform, a user whose memberships are **all hiring orgs** is **ejected out of `apps/web`**
  to the standalone Hiring product (`apps/web/src/context/UserStatusContext.tsx` → `buildSwitchHandoffUrl({
  targetProduct: 'hiring' })`, **by design** — a global guard that fires even on a direct navigation), and
  `useGetClerkOrganization` filters hiring orgs out of the workforce list. **⇒ "reads as hiring in the browser"
  and "scoreboard reachable in `apps/web`" are mutually exclusive on the real platform.** (This falsified M222's
  `apps/web` premise — #M224-D-TOK02.)

**Resolution — run the genuine `apps/hiring` as a second UI container (TOK-02).** The real candidate-comparison
Results screen ships in `apps/hiring`
(`.../enterprise/activity-dashboard/@tabs/ai-simulations/[simId]/page.tsx` → `InsightsByMembersContainer`), and it
federates the **same** `insightsJobSimulationByMemberships` field (in the **app** subgraph SDL, **no feature
gate**) over the **same** GraphQL endpoint the demo already serves — which since platform `2adcf71`
(2026-07-31) is **`backend`'s own `:8082/graphql/query`**, the Cosmo/WunderGraph router having been deleted
from compose — reading the **same** seeded `public.job_simulation_sessions` M223 wrote. So the demo builds `apps/hiring` from the **untouched clone** as
a second offset-port UI container (same recipe as `apps/web` + `studio-desk`), wired to the same fake FAPI +
the same `backend` GraphQL endpoint + Postgres — **no re-skin, no new resolver, no data migration, zero
platform-repo edits.** The recruiter logs in
straight onto the hiring Results page (the cockpit's `CockpitHero.IsHiring` routes her to the hiring base); the
platform's own symmetric guard keeps her *in*.

**What renders (gate met, ≥3 cold runs, 4/4 flake).** For **each** of the 5 shared sims the scoreboard paints its
comparable-candidate cohort — **~8 candidates per position** since **v2.4 "casting call" M227 fix #3** (each
candidate auditions on the ONE position they applied for, round-robined evenly across the 5; before M227 every
candidate took all 5 → ~43 on each, paged at the platform-native `useTablePagination` default 20 — **GATE-DECISION
D1**). With ~8 per position they all fit on page 1 (no pagination needed), non-degenerate (scores 27–100), **0 junk**
(closure green), **0 prod-eject**. **The compare gate retuned `≥40 → ≥6` (M227 fix #3** — a small margin below the
seeded min of ~8; `hiringComparableFloor` / the render-probe `RENDER_GATE_FLOOR`). Four demo-patches on the
hiring image make it land — 2 net-new (`next-hiring-role-remap`, `next-hiring-members-pagination`) + the 2 chained
shared `urls.ts` patches (the Studio-eject kill) — see [`../ops/demo/demopatch-spec.md`](../ops/demo/demopatch-spec.md)
§ the four hiring-image patches, and the cockpit trio in
[`../ops/demo/cockpit-spec.md`](../ops/demo/cockpit-spec.md) § the hiring vantage.

**Believability (v2.4 M227 "the notes", seed-only).** The recruiter's AI-Simulations list reads **hiring-only** — the
generic workforce activity seeders skip a hiring org (`hiring_scope.go` `IsHiringOrg()`, #M227-D1), so its whole sim
footprint is these 5 HIRING sims (no training/assessment leakage into the mirror the list groups by). Candidates read
as **outside applicants**: emails are keyed on **role** → an external consumer domain (gmail/outlook/…), only
admins/recruiters keep `@meridian-talent.com` (#M227-D2). See
[`../ops/demo/stories-spec.md`](../ops/demo/stories-spec.md#the-m223-hiring-chain--two-seeders-hiring-config--hiring-funnel).

**The M227 fix-#1 guard was INCOMPLETE — caught by the M228 live re-prove.** M227's "skip the generic seeders for a
hiring org" listed the obvious activity seeders but **missed two mirror/FK writers**, and the gap only surfaced when
the corrected seed actually ran on `billion` (the deterministic unit test had simply omitted them). **FeedbackSeeder**
has written app-side session rows on GENERIC sims since v1.10 M42m (then
`public.local_jobsimulation_sessions` MIRROR rows; `public.job_simulation_sessions` since the M257 re-point) → unguarded it leaked a
training sim + a 2nd session per candidate into the recruiter's list; **SuccessionSeeder** FKs each member's first
population session (now skipped for the hiring org) → the FK VIOLATED and the whole seed reported *"failed"*. Both now
consult `skipGenericActivityForHiringOrg`; the regression enumerates **all 8** generic seeders (#M228 F1/F2/F3). The
lesson is in-code: **a new mirror-writing / session-FK seeder MUST be added to the guard's consult-list AND the
enumerated `TestGenericActivitySeeders_SkipHiringOrg` table.**

**The render probe is drawer-aware (M228; mechanism corrected v2.5 M231).** The recruiter comparison drawer is a
**plain Ant `<Drawer>`** rendered by `InsightsByMembersContainer` (`InsightsByMembersContainer.tsx:359`) on the
ordinary `@tabs/ai-simulations/[simId]` leaf route — **NOT** a Next.js intercepting route (verified M231: zero
`(.)`/`(..)`-prefixed dirs exist anywhere in `apps/`; the earlier `…/@tabs/(.)ai-simulations/[simId]` intercepting-route
description was wrong). It still mounts as a detectable `.ant-drawer` — firing its client
`insightsJobSimulationByMemberships` POST + becoming DOM-visible — **only for the FIRST sim clicked per page-load**. Later sims in the same session *do* render (server-path RSC, screenshot-proven) but aren't cleanly
detected. So the render gate proves each of the 5 positions in its **own** run (`RENDER_ONLY_SIM`, each sim as "the
first") rather than clicking all 5 in one session — a clean automated **5/5** (Talent-Mgr 8, BD-Lead 8, Inside-Sales
9, Project-Mgr 9, AWS-Security 8; each ≥ floor 6, junk=0, 0 eject). Proven live on `billion` with recruiter p95
click→ACCESS **1.27 s**. (#M228, render-probe `stack-verify/e2e/tests/render-hiring-comparison.spec.ts`.)

## Cross-references

- The frontend split that hosts the surface: [`next-web-app.md`](next-web-app.md) (Workforce `apps/web` vs Hiring
  `apps/hiring`). **⚠️ M222 inferred the scoreboard was reachable in `apps/web`; M224 rendering proved it is not for a
  *genuine* hiring org — the demo serves the real `apps/hiring` as a second container (see § The render path above).**
- The simulation runtime — now a **domain inside `app`** (`app/internal/jobsimulation/`), owning
  `public.job_simulation_sessions` + the `public.validation_*` drill-down tables:
  [`jobsimulation.md`](jobsimulation.md).
- The `app` service that owns the session table + the `IntelligenceManager` read-model + the
  `OrgFeatureInsights` Casbin gate: [`backend.md`](backend.md).
- The Clerk mock that must emit `publicMetadata.isHiring` (M224): [`clerkenstein.md`](clerkenstein.md).
- The closest seeding precedent (a narrative-gated org feature with a funnel seeder + the same
  render-gate-bypasses-the-seed lesson): [`ai-readiness.md`](ai-readiness.md).
- The seeded demo world the hiring org joins as the 4th story: [`../ops/demo/stories-spec.md`](../ops/demo/stories-spec.md)
  + the blueprint gate in [`../ops/seeding-spec.md`](../ops/seeding-spec.md).
