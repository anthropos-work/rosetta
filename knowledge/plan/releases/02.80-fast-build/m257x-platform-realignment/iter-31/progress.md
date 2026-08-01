---
milestone: M257x
iter: 31
---

**Type:** tik (under `TOK-01: instrument first, then follow`)

# iter-31 — the org that had one job title per person

## What landed

`FIX-M257x-iter30-succession-role-tiebreak` is closed, in two cuts, the second forced by a regression the
suite caught inside the same iteration.

### Cut 1 — bound the population's role set

Supporting members drew a role via `pool.at(hash(prefix:role:i))` over the full 300-role public pool, one
independent hash each. Measured on `demo-1`: **40 members, 39 distinct titles**, a single title with two
holders. Members now draw from a **bounded per-run role set (12 titles)** that always contains the
heroes' declared roles.

Twelve is chosen to hold across the seeded size range rather than tuned to one org — ~3 incumbents per
title at size 40, ~17 at the showcase's 200.

### The sibling sweep, which was the larger half

The derivation was **hand-copied**. The identical
`roles.at(int(hashInt(fmt.Sprintf("%s:role:%d", prefix, i))))` expression stood in **six** production
seeders — `users`, `certificates`, `membership_skills`, `population_evidence`, `profile`, `target_roles`
— and a **seventh** copy inside `target_roles_test.go`.

**The first sweep of this fix found four.** The other two spell the index `idx` and read a different
field off the result, so the obvious grep missed them. Left at four, a member's stored
`memberships.job_role_name` would have silently diverged from the role their profile title and
role-mobility rows were generated for — a cross-seeder divergence no test asserts. This is iter-25's
half-swept `stackseed` fix a third time, and the reason it did not land a third time is that iter-25
wrote the rule down.

The seventh copy — the one in the test — went RED the moment the six production copies were unified.
That is the desync in miniature, demonstrated rather than argued.

### Cut 2 — the regression, and the control that caught it

The first cut gave **every** org **every** story's hero roles from one shared set. The seeded worlds are
multi-tenant and several Playthroughs prove exactly that they are. 52 specs into the binding run:

    ✘ negative-controls.spec.ts:429
      CONTROL for pt-workforce-succession: the proof org's seeded role ("DevOps Engineer")
      must NOT be one of this tenant's key-role cards — expected 0, received 1

**A fix for a tenancy Playthrough that breaks the tenancy CONTROL is not a fix.** Left in, it would have
reported `pt-workforce-succession` green while deleting the isolation the suite exists to prove. Hero
roles are now scoped to their own story and the pool window is offset by the story prefix, so two orgs
draw largely different titles.

The offline test now asserts both directions — no hero role crosses an org boundary, each org still holds
its own hero's title with multiple incumbents, and the two orgs' sets do not overlap wholesale — so the
next attempt at this fails in a second rather than 20 minutes into a live suite.

## Measurement

**Predicted in the overview before the run: `pt-workforce-succession` passes after a reset-to-seed.**
It does.

Live on `demo-1`, cold reset-to-seed, after cut 2 — **7 of 7 green**:

    ✓ pt-workforce-funnel                                    1.2s   (iter-30's fix, holding)
    ✓ pt-workforce-succession                                1.2s   ← the target
    ✓ negative-controls ×5, including :429                   5.7s   ← the one that caught cut 1

The overview's declared-acceptable alternative (concentration dropping the role *below* the critical
threshold, so the correct fix would have been the opposite) **did not occur** — recorded because it was
on record as a real possibility, not a rhetorical hedge.

## Honesty on the clause-2 number

**Clause 2 is NOT claimed to have moved.** The binding run was started, reached 52 of 209 specs, and was
**deliberately killed** when the negative control failed — it was measuring a world carrying a regression,
so its number would have been a number about nothing. Killing it cost ~30 minutes of wall clock and saved
a wrong headline.

So two ids are fixed on **scoped, reproducible** evidence (iter-29 proved the suite deterministic, so a
scoped green on a cold reset is real), and the expectation stands at **`27 / 3 / 1`** — written down
before the confirming run, per the iter-28 discipline. **The full binding run is iter-32's whole job.**

## Close — 2026-08-01

**Outcome:** the seeded org no longer has one job title per person; the member-role derivation is in one
place instead of seven; `pt-workforce-succession` and `pt-workforce-funnel` are both green on a cold
reset, with all five tenancy controls holding.
**Type:** tik
**Status:** closed-fixed (the planned target landed, plus the tenancy correction its first cut required)
**Gate:** NOT MET (3 of 5; clause 2 last *bindingly* measured `25 / 5 / 1`, two ids since fixed on scoped
evidence, confirming run routed to iter-32)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (platform origin `2adcf71`
re-checked at open and close; occurrence stays 1 of 2) — (4) user-blocker: n — (5) cap-reached: n —
(6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (bounded per-org role set + the 12 constant), D2 (the six-plus-one sweep), D3 (kill the
binding run rather than let it finish on a regressed world).
**Side-deliverables:** none.

**Routes carried forward:**

| item | why | target |
|---|---|---|
| `MEASURE-M257x-iter31-clause2-binding-run` (supersedes `MEASURE-M257x-iter30-…`) | Two fixes are in on scoped evidence. Budget a full `--reset` run as an ENTIRE iteration; the artifact is now self-describing and preserved (iter-30), so it will survive later diagnostics. | iter-32 |
| `CHECK-M257x-iter27-drilldown-target-coupling` | `pt-activity-drilldown` still fails on the hero's NAME being absent from the first content row's per-member breakdown. Its own spec comment says the first row is a hero-session content "because the seeded heroes' sessions are dated today" — a coupling to grid ordering that iter-27's new hero sessions plausibly disturbed. Best-evidenced remaining id. | iter-33 |
| `CHECK-M257x-iter30-scoped-classifier-misses-filenames` | Unchanged, and now slightly sharper: a bare spec filename would write a *binding* snapshot from a one-spec run. | later tik |
| `CHECK-M257x-iter28-assignment-flip-is-stateful` | Unchanged. | later tik |
| `DOC-M257x-iter31-role-concentration-believability` | The 39-titles-for-40-people shape was a **believability** defect the coverage gate never caught, because that gate checks per-section cardinality and persona self-consistency, not population-level plausibility. Worth a line in `profile-completeness-spec.md`. | later tik |
| `DOC-M257x-iter30-job-role-title-unfilled` | Unchanged. | later tik |

**Lessons:**

- **Grep for the expression, then grep for its shape.** The sweep found four of seven copies because two
  spelled the index differently and one lived in a test. The rule that saved it was iter-25's, already
  written down: assert **completeness**, not presence.
- **A negative control is worth more than the test it guards.** Cut 1 made the target Playthrough pass and
  broke the control that says the pass means something. Only the control could tell those apart — the
  target spec itself was green on a world where tenancy had been quietly deleted.
- **Kill a measurement the moment you know what it is measuring.** The binding run was 25% done when the
  control failed. Letting it finish would have produced a precise number about a world that was about to
  be thrown away.
- **Say what the opposite result would have meant, before measuring.** The overview recorded that more
  incumbency might *lower* key-person risk and push the role off the list entirely — the opposite fix. It
  didn't happen, but writing it first is what would have stopped a "concentration doesn't work" conclusion
  from being drawn from one ambiguous run.
