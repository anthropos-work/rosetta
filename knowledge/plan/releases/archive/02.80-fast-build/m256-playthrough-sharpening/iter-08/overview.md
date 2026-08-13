---
milestone: M256
iter: 08
iteration_type: tik
status: closed
active_strategy: TOK-01
created: 2026-07-28
---

# M256 · iter-08 — the onboarding product exists

**Type:** `tik` · **Active strategy:** `TOK-01` move 3 (land the coverage clusters), unblocked by iter-07 D28.
This is `ONBOARD-M256-build`.

## Step 0 — re-survey

`ptvalidate --manifest-dir manifest --e2e-dir e2e/tests --seed-worlds …` VALID: 9 products, 23 use cases,
21 live, 2 TODO. Clause 1 at 0.6245×; clause 2 mutating 5/5 MET, negative controls 5 of 21, `blocked` 0;
onboarding 0 of 5. Target current and unchanged.

## Cluster / target identified

Onboarding: **0 of 5 covered, un-homed for five releases** — *the first thing every real user does was the
largest untested surface in the product* — and **5 of the 9 use cases clause 3 must land**. iter-07 D28 retired
the blocker (`user_params.onboarding` is NULL for every seeded user; `/onboarding` drives), so the cluster is
buildable now and nothing else in the milestone unlocks as much coverage per unit of work.

## Hypothesis

**H1.** `/onboarding` is **its own read-back**: it SERVES the flow while onboarding is incomplete and REDIRECTS
to `/home` once complete. If so, one route yields both halves of a mutating proof — the pre-state absence and
the persisted post-state — each on a fresh navigation, with no toast and no DB backdoor.

**H2.** The `pt-free` seat can absorb the write safely: it is registered in `seed-worlds.yaml` and driven by
**0** other use cases, so a persistent hero mutation cannot perturb another Playthrough's expectations.

## Expected lift

- The onboarding product opens: **1 live Playthrough** + **all 5 curated onboarding UCs declared with written
  verdicts** (clause 3's zero-silent-gaps requirement for this cluster).
- Clause 2's mutating count **5 → 6**, with its negative control free from H1's route flip.

## Phase plan

Steps 3 → 4 → 5 → 6 of `corpus/ops/demo/playthroughs.md` § The iteration protocol. Probe the whole flow live
first (the milestone's standing lesson), then page object → spec → manifest → gate run ×3 → close.

## Escalation conditions

- If the completion cannot be read back through the app (only via the DB), the proof shape weakens to a DB
  assert — record that honestly rather than claiming a UI read-back.
- A platform defect escalates or gets a diagnosed draft. Never a platform edit.

## Acceptable close-no-lift outcomes

- The flow proves undrivable past its first step for a reason iter-07's probe did not see; that would re-open
  the clause-3 onboarding question with evidence.
