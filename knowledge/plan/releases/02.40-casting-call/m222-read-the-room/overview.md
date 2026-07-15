---
milestone: M222
slug: read-the-room
version: v2.4 "casting call"
milestone_shape: section
status: archived
created: 2026-07-15
last_updated: 2026-07-15
depends_on: none
delivers: corpus/services/hiring.md (BLIND AREA — the hiring READ-model BA-1/BA-2 + the isEnterprise divergence + the isHiringOrg derivation) + the is_hiring gate + the narrative:hiring discriminator + a render-proof GO/NO-GO decision + the exact seeder-output contract M223/M224 build against; registrations pending in seeding-spec.md/stories-spec.md/README.md
---

# M222 — Read the room

## Goal
Author the missing hiring-model doc, and **prove by rendering** — on a throwaway hand-seed against a live dockerized
`apps/web` — that the recruiter comparison surface (a) exists in the demo-servable app and (b) renders a comparable
score from seedable data. Land the `is_hiring` gate + the `narrative: hiring` discriminator. This milestone builds no
seed and no cockpit — it makes the foundation real and **proves the path before the risky work commits**.

## Why section
The deliverables ARE enumerable up front — a corpus doc, a set of render-probes, a go/no-go decision, and a one-value
gate thread. **What is uncertain is the *result* (which then re-scopes M224), not the task list.** This is the "clean
stage" analog: make the foundation real + prove the path before committing the risky seeder/render work.

## Why this is a HARD go/no-go barrier

Two blind areas gate the whole release, and neither has a doc or code anchor today. Each needs a `Delivers →` doc
line or a render-proof **before** any code commits against it:

- **BA-1 — the hiring READ-model.** No corpus doc. The exact read-path that fills the activity-dashboard per-sim
  candidate list with a **comparable score** was not traced end-to-end. **Unproven whether the score/best-session
  comes purely from `jobsimulation.sessions.score` (which the seeder writes) or ALSO needs a
  `validation_attempt_results`/evaluation row per session** the current seeder may not populate — a potential extra
  seeding surface (data-only, not a platform edit).
- **BA-2 — does the comparison surface render for an `is_hiring` org?** Many `apps/web` pages hardcode
  `isHiringOrg={false}`; the exact set of surfaces that change vs stay identical under `isHiring` was not enumerated.
  Also unresolved: the `isEnterprise=Boolean(organization)` (nav) vs `isEnterprise=!isHiringOrg && organizationId`
  (billing) **divergence** — the two definitions differ and the blast radius is unverified (R5).
- **BA-3 — is the demo-servable surface actually in `apps/web`, or only `apps/hiring`?** Strong evidence it is
  `apps/web` (route files present; `Dockerfile.dev` builds `@anthropos/web-app`), but a full route-by-route diff
  `apps/hiring` vs `apps/web` was not done. **This is the make-or-break go/no-go.** If `apps/hiring`-only, showing it
  = containerizing a Vercel-only app = large net-new + likely a platform edit → **ESCALATE. Do not proceed.**
- **BA-6 — does the cold-primed public snapshot contain ≥5 usable HIRING sim blueprints?** 715 HIRING sims / 127
  public exist *in prod*, but no anchor confirms the **captured** snapshot pool carries ≥5 HIRING-typed sims tied to
  job positions vs being TRAINING-dominated. Probe the captured pool.

## Scope

### In
1. **Author `corpus/services/hiring.md`** (the model): `is_hiring`, the `candidate` role, hiring sims ↔
   `job_position`, `jobsimulation.sessions` HIRING, the read-path for the comparison surface, the `isHiringOrg`
   Clerk-`publicMetadata` derivation, the `isEnterprise` divergence blast radius.
2. **Add the blueprint `Org.IsHiring` field** + the `narrative: hiring` discriminator (the `OrgSpec` has no hiring
   notion today — just `{Name, Slug, Industry, Narrative}`).
3. **Thread `is_hiring=true` into `org.go`** — read `is_hiring` from the spec instead of the hardcoded `false` at
   `stack-seeding/seeders/org.go`. Small, but **it is the load-bearing gate** (D-DESIGN-1: genuine hiring org, so we
   DO flip it; M222 de-risks the flip's blast radius rather than dodging it).
4. **Reserve the hiring-org deterministic `OrgID`.**
5. **Render-probe answers to BA-1/BA-2/BA-3/BA-6** and record a **GO/NO-GO + the exact seeder-output contract**
   M223/M224 build against — critically: *does the score render from `sessions.score` alone, or does it need a
   `validation_*`/eval row per session?*

### Out
- The full 50-person seed · the assessment funnel · cockpit heroes · any latency work. (All M223/M224.)

## Depends on
Nothing — the release entry point. **The v2.3 `billion` demo LEFT LIVE is a free render substrate** for the probes.

## KB dependencies
- `kb-ant-business` hiring.md / ai-interview.md · `corpus/ops/demo/stories-spec.md` (the 7-table chain)
- `corpus/ops/snapshot-spec.md` · v2.3's `corpus/ops/demo/demopatch-spec.md`

## Delivers → knowledge/corpus
**`corpus/services/hiring.md`** — **BLIND AREA (BA-1/BA-2).** Grep of `corpus/**` for a hiring model returns no doc
anchor; hiring exists only as a "distinct-frontend" line in `next-web-app.md` + the business KB. This doc must be
authored before the seeder codes against the contract. Registrations pending in
`seeding-spec.md`/`stories-spec.md`/`README.md` (the 4th story + the `is_hiring` gate).

## Demo-patch?
**No patch built here** — but M222 *decides whether one is needed downstream.* If the render-probe shows the surface
hard-gates on `isHiringOrg`, or the score demands an un-seedable resolver path, that becomes an **M224 demo-patch**
(the M219 precedent). If BA-3 comes back "`apps/hiring`-only," **escalate — do not proceed.**

## Risks carried
- **R2 (blocks-release)** — the comparison view may be `apps/hiring`-only. **This milestone's render-probe is the
  make-or-break go/no-go.** Escalate if so.
- **R5 (degrades-quality)** — the `isHiring` blast radius (two `isEnterprise` definitions). M222 **enumerates** the
  surfaces that change vs stay identical under `is_hiring=true`.
- **R7 (degrades-quality)** — snapshot content starvation (BA-6). M222 probes the captured pool.
