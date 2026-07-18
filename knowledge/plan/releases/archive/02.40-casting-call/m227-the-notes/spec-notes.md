# M227 — Spec notes

Technical detail, code anchors, thresholds. All rext code lives in
`.agentspace/rosetta-extensions/stack-seeding/` unless noted.

## Fix #1 — hiring-only content

**Surface traced (the recruiter's main hiring view):**
`apps/hiring/.../enterprise/activity-dashboard/@tabs/ai-simulations/page.tsx` → `InsightsByJobSimulationsContainer`
→ `useGetInsightsByJobSimulations` → app resolver `IntelligenceManager.InsightsByJobSimulations`
(`app/internal/organization/intelligence.go:1472`). It queries `LocalJobsimulationSession` (the MIRROR) filtered by
`user.IDIn(org members)` + `OrganizationID(org)`, grouped by `jobsimulation_id`. Each distinct sim shows with its
REAL Directus type (resolved via `resolveJobSimulationIdsByNameAndType`). So the list = every distinct sim that has
a mirror row for an org member.

**Mirror writers for the hiring org (only two):**
- `HiringFunnelSeeder` — HIRING-typed only ✓ (the 5 positions).
- `PersonaSeeder.seedVerifiedSkill` (`persona.go:213+`) — writes the candidate heroes Cara/Cody (`verified: 8` /
  `verified: 2` in `presets/stories.seed.yaml`) a verified-skill chain whose `simID = linkedRefDistinct(refs.sims, …)`
  draws from the GENERIC (type-blind, hiring-positions-withheld) pool. Those generic sims (training/assessment/…)
  surface in the recruiter list. **This is the defect.**

**Fix:** generic workforce activity/simulation seeders SKIP hiring orgs (`if st.IsHiringOrg() { continue }` in the
per-story loop). A hiring org's entire sim footprint is then the HiringFunnel (hiring-only). Candidate heroes'
usable profile = the HiringFunnel `/home` assignment (`appendHeroAssignment`) + scored positions; `/profile` (Skill
Spotlight, where the verified chain would show) is admin-gated → a candidate is redirected to `/home` and can never
see it, so dropping the verified chain for hiring candidates loses nothing visible while removing the pollution.

**Seeders to gate (they iterate `s.EffectiveStories()` and write sim/skillpath/assignment/activity footprint):**
`PersonaSeeder`, `JobsimSessionsSeeder`, `SkillpathSessionsSeeder`, `AssignmentsSeeder`, `ActivitySeeder`,
`HeroActivitySeeder`. The hiring-specific seeders (`HiringConfigSeeder`, `HiringFunnelSeeder`) already own the hiring
org. The M36 workforce-intelligence dashboard seeders (MembershipSkills/Tags/TargetRoles/Succession/Feedback/
PopulationEvidence/MemberLanguages/Certificates/Projects) render into Workforce-Intelligence surfaces that are GATED
OFF for `is_hiring` orgs (`hiring.md:165`), so their data isn't visible for the hiring org; gate them too ONLY if
they write jobsimulation/mirror rows (verify each) — otherwise leave to minimize risk. Decision recorded in
decisions.md.

**Shared helper:** add `func skipStoryForGenericActivity(st blueprint.ResolvedStory) bool { return st.IsHiringOrg() }`
(or inline the guard) — one source, documented, so the rule is discoverable.

## Fix #2 — external candidate emails

`users.go` currently: `domain := storyEmailDomainFor(st)` (per-story), then `email = emailFor(first, last, domain, i)`
for both hero + fill members. Role-blind.

**Fix:** compute the domain per-MEMBER from the resolved `role`:
- `role == "candidate"` → a deterministic external consumer domain from a bounded bank
  (`gmail.com`, `outlook.com`, `proton.me`, `icloud.com`, `hotmail.com`, `gmx.com`, …), keyed by `hash(uid) % len`.
- else → `storyEmailDomainFor(st)` (org domain, unchanged for admins/members).

New helper in `userprofile.go`: `externalCandidateDomain(seed string) string` + `emailDomainForMember(st, role, seed)`.
Applies to BOTH the hero branch and the fill branch in `users.go` (both call `emailFor`). The heroes Cara/Cody are
role=candidate → external. Rae (admin) → org domain.

**Consistency (Clerkenstein roster == seeded email == login):**
- `roster.go` exports the roster email. Verify it derives from the SAME `emailFor(...)`/domain path (single-source) so
  the fake-Clerk login email == the seeded `public.users.email`. If roster re-derives independently, route it through
  the same per-member domain logic.
- The preset `login:` for Cara (`cara@meridian-talent.com`) / Cody (`cody@meridian-talent.com`) / (Rae stays org).
  The `login:` is a cockpit display/seat hint; the ACTUAL login identity must match the seeded email. Check how the
  cockpit + Clerkenstein resolve a hero's login email (persona `login:` vs the derived `emailFor`). If the login uses
  the preset `login:` literal, update Cara/Cody's `login:` to their external address OR make the login derive from the
  seeded email. **The invariant to preserve: the email the cockpit logs in with == the email in `public.users` ==
  the Clerkenstein roster email.** Trace this end-to-end during implementation (roster.go + cockpit.go + clerkenstein).

## Fix #3 — 1 sim per candidate + gate retune

`hiring_funnel.go:seedHiringOrgFunnel` — currently an ASSESSED candidate loops `for pi, simID := range positions`
(all 5). Change to ONE position per assessed candidate:
- `posIdx := int(hashInt(fmt.Sprintf("%s:hiring-position:%d", prefix, i)) % uint64(len(positions)))` → even split.
- Write the co-written session+mirror pair for `positions[posIdx]` only.
- The candidate HERO (Cara, assessed) must land on a KNOWN position so her `appendHeroAssignment(positions[0], …)`
  and her scored session agree — pin an assessed hero to `posIdx = 0` (or make appendHeroAssignment use `positions[posIdx]`).
  Keep the assessed hero on ONE position consistently.
- Keep the aptitude/jitter score model (single position now → single score; still rankable across the ~8 peers).
- Keep the mirror pair (both `jobsimulation.sessions` + `local_jobsimulation_sessions`) — the M219 trap.
- Closure stays green (the funnel writes zero skill/role refs — unchanged).

**Per-position count:** ~40 assessed candidates (45 × (1 − 0.10 assigned-only)) split across 5 → ~8/position, min
possibly 6-7 depending on the hash split. **COMPUTE the exact deterministic per-position distribution** (a unit test
enumerating the 45 candidates) before setting the floor.

**Gate retune — set the floor = (min per-position count) − small margin.** Retune EVERYWHERE `≥40` is asserted:
- `contentref.go` `reservedHiringSimRefs` stays 5 (positions count, unchanged); the FLOOR is a separate constant.
- The render-probe `RENDER_GATE_FLOOR` default (`stack-verify/e2e/run-hiring-render.sh` + the spec).
- M224 `decisions.md` GATE-DECISION D1 (the ≥40 interpretation).
- M226 `overview.md` exit_gate condition (2) + the new M228 `overview.md` exit gate (2) in `roadmap.md`.
- The M225 coverage manifest (`reservedHiringSimRefs`/floor) + the hiring playthrough assertions.
- `corpus/services/hiring.md` (the "≥40 comparable" claim + "MOST take all 5" → "each on 1").
Grep for `40` / `≥40` / `all 5` / `reservedHiringSimRefs` across rext + corpus + knowledge to find every assertion.

## Fix #4 — gender-consistent avatars

`avatar.go:photoAvatarDataURI(seed)` → `assets.AvatarJPEG(hashInt("avatar:photo:"+seed))`. The 12 faces (visually
categorized): **female = face-01,03,05,07,10 (5); male = face-02,04,06,08,09,11,12 (7).**

**Fix:**
1. `assets/avatars.go`: add a gender partition — `AvatarIndicesByGender(g Gender) []int` returning the sorted indices
   (0-based, matching `avatarNames` order) for female {0,2,4,6,9} / male {1,3,5,7,8,10,11}. (Indices are 0-based;
   face-01=idx0.) Keep a package-level constant map so the partition is auditable + a test asserts each face's gender.
2. `avatar.go`: `photoAvatarDataURIForName(seed, firstName string) string` — infer gender from `firstName`, pick a face
   from that gender's subset by `hash(seed) % len(subset)`. Fallback (unknown name) → the full pool by hash (current
   behavior; never empty). Keep `photoAvatarDataURI(seed)` for callers with no name (or route them through the new fn).
3. Gender inference: `inferGender(firstName string) Gender` — a curated first-name→gender dictionary covering the 40
   fixed `firstNames`, the heroes, and common professional/generated names, + a light heuristic fallback (unknown →
   `GenderUnknown` → full-pool pick). Build the dictionary from the fixed bank first (deterministic coverage of the
   curated population), then extend with common names. A `GenderUnknown` name keeps the current gender-blind pick
   (honest degradation, never a wrong-gender forced pick).
4. Thread `firstName` at both call sites: `users.go` (curated + fill → `first`) and `generated_batch.go` (generated →
   `first` from `splitName(env.Name)`). Both already have `first` in scope.
5. Determinism/idempotency preserved: same (seed, name) → same face; re-seed byte-identical.

## Fix #5 — local re-prove
- Bring up a fresh LOCAL demo (`up-injected.sh` / `/demo-up N`, offset ports) consuming the M227 rext tag; cold
  reset-to-seed. Cached images (fast). NEVER billion.
- Assert from THIS Mac against the LOCAL offset ports: recruiter comparison ≥ floor per each of 5 positions; 1
  sim/candidate; hiring-only sims; candidates @external; avatars match names.
- Run the hiring coverage sweep + the hiring playthrough GREEN.

## Pre-flight audits — Hiring-only content (section 1)
**Phase 0b KB-fidelity (audit-kb-fidelity --milestone=M227): GREEN** (2026-07-17). Report:
`knowledge/plan/releases/archive/02.40-casting-call/m227-the-notes/kb-fidelity-audit.md`. All 6 topics PAIRED, all audited
claims ALIGNED with current code; no blind areas, no stale load-bearing claims. The `≥40 / all-5 / gender-blind
avatar / org-domain email` descriptions correctly document TODAY's behavior — M227 revises them as deliverables
(`overview.md` `Delivers →`). Triples recorded in the Topic Inventory table of the report. Reused for all sections
this session (same subsystem: stack-seeding + the hiring probe/coverage/playthrough) per the audit-reuse rule unless
a knowledge doc load-bearing for the milestone changes.
