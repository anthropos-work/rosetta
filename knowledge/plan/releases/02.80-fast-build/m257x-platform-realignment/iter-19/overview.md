---
milestone: M257x
iter: 19
iteration_type: tik
status: closed-no-lift
opened: 2026-08-01
closed: 2026-08-01
---

# iter-19 — re-measure clause 2 through an instrument that now works

**Type:** tik
**Active strategy reference:** `TOK-01: instrument first, then follow` — step 5, *"prove it cold"*, applied
to the clause the instrument was blind for.

## Step 0 — re-survey (mandatory)

- **Platform origin HEAD:** `2adcf71`, unchanged (re-checked at iter-18 close, minutes ago). Re-scope
  trigger stays at **occurrence 1 of 2**.
- **The routed targets say to re-measure before touching them.** iter-18's close routed
  `FIX-M257x-iter15-directus-versions-403` and `FIX-M257x-iter15-library-category-expansion` forward with an
  explicit condition attached: *both were measured on a stack whose Directus served nothing.*

## Cluster / target identified

**Clause 2's own measurement, not any of its named causes.** iter-15 measured `20 live / 10 failing / 1
unimplemented` on a stack whose per-stack Directus held the content and served it to nobody — the exact
condition iter-17 later measured directly as `anon GET … -> 403` and iter-18 fixed. Every clause-2 failure
attributed to a content-model cause was therefore diagnosed **through a broken instrument**.

`platform-alignment.md` §5's closing rule and rule 12 both point the same way: *a cause measured through a
broken instrument is a hypothesis*, and *say which invocation produced the number*. The cheapest correct
next action is to take the number again, on the stack iter-18 just proved green, before spending an iter on
a cause that may no longer exist.

## Hypothesis

Some non-zero part of the 10 failing Playthroughs was downstream of the unserved content layer, and the
clause-2 denominator's real residual is smaller than 10. **Direction only** — the honest outcomes include
"identical, so the causes are independent", which is itself worth having in writing.

## Expected lift

A re-measured clause-2 number with its invocation recorded. If failures drop, the routed cause list is
corrected; if they do not, the causes are confirmed independent of the Directus defect and the next tik
starts from evidence instead of from an inherited attribution.

## Phase plan (single-step — a measurement iter)

1. `./run-playthroughs.sh 1 --reset` — the **full** suite (the ptreport gate is binding only on a full run)
   with the real reset-to-seed, because iter-15 established that an additive re-run poisons the negative
   controls and makes two runs incomparable.
2. Compare per-id against iter-15's recorded result; re-attribute each failure that moved.
3. Land no source change unless something is both in-scope and complete-able. Anything else is routed.

## Escalation conditions

- The suite cannot run at all ⇒ that is a clause-2 blocker in its own right; record and route, do not patch
  the runner mid-iter.

## Acceptable close-no-lift outcomes

An unchanged 20/10/1 is a **complete** outcome: it falsifies the "downstream of Directus" attribution that
iter-18's close attached to two of the routed causes, and that falsification is the deliverable.
