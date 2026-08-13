**Type:** tik · **Shape:** standard (protocol: `corpus/ops/demo/playthroughs.md` § *The iteration protocol*)

# iter-29 — the org-prepared Playthrough

## Phase A — the one measurement iter-28 routed

Driving the prepared variant's relabelled **`Start`**: it **CONFIRMS and ADVANCES**. `user_params.onboarding`
goes from `[{import}]` to `[{import},{role}]` — the platform persisted the role step — and the flow lands on the
**skills** screen: `Change Role` plus her declared role's **real taxonomy skills** (Business Requirement
Gathering, Data Analysis, Critical Thinking Fundamentals, … — Business Analyst's own `job_role_skills`), each
offered to keep or discard. So the curated flow's *"confirm the pre-filled role, refine the suggested skills"* is
that one click and its outcome.

## Phase B — the spec

`pt-onboarding-org-prepared`, `@pt-mutation: MUTATES`. Four accessors added to `OnboardingPage`
(`preparedSummaryName`, `preparedStartControl`, `changeRoleControl`, `suggestedSkill`).
`preparedStartControl` is held **apart** from `forwardControl()` on purpose — that accessor documents the
IMPORT step's `/^(Next|Import)$/`, and widening it to a third label that can never fire on the import path is
the dead-coverage shape iter-27's mutant P3 exposed.

## Phase C — mutants, and the one that PASSED

| mutant | outcome |
|---|---|
| **S1** delete the click | **RED** at "Change Role" |
| **S2** remove `onboarding: org_prepared` + RESEED | **RED** at liveness — the summary never renders, the import form returns |
| **S1c** delete the action AND the intermediates, leaving only the fresh-navigation read-back | **PASSED** |

**S1c is the finding (D115).** The read-back asserted *"the prepared summary is gone"* after
`goto('/onboarding')`. The reasoning was correct — once a `role` step exists `managerImport` cannot be true —
and the assertion was still **worthless**: `toHaveCount(0)` right after a navigation is satisfied by a page that
has not hydrated. **Removed, not weakened.** The honest repair needs a POSITIVE locator on the ROLE step, which
nothing has driven, so the persistence half is routed.

**This is the third variation on one theme:** iter-12 (a dead page satisfies every absence), iter-22
(`rows > 0` satisfied by *"No roles match your filters"*), and now **time** as the confounder — the page is
fine, it just is not there yet. *An absence assertion needs a companion that proves WHEN it was read, not only
WHERE.*

## Phase D — the gate

- **`197 passed` ×3 consecutive cold reset-to-seed, rc 0 each, 0 flake** (1.7 / 1.8 / 2.0 m).
- `ptreport`: **29/31 passing**, **0 failing**, **2** `unimplemented`, **0 `unimplementable`**.
- Controls **27 of 29** (13 self-declared + 14 via the control spec); `@pt-mutation` **MUTATES=11**.
- `--policy-check` green; 16 containers Up / 0 exited; `gofmt -l` clean; fixture restored (`99e2f315`).

## Close — 2026-07-30

**Outcome:** onboarding **3 of 5 → 4 of 5** — only `individual.UC1` remains. Controls 26/28 → **27/29**,
mutating 10 → **11**. And the milestone's signature defect was caught in this iter's *own* work for the third
time, by the standing mutant iter-27 routed for exactly that purpose.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — onboarding 4 of 5, one UC short — (2) triggered-tok: n — (3) re-scope: n
— (4) user-blocker: n — (5) cap-reached: n — 4 tiks this session (26, 27, 28, 29); the cap is 5 —
(6) protocol-stop: n — Outcome: continue
**Decisions:** D115 (the read-back that could not fail — removed, not weakened)
**Side-deliverables:** none.
**Routes carried forward:**
- **`ONBOARD-M256-prepared-persistence`** → the persistence half of `standard.UC2`. Needs ONE measurement: drive
  a fresh `/onboarding` after the role step is confirmed and record a POSITIVE locator on the **Role** screen
  (the component opens on `lastStep || Import`, and `lastStep` is `role`). Then the read-back becomes
  liveness-then-absence instead of absence-alone, and S1c goes RED.
- `ONBOARD-M256-import-path` → **`individual.UC1` is the last onboarding UC**, and the only one left with an
  unpriced capability: a **member-less user**, i.e. excluding a hero's slot from every per-index seeder that
  writes org-scoped rows. `standard.UC1` stays behind the measured CV-upload product defect.
**Lessons:**
1. **An absence assertion needs a companion that proves WHEN it was read, not only WHERE.** Liveness-before-
   absence (iter-12) is usually read as "prove the page is alive"; S1c shows the same rule has a *temporal*
   reading — prove the page has ARRIVED. `toHaveCount(0)` immediately after a navigation is the trap.
2. **The standing Q1 mutant earned its place immediately.** iter-27 routed *"delete the action and see whether
   anything fails"* as a check to run against every mutating Playthrough; run here on brand-new work it found a
   green-but-worthless assertion in the very first spec it was applied to. One edit, one run.
3. **Removing an assertion can be the honest outcome.** The pressure is to weaken it into something that passes;
   the right move was to delete it and name what would be needed to do it properly.
