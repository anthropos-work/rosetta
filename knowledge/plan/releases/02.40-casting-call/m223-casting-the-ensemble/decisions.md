# M223 — Decisions

_Implementation choices with rationale. One entry per decision; cite the code/doc it binds._

---

## D1 — The insights substrate needs NO net-new Casbin grant: the hiring org's admins inherit `org:feature:insights` from the GLOBAL `p3` admin policy via their standard `admin` g2 grant.

**Decision:** the `OrgFeatureInsights` gate the recruiter scoreboard sits behind
(`resolver_queries.go:1089`) is satisfied **for free** for the hiring org — no hiring-specific grant is
written. The seeder just needs the org's 5 admins to carry the standard `admin` g2 grant (which
`users.go` already writes for `role == "admin"` members).

**Evidence (traced through the running Sentinel + app clones, READ-ONLY):**
- The feature check maps to enforce-context **`3`** → matcher
  `m3 = g2(r3.org, r3.sub, p3.sub_role) && (r3.org != r3.sub) && ('default'==p3.org || r3.org==p3.org)
  && r3.feat==p3.feat` (`sentinel/internal/authorization/casbin.go`,
  `enforcer_conversions.go:orgCheckFeatureRequestToEnforceRequest`).
- `init_policy.sql:54` seeds a **GLOBAL** (`org='default'`) policy row
  `('p3','default','admin','org:feature:insights','','','')` — so ANY member with a `g2(org, user,
  'admin')` grant in ANY org clears the insights gate.
- `roleForIndex` (users.go) makes the first `RoleMix.Admin` fraction of a story's members `admin`; each
  gets a `CasbinGrant{Role:"admin"}`. With `role_mix ≈ 0.1 admin`, the 5 hiring admins are `admin` → they
  inherit insights.

**Consequence:** S4's "replicate the `OrgFeatureInsights` Casbin grant" is realized by ENSURING the hiring
org has `admin` members (via the role mix) — not by a new grant. A unit test asserts the funnel/story
produces ≥1 `admin`-role candidate-cohort admin. `p2` also grants `admin` the `admin-vs-candidate`
action rows (`init_policy.sql:36-40`), so an admin reading a candidate's sessions is not action-403'd.

---

## D2 — There is NO platform "hiring-config" table: the recruiter scoreboard AND its sim list both derive from `local_jobsimulation_sessions`. So the 5 positions materialize purely from the funnel's mirror rows.

**Decision:** the "AI-Simulations" list (`InsightsByJobSimulations`) and the per-sim scoreboard
(`InsightsJobSimulationByMemberships`) BOTH read only `m.ent.LocalJobsimulationSession`
(`intelligence.go:1527,1692`), org-scoped by `organization_id`. A sim appears as a "position" **because
the org has candidate mirror-sessions on it** — there is no config/registry table (unlike AI-readiness's
`ai_readiness_sims`). So the 5 positions are made real by S4's funnel writing candidate mirror rows; no
config-table write is required for the surface to render.

**Consequence (shapes S2):** the "HiringConfigSeeder / AI-readiness-config analog" collapses to the
**resolution layer** (the type-aware hiring-sim reader + the shared `hiringSimRefs`) plus writing the
org's 5 positions as `organization_sim_invitation_links` (D3) — NOT a scoreboard-read config table (none
exists). The scoreboard's disjoint-cohort property is guaranteed by ORG-SCOPING + the fact that the ONLY
writers of `local_jobsimulation_sessions` are the PersonaSeeder (heroes — the hiring story has none) and
S4's funnel (org-scoped to the hiring org).

---

## D3 — S2's "HiringConfigSeeder" writes the org's 5 positions as `organization_sim_invitation_links` (folding in S6); S6 is DONE, not skipped.

**Decision:** the config seeder writes exactly one `organization_sim_invitation_links` row per shared
hiring sim (5 per hiring org) — the faithful "the recruiter created 5 positions" analog of
`AIReadinessConfigSeeder` writing `ai_readiness_sims`. The table's partial `UNIQUE(simulation_id,
organization_id) WHERE deleted_at IS NULL` (app migrate/schema.go:2443) makes this exactly-5, no balloon;
`options` is a simple `{"voiceOnly":false,"enableRecording":false}` JSON
(`repository.CreateOrganizationAssignmentInput`), `token` a deterministic 22-char URL-safe string,
`created_by` an admin membership. **S6 (invitation links) is therefore SATISFIED by S2, not deferred** —
the comparison reads sessions not links, but the links make the org's positions real + auditable.

**Consequence:** S2 and S6 are one seeder. Idempotent via deterministic id (ON CONFLICT id).

---

## D4 — The 5 positions are 5 REAL captured `SIMULATION_TYPE_HIRING` sims via a type-aware reader; disjoint-reserved from the generic pool. (S2/S3 — S3's job_position replay stays DROPPED per M222 D4.)

**Decision:** a dedicated `readHiringSimPool` reads
`directus.simulations WHERE type='SIMULATION_TYPE_HIRING'` (the captured public snapshot already applied
`private=false AND tenant_id IS NULL AND status='published'`), ORDER BY id — the type-aware precedent of
`readAIReadinessSkillPool`. The current `contentref` sims pool is type-BLIND
(`SELECT id ... ORDER BY id LIMIT 50`), which cannot distinguish a hiring sim. `hiringSimRefs()` takes the
first `reservedHiringSimRefs` (5) — SHARED by the config seeder (positions) and the funnel (session
sim_ids) so they co-derive (the `aiReadinessSimRefs` pattern). The 5 chosen ids are EXCLUDED from the
generic sims pool's pickers (the M219 R-3 reserved-disjoint fix) so a generic activity session can never
reference a hiring position. `directus.simulations.type` carries `SIMULATION_TYPE_HIRING` (cms
`jobsimulation.go:741` oneof; M222 measured **87** captured public hiring sims). **S3 (`directus.job_position`
replay) stays DROPPED** — 0 rows captured, the scoreboard doesn't read `job_position` (M222 D4).

---

## D5 — The hiring funnel writes the `local_jobsimulation_sessions` MIRROR pair (like the PersonaSeeder), NOT the AI-readiness-funnel shape (which skips the mirror). The M219 render-gate trap.

**Decision:** S4 reuses the PersonaSeeder's session-pair shape (`persona.go:267-288` — a
`jobsimulation.sessions` row + its `public.local_jobsimulation_sessions` mirror row with the `score`),
NOT the `AIReadinessFunnelSeeder` shape. The AI-readiness funnel deliberately does **not** write the
mirror (it scores from frozen `ai_readiness_snapshots`), but the recruiter scoreboard's score comes from
`local_jobsimulation_sessions.score` (M222 D2). Seeding only `jobsimulation.sessions` renders an EMPTY
comparison — the exact M219/M222 render-gate trap. The funnel writes the `#1 (mirror) + #2 (session)`
pair per (candidate × taken-sim), org-scoped to the hiring org, `sim_type='SIMULATION_TYPE_HIRING'`,
G14-valid enums.

**The funnel shape (differentiated, non-degenerate — the M219 anti-flat-arc lesson):** each candidate
gets a base aptitude spread across ~[30,95] (deterministic hash), + a small per-sim jitter, so within a
sim the ~40 candidates are RANKABLE (not 45 identical scores) and a strong candidate is consistently
strong across the 5. `completition_status = passed` when score ≥ 60 else `failed`. **MOST candidates take
all 5** (`hiringAssignedOnlyShare ≈ 0.1` of candidates are assigned-not-taken → no mirror rows → absent
from the ranked list — the 2nd candidate hero's future state, M224). Closure is UNAFFECTED: the funnel
writes zero skill refs (`measure-closure` only checks `user_skills`/`user_skill_evidences`/
`validation_attempt_skill_results`).

## Adversarial review (M223 close, 2026-07-16)

**Scenario — the captured snapshot has 1–4 HIRING sims (a starved pool).** `hiringSimRefs` takes the first 5
of the type-filtered `SIMULATION_TYPE_HIRING` pool; `HiringConfigSeeder` guards `len(positions)==0` (honest
skip — the M20 graceful-degradation contract: no hiring content ⇒ structural-only, never abort/pad). The
dangerous middle is **1–4**: the seeder would write fewer than 5 shared positions and the funnel would assess
candidates on fewer than 5 — a silent incomplete comparison.
**Verdict — handled by compensating controls, not a live risk:**
1. M222 empirically measured **87** captured HIRING sims (published+public) — the pool is not starved today.
2. The reader is **type-filtered + reserved-disjoint** — it can never pad from the generic pool (the M219
   anti-pad discipline), so a short pool surfaces as *fewer real positions*, never as generic sessions
   masquerading as hiring assessments.
3. **The downstream render gate catches it loud:** M224/M226's exit gate requires **≥40 comparable rows per
   EACH of the 5 sims** — a <5-position seed fails that gate at verification, on a cold reset-to-seed, before
   any ship. So a starved re-capture cannot silently reach a demo.
Recorded, not code-guarded: adding a fatal `<5` abort would regress the M20 graceful-degradation contract for
a genuinely cold box; the render gate is the correct, already-planned control.
