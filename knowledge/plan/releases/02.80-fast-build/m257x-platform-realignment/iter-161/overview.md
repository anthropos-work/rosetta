---
iter: 161
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-08
---

# iter-161 — run the frozen-expectation census to zero

**Active strategy reference:** [`TOK-08`](../decisions.md) — *build or extend a fence that enumerates
every instance in the corpus, **run it to zero, and keep it green***. iters 159–160 built the instrument
and enumerated the population; this iter does the third clause, which is the one that turns a census into
a fence.

## Step 0 — re-survey before targeting

`FIX-M257x-iter160-iter155-fixture-is-frozen-at-8-sites` was routed **one iter ago**, so it is current by
construction. The re-survey is the census itself, re-run at HEAD: **10 unexempted candidates over 9,370
multi-token literals**.

**Why this population and not iter-159's 961:** TOK-08 says work classes in **descending measured size**,
and on candidate count 961 > 10. But the deliverable is *a class run to zero and kept green*, and these
are two orders of magnitude apart in cost while both close **one** class. The 961 is also, by iter-159's
own reading, mostly rule 71's **sanctioned** residual — so it is a grading job, where these 10 are
individually attributed to a named derivation. Taking the tractable one first is not scope avoidance; it
is the only one of the two that can reach zero inside an iter.

## Cluster / target identified

All 10 enumerated candidates, each graded, none left ungraded:

| site | shape | first read |
|---|---|---|
| `test_bringup_verify_scope_m257x.py` ×6 | the platform-side argument to `scope-union.sh` / the tail | frozen copy → **derive** |
| `test_bringup_verify_scope_m257x.py` ×2 | `st.write_override([...])` — the synthetic override | frozen copy → **derive** |
| `test_platform_predicate_guard.py:193` | `assertEqual(c.select("core"), {…})` against the **live** platform | needs grading — a golden over a *live* input is the class |
| `stack-injection/tests/test_platform_topology.py:120` | expected output of a **synthetic** compose the test wrote | expected **exemption**, not a repair |

**These first reads are hypotheses, per iter-158.** Each is confirmed at source before it is acted on.

## Hypothesis

Eight sites derive cleanly. The tenth is a legitimate golden (a literal expected output of a *controlled*
input is not a frozen copy of a *live* derivation — the distinction the exemption mechanism exists for).
The ninth is the interesting one and is graded on its merits rather than assumed either way.

## Expected lift

The frozen-expectation census reaches **0 unexempted candidates** and becomes a fence that stays green —
completing TOK-08's three-clause shape for this class.

## Phase plan

- **A** — grade all 10 at source.
- **B** — repair by deriving; declare exemptions with reasons at the site.
- **C** — census to zero; re-run the affected suites (`stack-core` scope + predicate-guard tests,
  `stack-injection`'s topology suite — the section iter-160 named as carrying a candidate whose suite it
  did not run).
- **D** — a regression fence that the census stays at zero; close.

## Escalation conditions

- A repair changes what a test actually asserts → **stop and revert that site**; a census-driven edit
  that weakens a test is worse than the frozen literal it replaced.
- The ninth site's repair needs a decision about the guard's contract → grade it, route it, do **not**
  half-land it (iter-155's rule).

## Acceptable close-no-lift outcomes

Any site whose repair would weaken an assertion, reverted with the reasoning recorded.

## Explicitly OUT of scope (tripwire pre-declared)

iter-159's **961** haystack candidates. Different instrument, different cost class, still routed.
