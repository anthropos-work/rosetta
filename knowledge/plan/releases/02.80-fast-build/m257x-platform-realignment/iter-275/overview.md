---
iter: 275
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
route: FIX-M257x-274-devops-occupancy-must-stay-at-two
---

# iter-275 — the occupancy invariant is unpinned in BOTH directions

**Type:** tik, under `TOK-08` (*census the mechanical classes; stop sampling them*).

## Step 0 — re-survey (mandatory, before targeting)

iter-274 named the fix as *"hold `DevOps Engineer` occupancy at ≤ 2"* and called it **one number**.
Re-surveying the mechanism before implementing shows that framing is **too small**, and implementing
against it would have produced a re-roll rather than a fix.

`seeders/jobroleref.go` is the single derivation of a supporting member's role:

- `orgRoleSet(prefix, storyHeroRoles)` — the story's own hero roles **first**, then a per-org window of
  the pool, capped at `orgRolePoolSize = 12`. Its comment states the invariant it exists for:
  *"so no hero is ever the sole holder of her own title"*.
- `memberRoleAt(...)` — `set[hashInt("<prefix>:role:<i>") % 12]`. A hash draw. **Uniform.**

**`orgRoleSet` only guarantees the hero's role is IN the set. Nothing makes any supporting member
actually draw it, and nothing caps how many do.** Occupancy of a hero's role is therefore a lottery with
mean ≈ 28/12 ≈ 2.3, against an at-risk threshold that flips between 2 and 3 (iter-274).

**So the M256 iter-14 green was a draw, not a property** — and the same lottery is what this iter must
measure in both tails, because the *other* tail is a violated invariant, not just a broken test.

## Cluster / target identified

The **occupancy distribution over Org A's 12 roles**, read from the payloads already on disk, and the two
hero roles' positions in it:

- `DevOps Engineer` (Pat Ellis) — **3** incumbents → `riskScore 45` → below the guard → **the Playthrough
  fails**.
- `Engineering Manager` (Morgan Reyes) — **1** incumbent → **Morgan is the sole holder of her own title**,
  which is precisely the state `orgRoleSet`'s comment says the mechanism prevents.

One mechanism, two failures, opposite tails.

## Hypothesis

The invariant is **unenforced**, not merely mis-tuned: no code path bounds hero-role occupancy from above
or below. If so, the fix is structural (bound the draw) and a re-roll of the seed would only move which
hero is broken.

## Expected lift

- The full occupancy distribution, with both hero roles located and the total reconciled against the
  member count (a census that must sum back, or it is not a census).
- The invariant's status established by **reading the code path**, not inferred from one org.
- The implementation **specified against the real mechanism**, so iter-276 edits with a bounded blast
  radius rather than discovering it mid-edit.

## Phase plan (declared multi-step — the tripwire counts UNPLANNED lines only)

1. Seal these pre-registrations (first commit).
2. Census the occupancy distribution from the captured payloads; reconcile the sum.
3. Read `orgRoleSet` / `memberRoleAt` for any bound on hero-role occupancy; state the verdict.
4. Specify the implementation, including its blast radius, and record the decision.

## Out of this iter's planned scope (declared, so the tripwire is clean)

**The edit itself.** `memberRoleAt` is the **single** derivation for **six** production seeders, and its
own comment records that the previous unification *"found only FOUR"* of them and that a seventh copy in a
test went RED when the production six were unified. A bounded-occupancy fix needs the population size,
which the current signature does not carry — so it is an **11-site** change (6 production + 5 test) to a
shared derivation. Starting that with a partial budget is the exact *half-swept fix* the file warns about
**twice**. It is iter-276's work, specified here.

## Escalation conditions

- **No platform edit** (v2.8), and no edit to the pinned `stack-demo` clone.
- **Do not start the 11-site edit in this iter.** If the specification turns out to be cheap enough to
  land safely, that is still a *next-iter* decision — the tripwire exists for exactly this temptation.
- No new stack runs are required; the payloads are on disk.

## Acceptable close-no-lift outcomes

The invariant proving to be **enforced** somewhere I have not read — making Org A's state impossible and
my census wrong — is a complete result, and the one that would most change iter-276's plan.

## Pre-registrations (sealed in this iter's FIRST commit, before any measurement)

- **PR-1 — the census sums back.** Per-role incumbent counts over the 12 roles total exactly the member
  count the projection reports (**28**, iter-267's `querySuccessionMembers`). *Refuted by:* any other sum
  — which would mean incumbency and membership are different populations and iter-274's step function was
  read against the wrong denominator.
- **PR-2 — Morgan is the sole holder of her own role.** `Engineering Manager` has exactly **1** incumbent,
  and it is the manager hero. *Refuted by:* any other count, or a non-hero occupant.
- **PR-3 — nothing bounds hero-role occupancy.** Neither `orgRoleSet` nor `memberRoleAt` contains any
  cap, floor, reservation or re-draw keyed on a role being a hero's. *Refuted by:* any such bound.
- **PR-4 — the two hero roles sit in different tails.** Their occupancies differ, and at least one of them
  is not the modal occupancy. *Refuted by:* both equal to each other.
