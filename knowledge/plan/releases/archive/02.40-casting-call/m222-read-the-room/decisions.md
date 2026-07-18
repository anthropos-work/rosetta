# M222 — Decisions

_Implementation choices with rationale. One entry per decision; cite the code/doc it binds._

---

## D1 — **GO.** The recruiter comparison surface renders from seedable data in the dockerized `apps/web`. (BA-3 refuted; the release proceeds.)

**Decision:** the release's make-or-break go/no-go is **GO**. The candidate-comparison surface
`/enterprise/activity-dashboard → AI-Simulations → [simId]` lives in the **dockerized `apps/web`** (the
`@anthropos/web-app` the demo already builds), renders **from DATA ALONE** (no platform edit), and **survives the
`is_hiring` flip**. BA-3's escalation trigger — "the view is `apps/hiring`-only, so showing it = containerizing a
Vercel app + a platform edit" — is **refuted**. M223/M224 may commit against the contract in D2.

**Evidence (empirically traced/reproduced on the live `billion` substrate):**
- The route `/enterprise/activity-dashboard` has **no `is_hiring` guard**; under `is_hiring` it stays in
  `enterpriseAdminNavbarMenuItems` and is only **RELABELED** "Results" (vs "Activity dashboard") —
  `packages/ui/src/NavBar/useNavbarSections.tsx:300-307` (`label: isHiringOrg ? 'results' : 'activityDashboard'`).
- The score column renders `row.score` from the seedable read-path (see D2) —
  `apps/web/.../simulationScoreColumn.tsx:54,95-97`.

**Consequence:** no escalation. The risk **R2 (blocks-release)** is retired.

---

## D2 — The seeder-output contract: the score is a **MIRROR table** in `app`, NOT `jobsimulation.sessions`. (The M219-class render-gate trap, spelled out.)

**Decision:** the comparable per-candidate score the surface renders comes from
**`app.public.local_jobsimulation_sessions.score`** (a `Float32` MIRROR table in the `app` service's `public`
schema, read directly by the resolver) — **NOT** from `jobsimulation.sessions.score`. Seeding only
`jobsimulation.sessions` renders an **EMPTY** comparison (the M219 "render gate bypasses the seed" failure class).
M223's funnel seeder writes the **local-mirror** table as the score source.

**The read-path trace (FE → GraphQL → resolver → Ent → table):**
1. FE reads `row.score` — `apps/web/.../simulationScoreColumn.tsx:54,95-97`.
2. Query `insightsJobSimulationByMemberships` — `packages/graphql/src/query/insights.ts:31-82`.
3. Resolver — `resolver_queries.go:1088,1134` → `IntelligenceManager.InsightsJobSimulationByMemberships`.
4. The resolver reads **ONLY** `m.ent.LocalJobsimulationSession` — `app/internal/organization/intelligence.go:1692`;
   best-attempt via `row_number() ORDER BY score DESC` (`:1728-1735`); `Score` from `ls.Score` (`:1801`).
5. Ent table = `public.local_jobsimulation_sessions`, `field.Float32("score")` —
   `app/internal/data/ent/schema/local_jobsimulation_session.go:52`.

**The minimal write-set per (candidate × sim):**
1. **`public.local_jobsimulation_sessions`** — the score source + row generator. Fields: `score` (0–100),
   `completition_status` (**note the misspelled column**; values `passed`/`failed`/`pending`/`SIMULATION…`),
   `user_id`, `jobsimulation_id`, `jobsimulation_session_id` (FK → #2), `organization_id`, `tenant_id`
   (NULL or `=org`), `status`, `session_started_at`/`session_ended_at`, `validation_version`,
   `anticheat_summary` (optional; a decorative icon only).
2. **`jobsimulation.sessions`** — required so the federated non-null `Session!` (status/startedAt) resolves; else
   the list **NULL-bubbles**. Fields: `id` (= #1's `jobsimulation_session_id`), non-null `status`, `started_at`,
   `ended_at`, `owner_id`, `sim_id`, `sim_type`. **Empirically: 393/393 local rows on `billion` carry this matching
   pair.**
3. **`public.memberships`** — the candidate must be active (`GetMemberships`; status `active`/`invited`).

**Org prerequisites:**
- `public.organizations.is_hiring = true` (the seeder writes it directly — see D3).
- The **`OrgFeatureInsights` Casbin permission** substrate — else a **silent 403** at `resolver_queries.go:1089`.
  Replicate whatever grants the existing demo orgs the insights permission.

**Sort / comparability:** grouped per `user_id`, **ONE best-attempt row (highest `score`) per candidate**, sorted
`score DESC, completition_status ASC, session_started_at DESC` (`intelligence.go:1738-1751`); the same
`jobsimulation_id` + `organization_id` = **one comparable cohort**.

**M223 reuses existing machinery — this is NOT net-new.** The current `PersonaSeeder` **already writes exactly this
pair** — `rosetta-extensions/stack-seeding/seeders/persona_write.go:68-73` writes `jobsimulation.sessions` +
`public.local_jobsimulation_sessions` (col builders `sessionCols()` / `localSessionCols()` at `:125-141`). M223
generalizes the same fan-out across 45 candidates × 5 shared sims.

**BA-4 (the drill-down, NOT the scoreboard):** the per-session competency/Job-Fit detail (`[simId]/[userId]`) reads
`jobsimulation.validation_attempt_results` / `_skill_results` / `_criterion_results` — the PersonaSeeder also writes
these (`persona_write.go:69-71,143-167`). They are needed **only for the drill-down**, not the comparison list. So
BA-1's open question — "does the list score need a per-session `validation_*`/eval row?" — is answered **NO**: the
comparison list scores from the 2-table pair (+ memberships + the Casbin gate) alone.

---

## D3 — `is_hiring` is a **DUAL-WRITE**: the `organizations` table AND Clerk `publicMetadata`. (M224 extends Clerkenstein.)

**Decision:** `is_hiring` must be set in **BOTH** places, because the surface's relabel/blast-radius is derived
**client-side from Clerk**, while the read-model's org-type is derived server-side from the DB:
1. **`public.organizations.is_hiring = true`** — the seeder writes it directly (M222 lands the gate — see the S2
   thread; M223 declares the 4th story with `is_hiring: true`).
2. **Clerk `publicMetadata.isHiring = true`** — `isHiringOrg` is computed **client-side** from Clerk:
   `isHiringOrg = Boolean(organization.publicMetadata?.isHiring)`, `organizationId =
   organization.publicMetadata?.eid` (`apps/web/src/hooks/useGetClerkOrganization.tsx:20-21`).

**⇒ M224 must extend Clerkenstein's fake API to emit `publicMetadata.isHiring = true`** for the hiring org — today
it emits `{eid}` only (`clerkenstein/clerk-backend/resources.go:38-47`). **This is a rext change, NOT a platform
edit.**

**The `is_hiring` blast radius (R5 enumerated — what changes vs stays identical under the flip):**
- The comparison surface **survives** — only relabeled "Results" (D1).
- **Two `isEnterprise` definitions diverge:** nav `isEnterprise = Boolean(organization)` (`template.tsx:90`) stays
  **TRUE** (enterprise nav still renders); billing `isEnterprise = !isHiringOrg && organizationId`
  (`FreeTrialContainer.tsx:29`) flips **FALSE** (hiring orgs are excluded from the Workforce free-trial —
  irrelevant to the comparison). The two definitions are **not** a bug; they answer different questions.
- Under `is_hiring` the nav also: relabels activity-dashboard → "Results"; trims the library to AI-Simulations;
  hides some member surfaces for non-admins; gates Workforce Intelligence **off**.

**Consequence:** R5 is enumerated (not merely flagged); M224 owns the Clerkenstein `publicMetadata.isHiring` wiring.

---

## D4 — BA-6: the 5 "job positions" = **5 real `SIMULATION_TYPE_HIRING` sims**; **no `job_position` replay**. (M223 Scope.In refinement — Fate-3.)

**Decision:** M223 seeds **real content, zero synth** by picking **5 of the 87 captured
`SIMULATION_TYPE_HIRING` sims** (published + public) as the org's "positions". The original plan's
**"replay `directus.job_position` rows" is MOOT** and dropped from M223's Scope.In.

**Evidence (the captured snapshot, not prod):**
- The cold-primed public snapshot carries **87 real `SIMULATION_TYPE_HIRING` sims** (published + public) — a rich
  pool; **≥5 usable, so BA-6 clears GO.** (The prod count was 715/127; the *captured* pool is what matters, and it
  is ample.)
- **`directus.job_position` has 0 rows in the snapshot** (the "443" was a **PROD** count, never captured). The
  comparison surface does **NOT** need `job_position` entities: `JobSimulation.jobPosition` is optional and unused
  by the scoreboard. So there is nothing to replay.

**Consequence:** M223's "snapshot extension to REPLAY `directus.job_position`" (its Scope.In #4) is **removed** — the
5 positions ARE 5 real HIRING sims. Recorded as a **Fate-3** refinement into M223's `overview.md` (Scope.In) and
here. R7 (snapshot starvation) is retired — the pool is rich.

---

## Adversarial review (M222 close, Phase 2c)

Scenarios considered against the `is_hiring` gate — the only executable surface this milestone shipped (the rest is
docs). Recorded per the close-milestone Phase-2c contract: the *scenario*, and whether the code already handles it.

1. **The manifest `omitempty` must actually drop a `false` `is_hiring` — else every existing preset's auditable
   manifest DRIFTS (a honesty-gate break).** If `manifest.Org.IsHiring` were a pointer, or the yaml tag lacked
   `omitempty`, a Workforce (non-hiring) org would serialize `is_hiring: false` and the checked-in preset manifests
   would no longer be byte-identical — the exact silent-drift class the honesty gate exists to catch. **Handled:**
   `manifest/hiring_test.go::TestBuildPopulation_ProjectsIsHiring` serializes a mixed population and asserts **exactly
   one** `is_hiring:` key in the whole document (only the hiring org emits it; the Workforce org's `false` is dropped
   by `omitempty`). A regression here fails at the serialization layer, not just the struct field.

2. **The dual-signal OR (`IsHiringOrg() = IsHiring || Narrative == "hiring"`) resolves a contradictory input
   `is_hiring: false` + `narrative: hiring` to TRUE.** An adversary sets the explicit flag false but leaves the
   narrative discriminator — the OR makes it a hiring org anyway. **Intentional:** either signal is sufficient (the
   comment "the two signals can never disagree" documents the design — the narrative is a strong signal, mirroring
   the `aiReadinessNarrative` sibling). Presets never write contradictory signals; the recognition is deliberately
   permissive so `narrative: hiring` alone (M223's path) suffices. Not a defect.

3. **Exact-match narrative discriminator is case/whitespace-sensitive** (`Narrative == "hiring"`, so `"Hiring"` or
   `"hiring "` would NOT be recognized). **Accepted / consistent:** identical to the established
   `aiReadinessNarrative` convention; narrative values are authored in controlled presets, not free user input. No
   new failure mode vs the existing discriminator pattern.

All three are covered by existing tests or are intentional design — no Fate-1 code change required.
