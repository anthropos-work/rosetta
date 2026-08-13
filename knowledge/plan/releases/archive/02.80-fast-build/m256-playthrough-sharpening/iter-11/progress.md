# M256 · iter-11 — progress

**Type:** tik · **Active strategy:** `TOK-01` move 4 · **Handler:** `BLOCKED-M256-refusal-surface`.

## Phase A — the probe, and why it had to come first

The routing named the surface (`SimulationPage.orgMemberCannotStartModal()`) but the in-repo evidence about
what that surface *renders* was **contradictory in two places**: the page object names a text modal, while
`stack-seeding/seeders/identity.go:250` said the same condition renders *"the org-member deny modal (empty
`<main>`)"*. Those are different assertions and only one of them is safe — a spec built on "the page is blank"
passes for a broken page too. So: withhold the g3 grant for one org on `demo-2` (20 rows), Sentinel-Reload,
drive the real browser as `pt-free`, and read the DOM.

**H1 CONFIRMED, and the refusal is richer than the routing knew:**

| what was read | value |
|---|---|
| sim detail reachable | **yes** — `Start Simulation` still present (the gate is on LAUNCHING, not browsing) |
| deny dialog | **1** — a real `role="dialog"` |
| its text | *"You cannot start AI Simulations in this organization / **Please contact your administrator at Halcyon Retail to request access.**"* |
| launch confirmation | **0** |
| `atLaunchBoundary` | **false** — the URL never advanced to `/sim/<slug>/start` |
| `<main>` | **1, populated** (the sim detail renders behind the dialog) |

Two consequences beyond a yes/no on H1:

1. **The `identity.go` "empty `<main>`" note is wrong**, and it is the kind of wrong that produces a false
   green. Corrected in place.
2. **The dialog NAMES THE ORG.** So the assertion can prove *which tenant* was refused, not merely that
   something was refused — the M219 lesson applied in the negative direction. That upgraded the planned
   assertion before a line of spec was written.

Grants restored, probe spec deleted.

## Phase B — what landed

**The seed side (`stack-seeding`).** `StoryOrg.SimFeatureDisabled` (yaml `sim_feature_disabled`) → threaded
into `ResolvedStory` → read through **one** recognition point, `SimFeatureEnabled()`, the same shape as the
existing `IsHiringOrg()`. The `UsersSeeder` guards the per-membership g3 grant on it. It is an **opt-OUT**
deliberately: the grant has been unconditional since M42e iter-09 because a demo whose members cannot launch a
sim is a broken demo, so the default is byte-identical and only an org that asks is withheld.

**The world side.** `pt-world` Org B (Halcyon Retail) declares it — the org **no** sim Playthrough drives, so
Org A keeps its grant and `pt-aisim-chat-launch` is untouched. `seed-worlds.yaml` gains the
`sim-feature-disabled` capability as the deliberate complement of `sim-feature-enabled`.

**The journey.** `ai-simulations.access-denied.UC1` → `pt-aisim-org-feature-blocked`, `outcome: blocked`.
It pins the refusal from **four** directions, because a `blocked` outcome is the easiest one to satisfy by
accident: the dialog is PRESENT, it NAMES the member's own org, the launch confirmation is ABSENT, and the URL
is still the detail route. *A dead page satisfies exactly one of those.*

**The pairing, which is the cheap part.** `pt-aisim-chat-launch` already asserted that same locator ABSENT for
a granted hero. Naming the two as each other's negative control cost two comment blocks and delivers the
`NEGCTL-M256-cross-vantage` mechanism live: one locator, two orgs, opposite verdicts, both on every run.

**The count.** The mutation fence now computes the **negative-control** figure too (`8 of 24`, uncovered ids
named, no-regression floor) — clause 2's largest gap had been a prose number quoted from iter to iter, which is
exactly what iter-06's own fence header forbids.

## Phase B′ — the RED that was worth more than the green

**`pt-aisim-org-feature-blocked` failed on its first live run**, and the reason was not the spec.

`stackseed --reset` deleted only `g2` casbin rows. **`g3` had been accumulating since M42e iter-09** — measured
immediately after a reset + re-seed on `demo-2`: **731 g3 rows for 140 memberships, 540 of them ORPHANED**
(pointing at membership ids from worlds already truncated). Litter is the mild half. The correctness half is
that **seeded ids are deterministic**: a truncated membership is re-created with the *same* uuid, so a stale g3
row **silently re-granted** the feature to the new world. The org the blueprint declared as *not* having AI
Simulations came up granted **20/20**, and the Playthrough went red against a world that was never in its
declared state.

Fixed as a **class**, not a patch: `resetCasbinPTypes = {g2, g3}`, named once, rendered into the DELETE by
`quotedList`, and pinned by `cmd/stackseed/reset_casbin_test.go` as an **exact set** — too few leaks state, too
many widens a destructive DELETE past what the fleet seeds (the table also holds `init_policy.sql`'s global
policy, so it can never become a TRUNCATE).

**Why it survived four releases:** every test wanted the grant to be **PRESENT**. An additive leftover in a
reset path is invisible for exactly as long as nothing asserts an ABSENCE. The first negative assertion in the
suite's history found it on its first run.

*Possibly also a flake cause, recorded as a hypothesis:* run 1 (with the 540 orphans still in the enforcer)
also failed `pt-assignment-assign` on a 240 s timeout; it did not reproduce in any of the four runs after the
fix. A policy set 5× larger than the world is a plausible authz-latency contributor. **Unmeasured** — not
claimed.

## Phase C — re-measure

Four consecutive `--reset` runs, `demo-2`, `localhost/http`, offset 20000.

| | |
|---|---:|
| **Suite** | **148 passed ×4, 0 failed, 0 flake** |
| `ptreport` | **24/31 passing, 0 failing, 7 unimplemented, 0 unimplementable** |
| Clause 1 — median per non-studio Playthrough (n=3, 22 PTs) | **2.282 s = 0.6863×** of the 3.326 s baseline — gate `≤ 0.79×` **MET** |
| Honesty cross-check, the ORIGINAL 16 only | **2.014 s = 0.6055×** |
| Studio lane (excluded) | 1.35 s / 1.86 s |
| Suite wall-clock (reported, not gated) | 56.1 / 71.9 / 74.2 s |
| `@pt-mutation` registry (computed) | **MUTATES=6 READ-ONLY=16 UNKNOWN=2** |
| `@pt-negative-control` registry (computed, net-new) | **8 of 24** |
| `blocked` outcomes | **0 → 1** |

**On the clause-1 drift, stated plainly:** the figure moved from iter-09's 0.5652× to 0.6863×, and the
**cross-check on the untouched original 16 moved with it** (0.5284× → 0.6055×) with **zero code change to
those specs**. So it is the machine (a 9.70 GiB Docker VM against the documented 12 GB floor, four suite runs
back-to-back), not a regression attributable to this iter — the same variance iter-08 retired in the other
direction. Clause 1 holds with margin either way. The new Playthrough sits at 2.39 s, just above the median.

**Reproduced on the stack's OWN pinned tooling.** The first three runs used a `stackseed` built from the
authoring copy (the seeder change is unreleased). After tagging and pushing, `stack-demo/rosetta-extensions`
was re-pinned to `fast-build-m256-blocked-outcome` and its own binary rebuilt — run 4 is that binary:
`deleted 362 casbin grant(s) (g2 role + g3 feature)`, **148 passed, 24/31, 0 failing**. This matters as more
than hygiene: the moment the seed gained a field, the stack's previously-pinned `stackseed` **hard-failed the
reset** (`field sim_feature_disabled not found in type blueprint.StoryOrg`) — a live instance of the
*tagging-is-not-publishing* class, caught and closed inside the iter rather than left as a landmine for the
next run.

## Close — 2026-07-28

**Outcome:** **clause 2's `blocked` sub-clause is DISCHARGED — 0 → 1**, by a refusal that comes out of
Sentinel's own Casbin enforcer rather than the harness, and the `NEGCTL-M256-cross-vantage` mechanism is
**live** on its first pair (negative controls **6 → 8 of 24**, now machine-counted). The `blocked` Playthrough
earned its keep on its first run by exposing a four-release-old defect in `--reset`: it deleted `g2` grants and
not `g3`, so 540 orphaned grants accumulated and — ids being deterministic — **silently re-granted** the very
feature the new Playthrough needed absent.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — clause 1 **met** (0.6863×, re-verified on 24 PTs; cross-check 0.6055× shows the drift is
the machine); clause 2 mutating **6/5 MET**, **`blocked` 1/1 MET**, negative controls **8 of 24** (the last
open item in the clause); clause 3 verdict half **COMPLETE**, landed half short (org-admin 2/4, onboarding
1/5); **D-v28-5 still unfixed**, still blocked behind `FIX-M256-cockpit-manifest-drift`.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this tik progressed; iters 08/09 progressed) — (3) re-scope: n (0 of 31 curated UCs `unimplementable`; the new one landed first try) — (4) user-blocker: n (0 red, both trees clean, the batch is green) — (5) cap-reached: n (1st tik of the invocation) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D44 (the `blocked` outcome must come from a seeded ABSENCE, and the seed flag is an opt-out because the granted default is load-bearing), D45 (assert a refusal from four independent directions — a dead page satisfies one), D46 (the deny dialog names the org, so the assertion proves WHICH tenant was refused), D47 (`--reset` leaked g3 grants for four releases; fix the CLASS and pin it as an exact set), D48 (an additive leftover in a reset path is invisible while every test wants the thing PRESENT — the argument for negative controls, measured), D49 (the negative-control count is computed, not narrated), D50 (`identity.go`'s "empty `<main>`" deny-surface note was wrong and would have produced a false green), D51 (re-pin the consuming stack the moment the seed gains a field — the authoring copy and the pinned binary diverged into a hard reset failure).
**Side-deliverables:**
- The `--reset` g3 leak fix + its exact-set fence (**not** in the iter's planned scope — it surfaced as the
  cause of the planned scope's first RED, and landing it was the only way to observe the planned deliverable).
- `stack-demo/rosetta-extensions` re-pinned to the new tag + its `stackseed` rebuilt, so the stack's own
  tooling reproduces the run (run 4).
**Routes carried forward:**
- `NEGCTL-M256-cross-vantage` → **still the largest gap, but no longer unproven: 8 of 24.** The mechanism now
  has a working reference implementation (`aisim-chat-launch` ∥ `aisim-org-blocked`). The 16 uncovered ids are
  named by the fence's own output on every run. Note what the reference case teaches about cost: this pair was
  cheap because the two vantages differ by **seeded state**; a pair that differs only by test code is the
  O(tests) case the routing already priced.
- `FIX-M256-cockpit-manifest-drift` + `D-v28-5-cockpit-logout` → untouched by design (escalated; a reset-lifecycle
  change needing live bring-up verification).
- `FLAKE-M256-assign-under-bloated-policy` → **NEW, a hypothesis only.** `pt-assignment-assign` timed out at
  240 s on the one run made against a 731-row/540-orphan g3 policy set and in none of the four runs after the
  fix. If it ever recurs, measure the enforcer's policy size first.
- `DOC-M256-ptworld-reset-comment`, `FIX-M256-autoverify-fapi-libressl`, `PT-M256-orgadmin-role-create`,
  `PT-M256-orgadmin-member-tag`, `FIX-M256-studio-false-green`, `DOC-M256-llm-lane-premise`,
  `ONBOARD-M256-import-path`, `PT-M256-resume-fixture-pair`, `FENCE-M256-bounded-interaction`,
  `PT-M257-self-evaluation`, `PT-M257-talk-to-data`, `PERF-M256-parallel-lane`,
  `FIX-M257-content-stories-pair-count` — all stand.
- `DOC-M256-claudemd-pt-count` → **NEW (housekeeping).** `CLAUDE.md` still says "18 live Playthroughs"; it
  points at `playthroughs.md` as authoritative, which now reads 24. Reconcile once at milestone close rather
  than on every increment.
**Lessons:**
1. **A negative assertion is a different instrument, not a stricter one.** Twenty-three success assertions ran
   green for five releases over a reset path that was silently accumulating authz grants. The first assertion
   that required something to be **absent** found it on its first run. The value of a negative control is not
   that it is more rigorous — it is that it *looks in a direction nothing else looks*.
2. **Read the seed before writing the spec.** `blocked` was 0 not because nobody wrote the test but because the
   world contained no refusal. Coverage clauses that quantify *outcomes* are claims about seeded state as much
   as about test code; ask which is missing before assuming it is the code.
3. **When two comments in the repo describe the same surface differently, neither is evidence.** One said "deny
   modal", one said "empty `<main>`". A twenty-minute probe settled it and upgraded the assertion (the dialog
   names the org) at the same time. The overtaken comment was corrected in place, because leaving it would
   re-teach the wrong thing.
4. **The moment a seed file gains a field, every consumer pinned to an older tag is broken — loudly, and at the
   worst time.** The pinned `stackseed` hard-failed the reset *after* truncating the world. Re-pin (or build)
   as part of the same iter; do not leave the divergence for the next run to discover.
