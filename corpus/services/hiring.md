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

## Role & Responsibility

**Hiring** is a **sold product** (`hiring.anthropos.work`, the `@anthropos/hiring-app` at `apps/hiring`) AND an
**org-type** that re-skins the **Workforce** app (`apps/web`) for a recruiting buyer. The demo cares about the
**org-type**, not the standalone Hiring app: an `is_hiring` organization runs its members through
**`SIMULATION_TYPE_HIRING`** job simulations and reads a **candidate-comparison scoreboard** — the recruiter's core
value: *"line up every candidate who took this hiring simulation, ranked by score."* That scoreboard lives in the
**dockerized `apps/web`** (`/enterprise/activity-dashboard`), renders **from seedable data alone**, and **survives
the `is_hiring` flip** — so the demo can show it **without a platform edit** (M222 D1, the release go/no-go).

There is **no `hiring` microservice**. The feature is a composition: an org-type flag on `app`, the existing
`jobsimulation` runtime, an `app`-side read-model (`IntelligenceManager`), the federated GraphQL, and an `apps/web`
surface gated client-side on a Clerk org flag.

## The org-type gate — `is_hiring` is a DUAL-WRITE

`is_hiring` must be set in **TWO** places, because the platform derives the org's type differently on each side:

1. **Backend — `public.organizations.is_hiring boolean NOT NULL default false`.** The server-side org-type. The
   seeder writes it directly (M222 landed the gate — see § *The seeder-output contract*). The
   `resolver_queries.go` insights path (below) requires it `true` for the org's data to be treated as a hiring
   cohort.
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

> **Both, or the demo is half-lit.** DB-only → the browser doesn't re-skin. Clerk-only → the insights read-path
> won't treat the cohort as hiring. The seeder writes #1; the mock emits #2; M224 wires the pair.

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
  share the same `jobsimulation_id` — the sim IS the position for scoreboard purposes.

## The comparison read-model (THE HEADLINE) — the score is a MIRROR table in `app`

The recruiter's scoreboard is `/enterprise/activity-dashboard → AI-Simulations → [simId]`: one row per candidate who
took the sim, ranked by a comparable **score**. **That score does NOT come from `jobsimulation.sessions.score`.** It
comes from **`app.public.local_jobsimulation_sessions.score`** — a `Float32` **MIRROR** table in the `app` service's
own `public` schema, read directly by the resolver. Seeding only `jobsimulation.sessions` renders an **empty**
comparison.

**The read-path, traced end-to-end (FE → GraphQL → resolver → Ent → table):**

| Step | Location | What it does |
|------|----------|--------------|
| 1 | `apps/web/.../simulationScoreColumn.tsx:54,95-97` | renders `row.score` (the visible number) |
| 2 | `packages/graphql/src/query/insights.ts:31-82` | query `insightsJobSimulationByMemberships` |
| 3 | `app/.../resolver_queries.go:1088,1134` | resolver → `IntelligenceManager.InsightsJobSimulationByMemberships` |
| 4 | `app/internal/organization/intelligence.go:1692` | reads **ONLY** `m.ent.LocalJobsimulationSession` (the mirror) |
| 5 | `intelligence.go:1728-1735` | best-attempt: `row_number() ORDER BY score DESC` per candidate |
| 6 | `intelligence.go:1801` | `Score` ← `ls.Score` (the mirror's score column) |
| 7 | `app/internal/data/ent/schema/local_jobsimulation_session.go:52` | Ent table `public.local_jobsimulation_sessions`, `field.Float32("score")` |

**The best-attempt sort + the cohort** (`intelligence.go:1738-1751`): rows are grouped per `user_id`, reduced to
**ONE best-attempt row per candidate** (the highest `score`), then sorted `score DESC, completition_status ASC,
session_started_at DESC`. Candidates are **comparable** when they share the same `jobsimulation_id` +
`organization_id` — that pair defines **one comparable cohort** (one scoreboard).

**The silent-403 substrate:** the resolver gates on the **`OrgFeatureInsights` Casbin permission**
(`resolver_queries.go:1089`). Without that permission the query returns a **silent 403** and the scoreboard is empty
regardless of data — so the seeder must replicate whatever grants the existing demo orgs the insights permission.

**BA-4 — the drill-down is a DIFFERENT set of tables (not the scoreboard).** Clicking a candidate opens the
per-session competency / Job-Fit detail (`[simId]/[userId]`), which reads
`jobsimulation.validation_attempt_results` / `validation_skill_results` / `validation_criterion_results` — three
tables the `PersonaSeeder` also writes (`persona_write.go:69-71,143-167`). These are needed **only for the
drill-down**, NOT for the comparison list. `anticheat_summary` on the mirror row is a **decorative icon only**. So
the open BA-1 question — *"does the list score need a per-session `validation_*`/eval row?"* — is answered **NO**:
the scoreboard scores from the 2-table pair (+ membership + the Casbin gate) alone.

## The seeder-output contract (the write-set M223/M224 build against)

**Minimal write-set per (candidate × sim):**

1. **`public.local_jobsimulation_sessions`** — the **score source** + row generator. Fields: `score` (0–100),
   `completition_status` (**note the misspelled column**; values `passed`/`failed`/`pending`/`SIMULATION…`),
   `user_id`, `jobsimulation_id`, `jobsimulation_session_id` (FK → #2), `organization_id`, `tenant_id` (NULL or
   `=org`), `status`, `session_started_at`/`session_ended_at`, `validation_version`, `anticheat_summary` (optional).
2. **`jobsimulation.sessions`** — required so the federated **non-null `Session!`** (status / startedAt) resolves;
   else the whole list **NULL-bubbles** (a federation non-null on a missing row collapses the parent). Fields: `id`
   (= #1's `jobsimulation_session_id`), non-null `status`, `started_at`, `ended_at`, `owner_id`, `sim_id`,
   `sim_type`. **Empirically 393/393 local rows on `billion` carry this matching pair** — the mirror row and its
   `jobsimulation.sessions` twin are always co-written.
3. **`public.memberships`** — the candidate must be **active** (`GetMemberships`; status `active`/`invited`).

**Org prerequisites:** `public.organizations.is_hiring = true` (§ *the gate*) + Clerk `publicMetadata.isHiring =
true` (M224) + the **`OrgFeatureInsights` Casbin permission** substrate.

**The machinery already exists — M223 is NOT net-new.** The current **`PersonaSeeder` already writes exactly this
pair** — `rosetta-extensions/stack-seeding/seeders/persona_write.go:68-73` writes both `jobsimulation.sessions` and
`public.local_jobsimulation_sessions` (col builders `sessionCols()` / `localSessionCols()` at `:125-141`). M223's
candidate-assessment funnel is a **direct generalization** of the same fan-out — each candidate on the **one**
position they applied for (v2.4 "casting call" M227 fix #3; before M227 every candidate took all 5), round-robined
evenly across the 5 shared sims so ~8 candidates rank per position (the M51 `AIReadinessFunnelSeeder` shape, 2 shared
sims → 5) — with the M219 anti-junk discipline (a realistic non-degenerate score DISTRIBUTION, every skill/role ref
through the real resolvers, closure green, never fabricated), **not** a flat score grid.

## `isHiringOrg`, the `isEnterprise` divergence, and the `is_hiring` blast radius

`isHiringOrg` is **client-derived** (`useGetClerkOrganization.tsx:20-21`, above). What the flip changes:

- **The comparison surface SURVIVES** — it is only **RELABELED "Results"** (vs "Activity dashboard"):
  `packages/ui/src/NavBar/useNavbarSections.tsx:300-307` (`label: isHiringOrg ? 'results' : 'activityDashboard'`).
  It stays in `enterpriseAdminNavbarMenuItems`; the route `/enterprise/activity-dashboard` has **no `is_hiring`
  guard**.
- **Two `isEnterprise` definitions DIVERGE — and that is not a bug:**
  - **Nav:** `isEnterprise = Boolean(organization)` (`template.tsx:90`) → stays **TRUE** for a hiring org, so the
    enterprise nav still renders.
  - **Billing:** `isEnterprise = !isHiringOrg && organizationId` (`FreeTrialContainer.tsx:29`) → flips **FALSE**, so
    hiring orgs are **excluded from the Workforce free-trial** container. Irrelevant to the comparison; enumerated
    here so a future reader doesn't "fix" the divergence.
- **Also under `is_hiring`:** the nav trims the library to **AI-Simulations**, hides some member surfaces for
  non-admins, and gates **Workforce Intelligence off**. None of these touch the comparison scoreboard.

## Interface

- **GraphQL** (federated, `apps/web`): `insightsJobSimulationByMemberships` (`packages/graphql/src/query/insights.ts`)
  → `app` subgraph resolver `resolver_queries.go` → `IntelligenceManager.InsightsJobSimulationByMemberships`
  (`app/internal/organization/intelligence.go`). Gated on the `OrgFeatureInsights` Casbin permission.
- **The federated `Session!`** is resolved from `jobsimulation.sessions` (a different subgraph) — the mirror row's
  `jobsimulation_session_id` is the join key. A missing twin NULL-bubbles the row out of the list.
- **Surface:** `apps/web` route `/enterprise/activity-dashboard → AI-Simulations → [simId]` (list) +
  `.../[simId]/[userId]` (the per-candidate drill-down, reads the `jobsimulation.validation_*` tables).

## Local development

To make a hiring org's comparison scoreboard render on a demo/dev stack: seed an org with `is_hiring=true`
(+ Clerkenstein `publicMetadata.isHiring=true`, M224), an active membership per candidate, and — per (candidate ×
sim) — the co-written `jobsimulation.sessions` + `public.local_jobsimulation_sessions` pair (the score lives on the
mirror), plus the `OrgFeatureInsights` Casbin grant. Pick 5 real `SIMULATION_TYPE_HIRING` sims from the captured
snapshot as the org's positions. The scoreboard then reads `insightsJobSimulationByMemberships`, one best-attempt
row per candidate. The drill-down additionally needs the `jobsimulation.validation_attempt_results`/`_skill_results`/
`_criterion_results` rows (the PersonaSeeder pattern).

> **This is IMPLEMENTED as of v2.4 "casting call" M223** (`rosetta-extensions/stack-seeding`): the
> `stories.seed.yaml` 4th story (Meridian Talent, `narrative: hiring`, 5 admins + 45 candidates) + two seeders —
> **`HiringConfigSeeder`** (the 5 positions via the type-aware `readHiringSimPool`, written as
> `organization_sim_invitation_links`) and **`HiringFunnelSeeder`** (each candidate's scored `SIMULATION_TYPE_HIRING`
> MIRROR pair on the **one** position they applied for — round-robined evenly across the 5, ~8 per position (M227
> fix #3) — SOME assigned-only, a differentiated score spread). The
> `OrgFeatureInsights` substrate needs **no net-new grant** — the org's `admin` members inherit `org:feature:insights`
> from the global `p3` admin policy via their standard g2 grant. Seeder chain: [`../ops/demo/stories-spec.md`](../ops/demo/stories-spec.md#the-m223-hiring-chain--two-seeders-hiring-config--hiring-funnel)
> + [`../ops/seeding-spec.md`](../ops/seeding-spec.md#the-recruiter-vantage--the-hiring-org--candidate-comparison-funnel-v24-casting-call-m223).
> M223 does NOT ship the render proof or the cockpit heroes (M224); the per-candidate drill-down `validation_*`
> rows are also M224+ (the M223 scoreboard needs only the 2-table pair).

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
gate**) over the **same** Cosmo router the demo already bakes, reading the **same** seeded
`public.local_jobsimulation_sessions` M223 wrote. So the demo builds `apps/hiring` from the **untouched clone** as
a second offset-port UI container (same recipe as `apps/web` + `studio-desk`), wired to the same fake FAPI + Cosmo
+ Postgres — **no re-skin, no new resolver, no data migration, zero platform-repo edits.** The recruiter logs in
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

## Cross-references

- The frontend split that hosts the surface: [`next-web-app.md`](next-web-app.md) (Workforce `apps/web` vs Hiring
  `apps/hiring`). **⚠️ M222 inferred the scoreboard was reachable in `apps/web`; M224 rendering proved it is not for a
  *genuine* hiring org — the demo serves the real `apps/hiring` as a second container (see § The render path above).**
- The runtime that runs the sims + owns `jobsimulation.sessions` + the `validation_*` drill-down tables:
  [`jobsimulation.md`](jobsimulation.md).
- The `app` service that owns the **mirror** table `public.local_jobsimulation_sessions` + the `IntelligenceManager`
  read-model + the `OrgFeatureInsights` Casbin gate: [`backend.md`](backend.md).
- The Clerk mock that must emit `publicMetadata.isHiring` (M224): [`clerkenstein.md`](clerkenstein.md).
- The closest seeding precedent (a narrative-gated org feature with a funnel seeder + the same
  render-gate-bypasses-the-seed lesson): [`ai-readiness.md`](ai-readiness.md).
- The seeded demo world the hiring org joins as the 4th story: [`../ops/demo/stories-spec.md`](../ops/demo/stories-spec.md)
  + the blueprint gate in [`../ops/seeding-spec.md`](../ops/seeding-spec.md).
