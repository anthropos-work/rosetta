---
iteration_type: tik
status: closed-fixed
controlling_strategy: TOK-08
date: 2026-08-09
---

# iter-196 — the last unread language, and the sentence that would have kept it unread

**Type:** tik · **Active strategy:** `TOK-08` · **Protocol:** `corpus/ops/platform-alignment.md`

## Step 0 — Re-survey before targeting

iter-195 closed Go and routed `SURVEY-M257x-iter195-typescript-is-now-the-only-unread-language`
minutes earlier. Re-survey is trivially satisfied; what needed checking was **my own next sentence.**

iter-195's lesson is that *"no runner here collects it"* is a fact about the runner. The tempting next
claim — *"Playwright needs a live stack, therefore unreadable"* — is the identical shape, one language
over, and would have closed the route by assertion.

## Cluster / target identified

The 75 TypeScript specs: `stack-verify/e2e` (30) and `playthroughs/e2e` (45).

## Hypothesis

Enumeration does not need a stack even though execution does — so the population is measurable now, and
the run is not.

## Expected lift

A first TypeScript population count, **with the enumerate-vs-run distinction fenced** so the number can
never be quoted as a green.

## Phase plan

1. `playwright test --list` on each section — no stack, no browser.
2. Report `tests`/`files`, never `pass`.
3. Fence the vocabulary as well as the count.

## Escalation conditions

If listing needs a network install, record that as the measured prerequisite rather than installing.

## Acceptable close-no-lift outcomes

A measured *cannot-enumerate-here*, with the reason, closes the route.
