---
milestone: M258
iter: 11
iteration_type: tik
status: closed-fixed
opened: 2026-08-12
---

# M258 iter-11 — the space axis: where a bring-up's disk actually goes

**Type:** tik · **Active strategy:** `TOK-01`, extended by the **user's ruling of 2026-08-12** (`D52`,
`D57`): M258 is achieved on clauses 1/2/4/5, and the remaining budget goes to build-time fruit **plus a
net-new SPACE axis** — pre-build, post-build, post-teardown.

## Cluster / target identified

**There has never been a space budget the way there is a time budget.** `build-budget.md` prices the
cycle in seconds and names a free-disk *floor*, but nothing anywhere says what a bring-up **consumes**,
what a teardown **returns**, or what a `--purge` leaves behind. The user asked the question directly
and it had no documented answer.

## Hypothesis

Post-teardown is where the defect is, because it is the only phase nobody measures — and a leak there
compounds silently across every cycle this release has been running.

## Phase plan

- **A** — measure the real state (`docker system df` + ownership), correcting for the SIZE-vs-reclaimable
  trap before quoting anything.
- **B** — reclaim only what is provably safe, with ownership verified per target.
- **C** — find the producer, so the leak can be prevented rather than swept.

## Escalation conditions

- Any reclaim that cannot be proven not to touch `demo-2` or the dev stack → do not run it.
- Build-cache pruning to win space → **forbidden as a default** (`D58`): it converts a space win into a
  time loss on the next build, which is the thing this release exists to protect.

## Acceptable close-no-lift outcomes

- The measurement lands and shows the space is already tight. Then the budget is the deliverable.
