# M265 — Retro

**Milestone:** M265 "Prove it live" · iterative · `closed-on-gate` · 2026-08-16
**Release:** v2.9 "new alphabet" (02.90-new-alphabet) — the last of 7 milestones

## Summary

One tik. The gate — five clauses, all measured on cold bring-ups — fired, and on the way it caught
the defect the previous five milestones of this release could not: **a demo whose simulation library
rendered zero cards while every gate they had was green.**

The milestone's value was not the fixes. It was the discovery that v2.9 had **no measurement of a
rendered surface** anywhere in it, and that a canon swap produces a failure class — *hollow success*
— which row counts, liveness probes and unit tests are all structurally blind to.

## Incidents This Cycle

Six, all found and closed inside the milestone. Recorded in full because the pattern across them is
the reusable part: **five of the six are a verifier being wrong, not the system under test.**

1. **The content-realignment defect (the milestone's reason to exist).** 187 of 302 taxonomy node-ids
   embedded in replayed Directus content were retired by the canon swap. `skills[].name` is non-null,
   so one dead id nulled the entire `publicJobSimulations` list. Library: 0 cards. `/api/health`: 200.
2. **A hand-maintained column list, wrong within the hour of shipping.** Repaired its four columns,
   verified clean, exited 0 — and the next page load still failed on ids nested a level down.
3. **A swallowed error (`EXCEPTION WHEN others`) turned a `format()` arity bug into "nothing to do".**
   The remap reported success having changed nothing.
4. **An anti-vacuity rule with too broad a trigger** recorded a *successful* 55,116-row replay as a
   FAILED surface, twice — the second time because the probe's denominator had a different scope from
   the thing it qualified.
5. **A scan pattern narrower than its subject.** `[A-Z]{4,8}` matched 3,543 of 3,562 canon ids. It did
   not under-report; it reported **clean** while the app still failed.
6. **A wrong diagnosis, retracted inside the milestone.** `pt-assignment-assign` was called
   not-taxonomy-caused from two true measurements joined by a plausible story.

No P2 flakes. No regressions in unrelated suites.

## What Went Well

- **The gate did its job on the first run.** Clause 1 was green on the row counts and red on the
  rendered surface; only a clause that looks at what a user sees could tell those apart.
- **Fail-loud-by-default paid for itself repeatedly.** `realign` refusing to exit 0 on a residual is
  what turned incident 5 from a silent half-repair into a visible failure. The M217 pin guard, the
  hostlock and `fixture_nodeid_guard` each stopped a wrong run mid-milestone.
- **The retraction was recorded rather than overwritten.** The reasoning error — two true facts and a
  plausible join — is more reusable than the fix.
- **Both harden passes found real gaps**, and neither was in the code the iter had been fixing: the
  iter tested `realign` to 92.4 % and left the seams around it at zero.

## What Didn't

- **Four self-inflicted defects in one milestone is a lot**, and they share one root: I wrote
  verifiers whose scope was narrower than their subject, then trusted their green. The mitigation is
  cheap and I did not do it up front — *measure the verifier against the population it verifies*.
  One query (`[A-Z]{4,8}` vs 3,562 canon ids) would have caught incident 5 before it shipped.
- **The iter record was written retrospectively.** The work was executed and committed before
  `iter-01/` held anything, so the protocol's own artifact trailed the work rather than shaping it.
- **`state.md` had drifted three milestones behind** (`last_closed: M261` while M262–M264 were merged),
  so the contract surface disagreed with the branch for the whole milestone.
- **The sub-agent budget ran out mid-close** (300/300), so the lifecycle wrappers ran inline instead
  of in isolated sub-agents. Same skills, same order, no isolation layer — recorded here because the
  journal says so and a reader should not have to infer it.

## Carried Forward

All three are **Fate 3 → `/developer-kit:close-release`** — release-level by construction, not
cross-release, so none is an escape-hatch deferral.

1. **The claim-census ratchet is broken across six corpus files** (`dependency_map` 29→36,
   `external_services` 102→120, `backend` 30→32, `messenger` 13→16, `storage` 18→24, and the net-new
   `taxonomy-canon.md` at 37). Five predate M265. Re-baselining a ratchet is a target change, which is
   the user's call — doing it quietly at the tail of the last milestone is how a ratchet stops meaning
   anything (D-M265-6).
2. **Four pre-existing guard-family REDs** (`clone_drift`, `decommissioned_instruction`, `demo_knob`,
   `unreadable_repo_claim`), each flagging files M265 did not touch.
3. **Three archived-milestone scratchpads** (`work-m257`, `work-m257x`, `work-m258`) are sweep
   candidates. `work-m257x` alone holds hundreds of evidence artifacts, so the destructive sweep was
   **not** performed unilaterally.

## Metrics Delta

See `metrics.json`. Headline:

| | |
|---|---|
| Gate clauses met | **5 / 5** |
| Playthrough suite | **222 passed / 0 failed** (cold reset-to-seed) |
| Content refs repaired | **515 → 0 dangling** |
| `skill not found` in backend log | **258 → 0** |
| Tests added (harden) | **17** across 3 suites |
| Coverage, `cmd/stacksnap` | 75.8 % → **77.0 %** |
| Bugs fixed by harden | **2** (a false-GREEN guard regex; a fence pinned to a literal) |
| rext tags shipped | `v2.9.10-rext` → `v2.9.17-rext` (8) |
| Platform-repo edits | **0** |
