---
milestone: M257x
iter: 31
iteration_type: tik
status: closed-fixed
---

**Type:** tik (under `TOK-01: instrument first, then follow`)

# iter-31 — the org that had one job title per person

## Active strategy reference

`TOK-01: instrument first, then follow`. Specifically its second clause — **fix the mechanism, not the
symptom, and derive the list instead of maintaining it by hand.** That clause turns out to describe this
iter twice over: once for the seeded world, once for the derivation that assigns it.

## Step 0 — re-survey

iter-30 measured the cause on the live surface and routed it. Re-confirmed at open rather than assumed:
platform origin still `2adcf71`; the target (`FIX-M257x-iter30-succession-role-tiebreak`) is untouched
and still the best-evidenced remaining failure.

## Cluster / target identified

`pt-workforce-succession`, the one id iter-30 root-caused but deliberately did not fix (scope-creep
tripwire — it was a third line in a two-line iter). The cause as measured then:

- the succession view ranks **"Roles by risk"**; `Critical roles: 28`, and the surface renders **25**;
- every rendered card reads the identical `MEDIUM / risk 68` — all 28 are **tied**;
- the seeded org is **40 members across 39 distinct job roles**, one incumbent each bar one;
- so whether the hero's role gets a card is decided by a **tiebreak among 28 equal values**.

## Hypothesis

The assertion is legitimate — *did the manager's own tenant's projection reach her own org's roles?* — and
it was the **seeded world** that could not support it. Concentrate the supporting population into a
realistic number of job titles and the critical-role set drops below the view's render budget, so the
ranking has nothing left to truncate. The hero's role must be **in** that set, or she is again its sole
holder and nothing has changed.

Predicted side-benefit, and the reason this is not a test-shaped hack: a 40-person company with 39 job
titles is not believable, and the demo's believability is the actual product requirement the atomised
distribution was quietly failing.

## Prediction, recorded BEFORE the measurement

- `pt-workforce-succession` passes after a reset-to-seed.
- **Declared acceptable in advance:** the opposite result is informative and not a failure of the iter —
  if concentrating incumbency drops the role *below* the `risk ≥ 50` critical threshold (more holders =
  less key-person risk), the role disappears from the list for a new reason, and that measurement tells us
  the surface wants a *scarce* role, which is the opposite fix. **Measure; do not reason about the
  platform's risk formula from the outside.**
- No prediction is offered for the other four ids. Two are known-different causes; guessing would be the
  cluster mistake this milestone keeps making.

## Phase plan

1. Bound the supporting population's role set, ensuring hero roles are always in it.
2. **Sweep every call site of the member-role derivation** — it is hand-copied, so a partial change
   desynchronises a member's stored role from the role their other rows were generated for.
3. Offline unit test + reset-to-seed + a targeted scoped run.
4. **A full binding run** — which also discharges `MEASURE-M257x-iter30-clause2-binding-run`, since one
   reset buys both the succession fix and iter-30's funnel fix a real clause-2 number.

## Expected lift

Clause 2 was measured `25 / 5 / 1` (iter-29, deterministic over three runs). Two ids have since been
fixed on scoped evidence (funnel at iter-30, succession here), so the honest expectation is **`27 / 3 / 1`**.
Recorded before the run, per the iter-28 discipline: **a number that beats this deserves more suspicion
than one that meets it.**

## Escalation conditions

- A platform-repo edit required → route forward. Binding.
- Platform origin moves off `2adcf71` → re-scope trigger occurrence 2 → STOP.
- The binding run comes back **worse** than 25 → the role-concentration change has side effects across the
  209 specs; that is a real risk of touching the shared seed, and the answer is to measure which ids moved,
  not to defend the change.

## Acceptable close-no-lift outcomes

The concentration fix failing for a measured reason (e.g. the role drops below the critical threshold),
with the tie-break mechanism characterised — that is a complete iter.
