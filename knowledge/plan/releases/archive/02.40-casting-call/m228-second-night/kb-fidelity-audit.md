---
title: "KB Fidelity Audit — M228 second-night"
date: 2026-07-17
scope: milestone:M228
invoked-by: build-mstone-iters (Phase 0b)
---

## Verdict
**GREEN**

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Remote reach / tailscale serve (incl. apps/hiring front) | `corpus/ops/demo/tailscale-serve.md` | rext `demo-stack/gen_tailscale_serve.py` `UI_BROWSER_FACING` | PAIRED |
| Latency gate (recruiter 3rd vantage) | `corpus/ops/demo/latency-budget.md` | rext `stack-verify/e2e/run-latency.sh` (`recruiter`) | PAIRED |
| Bring-up verification / autoverify | `corpus/ops/verification.md` | rext `stack-verify/live/autoverify.sh` | PAIRED |
| Tooling safety contract | `corpus/ops/safety.md` | rext `stack-*` guards | PAIRED |
| Hiring read-model + M227 corrections | `corpus/services/hiring.md` | rext `stack-seeding/seeders/{hiring_funnel,users,userprofile,gender,hiring_scope}.go` | PAIRED |
| Coverage protocol / render gate floor | `corpus/ops/demo/coverage-protocol.md` | rext `stack-verify/e2e/tests/render-hiring-comparison.spec.ts` | PAIRED |
| Demo-patch mechanism (4 hiring patches) | `corpus/ops/demo/demopatch-spec.md` | rext `demo-stack/patches/*.yaml` | PAIRED |
| Seeding blueprint | `corpus/ops/seeding-spec.md` | rext `stack-seeding/` | PAIRED |

## Fidelity Findings (the M227-correction delta — the only audit-relevant change vs the M226 GREEN)

All 4 M227 corrections are ALIGNED in BOTH corpus and the `casting-call-m227-sections` rext tooling:

1. **fix#1 hiring-only content** — hiring.md:238-243 (`hiring_scope.go IsHiringOrg()`, #M227-D1) ↔ rext
   `stack-seeding/.../hiring_scope.go` + generic-seeder skip. **ALIGNED.**
2. **fix#2 external candidate emails** — hiring.md:241 (role-keyed external domain, admins keep org domain, #M227-D2) ↔
   rext `userprofile.go` `externalCandidateDomains = [gmail.com, outlook.com, proton.me, icloud.com, …]`,
   `emailDomainForMember` keyed on role. **ALIGNED.**
3. **fix#3 1-sim/candidate + retuned floor** — hiring.md:228-232 (`≥40 → ≥6`, `hiringComparableFloor` /
   `RENDER_GATE_FLOOR`, ~8/position) ↔ rext `hiring_funnel.go:127,156-157` (even round-robin
   `assessedOrdinal % len(positions)`) + `render-hiring-comparison.spec.ts:56` (`RENDER_GATE_FLOOR ?? '6'`) +
   `run-hiring-render.sh:18` ("default floor 6"). **ALIGNED.**
4. **fix#4 gender-matched avatars** — hiring.md believability section ↔ rext `gender.go` (curated first-name→gender
   dict, GenderUnknown → honest gender-blind fallback) + `users.go:180`. **ALIGNED.**

The M226 shared-infra fold-ins are present: the recruiter 3rd vantage + the Finding-1 serve prerequisite in
`latency-budget.md:28-38,92`; the apps/hiring `3001+off` serve front in `tailscale-serve.md:420-437`. All
`hiring.md` cross-references resolve.

## Completeness Gaps
None load-bearing. (Note for the measurement phase, not a doc gap: the render harness gates C2 on rows/floor; the
external-email + matched-avatar + hiring-only render *checks* are new live observations M228 will make — the harness
surfaces candidate names/emails/avatars in its capture, so these are observable at measurement time.)

## Applied Fixes
None needed — docs and code already aligned (M227 close blended the fix mechanisms into corpus with #M227-D refs;
M226 folded the recruiter vantage + hiring serve front).

## Open Items (require user decision)
None.

## Gate Result
**GREEN — proceed.** Every topic PAIRED, every M227-correction claim ALIGNED across corpus + rext, no blind areas,
no stale load-bearing claims, cross-refs resolve. The bootstrap tok may author strategy against verified docs.
