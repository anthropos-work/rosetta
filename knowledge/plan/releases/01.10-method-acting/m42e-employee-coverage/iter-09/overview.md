---
iter: 09
milestone: M42e
iteration_type: tik
iter_shape: production-fix
status: closed-fixed
created: 2026-06-25
---

# iter-09 — honest sim-start fix: seed the job-simulations entitlement + remove the dishonest skip

## Why this iter (the gate-honesty correction)
iter-08 closed `(failing=0, escapes=0)` and graded gate-MET — but an adversarial **gate-honesty audit**
(`.agentspace/scratch/work-m42e/verify-gate.md`) INVALIDATED that close. 3/4 lenses GREEN (re-capture provably
safe, state/supply-chain GREEN, result independently reproduced), but the **skip/allow-rule-honesty lens is
YELLOW**: the `SIM_START_SKIP` rule (iter-08 D2) rests on a **FALSE justification**. The `/sim/<slug>/start`
pages were excluded as a "runtime launch surface" (like the genuinely-runtime sim-RESULT pages), but they are
actually empty because **Maya's demo org lacks the `FEATURE_JOB_SIMULATIONS` entitlement** — the page renders a
**deny modal** (`AISimulationOrgMemberCannotStartModal`, empty `<main>`). That is a **SEEDABLE** failure
dishonestly scoped out. The TRUE gate with `/start` IN SCOPE is **failing-1**, not (0,0).

## Active strategy reference
**TOK-01** (`sweep-then-route-by-leverage`). This tik routes the sim-start cluster to its CORRECT fix surface
(`stack-seeding`, per the protocol's "seedable structural row vs runtime-computed artifact" distinction) —
correcting iter-08's mis-triage of it as a crawl-scope-out.

## Cluster / target identified
The single dishonest exclusion: the per-sim `/start` launch surface. iter-08 D2 classified it runtime-computed;
the audit + a code read prove it is **entitlement-gated**, which IS seedable under the zero-platform-edit line.

## Entitlement diagnosis (the exact data surface — done at iter open, read-only)
- `start page.tsx` L44-73: `canStartAsOrganizationMember` = `userMembership.organizationFeatures.some(f === FeatureJobSimulations)`.
  When `false` → renders the deny modal (empty `<main>`). When the feature is present → renders the real
  `AISimulationStartWithoutSession` launch UI.
- The `organizationFeatures` GraphQL field resolves (app `resolver_organizations.go:149`) via
  `OrgManager.IsMemberAllowedToUseFeature` → `authMan.OrgMembershipsAllowedToUseFeature` → **Sentinel RPC** →
  `GetFilteredNamedGroupingPolicy("g3", 0, org)` then `slices.Contains(membershipIDs, membership.ID)`.
- So the gate is **NOT** `app.organization_features` (that table is 0-rows even in normal operation — a
  red-herring the verdict named by symptom). The real gate is a **Sentinel Casbin grouping policy `g3`**:
  `casbin.go` model declares `g3 = _, _` → stored row `p_type='g3', v0=<org_id>, v1=<membership_id>`,
  `v2..v5=''`. The app grants it via `OrgAllowUserToUseFeature` → `AddNamedGroupingPolicy("g3", org, membership)`.
- Live demo-3 confirm: `sentinel.casbin_rules` has **0 `g3` rows** (g2=341, p/p2/p3/p5 present, g3=0). Maya Chen
  (`maya-thriving`) = membership `957cc282-…`, org `22222222-…` (Cervato Systems).

## Hypothesis
Seeding a `g3(org, membership)` grant for every seeded member (mirroring the existing per-member g2 grant in
`users.go` + the demo identity in `identity.go`) makes `userMembership.organizationFeatures` carry
`FeatureJobSimulations` → `/sim/<slug>/start` renders the real `AISimulationStartWithoutSession` launch UI
(non-empty `<main>`) instead of the deny modal. The start PAGE only needs to RENDER its launch UI to pass the
coverage assertion (not run a full simulation).

## Phase plan
- **Phase B/C (fix in rext, zero platform edit):** add a g3 feature-grant helper alongside the existing g2
  helpers in `identity.go` (reuse `resolveCasbinTable` + the idempotent WHERE-NOT-EXISTS guard +
  ::text-pinned params), and emit a g3 grant per membership in `users.go` (the member population) and in
  `identity.go` (the demo identity membership). Then **remove `SIM_START_SKIP`** from `coverage.spec.ts` so
  `/start` is back IN SCOPE and scored. Tests: rebuild stackseed + `go test ./...` on stack-seeding;
  `playwright test --list` compile-check on the harness.
- **Phase C re-apply:** re-run the seeder against live demo-3 (re-seed) so g3 grants land; verify g3 rows > 0
  + Maya's membership present + the start page renders via a probe.
- **Phase D re-sweep:** streaming foreground full frontier-EXHAUSTED re-sweep as `employee/maya-thriving` at
  the raised cap; quote the HONEST residual with `/start` scored.
- **Phase E close:** grade on whether the start cluster cleared + the gate is HONESTLY (0,0) with `/start` IN
  SCOPE.

## Expected lift
With `/start` back in scope, the pre-iter HONEST residual is `(failing≈1+, escapes=0)` (the start pages were
the dishonestly-hidden failures). Seeding the entitlement should render them → drive to a TRUE `(0,0)`.

## Escalation conditions
- If seeding g3 does NOT render the start page (some additional gate beyond the entitlement) → DIAGNOSE per the
  protocol's DOM+network+log probe; if the residual is a genuine runtime-computed surface (not the deny modal)
  re-document the skip with its TRUE justification + route forward. Do NOT re-instate a dishonest skip.
- If a 100%-blocker can ONLY be closed by a platform change → re-scope-trigger (the zero-edit line). Not
  expected — the entitlement is a Sentinel-policy seed, fully in-rext.

## Acceptable close-no-lift outcomes
If the probe proves seeding the entitlement makes the start page render BUT the re-sweep surfaces a NEW honest
residual elsewhere (a deeper crawl reaching a real gap) → that is real progress (the dishonest skip is gone);
route the new residual forward, grade on the start cluster's clearance.
