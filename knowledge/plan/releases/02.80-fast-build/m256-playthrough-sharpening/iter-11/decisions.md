# M256 · iter-11 — decisions

## D44 — the `blocked` outcome must come from a seeded ABSENCE, and the flag is an opt-OUT

`blocked` sat at 0 for 23 Playthroughs because **the seeded world contained no refusal**, not because nobody
wrote the test. Every membership in every `pt-*` org carried the g3 `FEATURE_JOB_SIMULATIONS` grant (20/20 ·
40/40 · 40/40 · 40/40, measured). Producing the outcome therefore had to be a **seed** change, and it is:
`StoryOrg.sim_feature_disabled` → `ResolvedStory.SimFeatureEnabled()` → the `UsersSeeder`'s guard.

It is an **opt-OUT**, not an opt-in, and that direction is a decision: the grant has been unconditional since
M42e iter-09 precisely because a demo whose members cannot launch a sim is a broken demo. Flipping the default
would have required every existing preset to opt in, and any preset that forgot would have degraded silently
into a demo that refuses its own headline feature. The default stays byte-identical; only an org that asks is
withheld.

**Rejected: faking the refusal in the harness** (route interception, a forced modal). A refusal produced by the
test proves nothing about the platform — the same objection iter-05 raised against `check({force: true})`
manufacturing a checkbox state the app never learned.

## D45 — assert a refusal from four independent directions

A `blocked` outcome is the easiest outcome to satisfy by accident: **a page that failed to load also fails to
show a launch confirmation.** So the Playthrough pins the deny dialog PRESENT, the org NAMED, the launch
confirmation ABSENT, and the URL still on the detail route. A dead page satisfies exactly one of those; a
mis-tenanted refusal satisfies three. This is the negative-outcome analogue of the coverage protocol's move
away from `textLen > 40`.

## D46 — the refusal names the org, so assert the org

The live probe found the dialog's second line: *"Please contact your administrator at **Halcyon Retail** to
request access."* That upgraded the planned assertion before any spec was written — the test now proves the
enforcer denied **this membership in this org**, which is the M219 lesson (*"a surface that renders is not the
same as the RIGHT surface"*) in the negative direction. A refusal that renders for the wrong tenant, or for
everyone, would have satisfied a bare "the deny dialog appeared" check.

## D47 — `--reset` leaked `g3` grants for four releases; fix the CLASS, pin it as an exact set

`resetCasbin` deleted `WHERE p_type = 'g2'` only. Measured on `demo-2` right after a reset + re-seed: **731 g3
rows for 140 memberships, 540 orphaned.** Because seeded ids are **deterministic**, a stale g3 row from a
previous seed re-applies to the re-created membership — so a `--reset` did not reset the authz state, it
**merged** it, and the org declared as having no AI Simulations came up granted 20/20.

The fix is written against the seeding fleet's grouping policies **as a class** — `resetCasbinPTypes = {g2, g3}`,
named once, rendered into the DELETE by `quotedList`, pinned by `cmd/stackseed/reset_casbin_test.go` as an
**exact set** in both directions: too few leaks state (the bug), too many widens a destructive DELETE past what
the fleet seeds. It stays a targeted DELETE and can never become a TRUNCATE — the table also holds
`init_policy.sql`'s global `p`-rows, and wiping those blanket-403s the stack (the M17 ISSUE-7 class).

**Not deferred**, even though it was outside the iter's planned scope, because landing it was the only way to
*observe* the planned deliverable. Recorded as a side-deliverable, not folded into the close status.

## D48 — an additive leftover in a reset path is invisible while every test wants the thing PRESENT

The generalisation of D47, and the iter's most transferable finding. The leak survived four releases under a
suite that was green on every run — **because all 23 assertions were success assertions**. The first Playthrough
in the suite's history that required something to be ABSENT found it on its first execution.

So the argument for negative controls is not "more rigour". It is that a negative control **looks in a
direction nothing else looks**. Two consequences for how the rest of clause 2 should be read: (a) the remaining
16 uncovered Playthroughs are not a formality, they are 16 unlooked directions; (b) when auditing any reset or
teardown path, enumerate what it deletes against what the seeders **write**, not against what previous bugs
taught you to check.

## D49 — the negative-control count is computed, not narrated

The mutation fence has counted MUTATING Playthroughs since iter-06 under an explicit doctrine (*"a gate whose
metric is a prose claim is not a gate"*), while the **negative-control** figure — the milestone's largest
remaining gap — stayed a prose number carried between iter docs. It is now computed from the same tags, reports
the uncovered ids by name on every run, and carries a **no-regression floor** (8).

A floor and not an equality, deliberately: the gate's target is *every* Playthrough, the count climbs across
iters, and a fence that must be edited on every increment gets edited without being read. What it cannot do is
go quietly backwards.

Also pinned in the same place: a **cross-vantage pair contributes two**, because the relation is symmetric —
each member asserts the same locator in the opposite direction, so each is genuinely the other's control.

## D50 — the "empty `<main>`" deny-surface note was wrong, and was corrected in place

Two comments described the same surface incompatibly: the page object named a text modal, `identity.go` said the
deny condition renders an *"empty `<main>`"*. The live probe settled it — the `<main>` renders the full sim
detail and the refusal is a `role="dialog"` over it. **Neither comment was evidence**; a spec built on the
"empty main" reading would have passed for any broken page. The overtaken note was corrected where it lived
rather than only in the new code, because a wrong comment left in place re-teaches the wrong thing to the next
reader.

## D51 — re-pin the consuming stack in the same iter that changes the seed schema

Adding `sim_feature_disabled` to `pt-world.seed.yaml` made the file **unreadable to the stack's previously
pinned `stackseed`** — and the failure mode is the worst available: the reset had already TRUNCATED the world
when the re-seed died on `field sim_feature_disabled not found`. That is the *tagging-is-not-publishing* class
(rung zero) reached from a new direction: not a missing tag, but a consumer left behind by a schema change.

So the iter closes with the tag pushed to origin **and** `stack-demo/rosetta-extensions` re-pinned to it, its
own `stackseed` rebuilt, and the suite re-run on that binary (run 4: 148 passed, 24/31, 0 failing). The rule:
**a seed-schema change is not landed until a consumer built from the tag can read it.**

## D52 — the clause-1 drift is the machine, and the cross-check is how we know

Clause 1 read 0.6863× against iter-09's 0.5652×. The **honesty cross-check over the original 16 specs, which
this iter did not touch, moved by the same factor** (0.5284× → 0.6055×). A subset with zero code change cannot
regress from code, so the drift is the environment — four back-to-back suite runs on a 9.70 GiB Docker VM
against the documented 12 GB floor. Reported as variance, with the control shown, exactly as iter-08 did in the
opposite direction. **No lift claimed and none needed:** clause 1 was already MET and this iter's job was to
re-verify it on the grown denominator, not to improve it.
