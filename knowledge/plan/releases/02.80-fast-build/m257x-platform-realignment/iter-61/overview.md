---
iter: 61
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-04
active_strategy: TOK-05
refs:
  platform: 0dab54d
  app: v1.366.0
  rext_pin_at_open: fast-build-m257x-iter-60.1
---

# iter-61 — widen G1/G4 to the two forms iter-60 could not reach, then repair what they name

**Active strategy:** [`TOK-05`](../decisions.md) — repair by predicate, and close the predicate by
making it underivable-when-false.

## Step 0 — re-survey before targeting (mandatory), and a SUBSTITUTION

`TOK-05`'s ordering named iter-61 = *land §7's citation-safety half and spend it on the 21 outstanding
`main.go:N` sites.* Re-measured at open, both candidates are real:

| candidate | measured now |
|---|---|
| the citation class | **real** — of 16 distinct `app/main.go:N` citations, 6 still land on the construct they claim (`:446`, `:524`, `:604`, `:816`, `:992`) and ~10 have moved (`:971-973` is now a comment about collapses, `:1178` is `defer cancelServerContext()`, `:1196` is the *skiller* handler) |
| **the G1/G4 prose gap** | **real, larger, and newly discovered** — iter-60's fence closed the *command* and *table-first-cell* forms; a whole-tree grep finds **~12 further sites** in the noun-phrase form (*"the default `graphql` profile"*) and **2** RPC values in *table-row* form, all invisible to iter-60's constructs |

**Substituting** (permitted by Step 0 — the TOK strategy holds, only the named next-target is stale).
Rationale: the gap was created *by this milestone's own fence one iteration ago*, and leaving it open
means iter-60's GREEN over-reports. A fence whose reach is narrower than its class is the failure this
milestone has found five times; closing it now is worth more than the citation class, which is
already routed, named, and does not silently misreport anything.

`FIX-M257x-iter58-mainline-shift` is **re-routed to iter-62** with its measurement refreshed above, so
the next handler starts from a re-derived list rather than iter-58's.

## Hypothesis

`graphql`-as-a-profile survives in a **third parsed construct** — a backticked token immediately
followed by the word *profile* — and `*_RPC_ADDR` values survive in a **fourth**: a markdown table row
whose first cell is the variable and whose second is the value. Both are parseable; neither is a
substring match on prose.

## Expected lift

G1 and G4 widened to those two forms, watched RED, repaired to GREEN, with the mutation battery
extended. iter-60's GREEN becomes a claim about the whole class rather than about two forms of it.

## Escalation conditions

If the noun-phrase rule cannot be made false-positive-free by deriving it from the corpus's own
structure, **route it forward rather than ship it** — a fence with false positives is disabled on
first contact (§4 Trap A).

## Acceptable close-no-lift outcomes

If the widened rules find nothing beyond what iter-60 already repaired, that falsifies the gap
hypothesis and iter-60's GREEN stands as a whole-class claim — a complete result.
