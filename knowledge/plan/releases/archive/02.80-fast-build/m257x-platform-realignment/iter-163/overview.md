---
iteration_type: tik
status: closed-fixed
---

# iter-163 — does the cited line SAY THE THING?

**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
— *intra-corpus mis-citation is now the largest single class … and that class is mechanical.* This
iter takes the sub-class that had a **live, measured instance** as of yesterday.

## Step 0 — re-survey

iter-162 closed on a side discovery: `corpus/services/backend.md:182` cited a test line **5 lines
away from its subject**, and was **green**, because `:435` happened to be ordinary code. A one-line
edit elsewhere shifted a `)` onto it and `anchor_construct_guard` went RED the same second.

`anchor_construct_guard`'s own docstring already says why: it checks whether a line *exists* and
whether it carries *content*; nothing checks whether the line **says the thing**. Catching the
general case *"requires deciding what a sentence claims, which is the line this whole fence family
does not cross."*

**Target confirmed, and narrowed to the decidable slice.** The general case is semantic. But when the
prose puts a **backticked literal** beside a citation and that literal occurs in the cited file
**nowhere near the anchor**, nothing has to be interpreted — and the census can name the line the
citation should have carried.

## Cluster / target identified

`FIX-M257x-iter138-anchor-rot-fence`, which iter-162 upgraded from a suspicion to a measured
instance. This is its first instrument.

## Hypothesis

A pairing-aware census over (citation, adjacent quoted literal) will enumerate the decidable slice of
anchor rot corpus-wide, at a stated denominator, with each finding carrying its own proposed repair
(*the line the literal is actually on*).

## Expected lift

Instrument + a run to zero on the enumerated class. No `P`/`N` reading is taken (`§9`: the metric
stays **UNMEASURED**, not unmoved).

**Falsifiable:** if the class is empty after honest pairing, the iter reports that. What refutes the
approach is being unable to separate signal from the cross-product — the first draft's failure mode.

## Phase plan

- **A** — build the predicate; measure it; sharpen it until it measures the class, not the pairing.
- **B** — grade **every** survivor at source (iter-158: a proposed repair is a hypothesis).
- **C** — repair the true positives **against the subject**, never by bumping the offset (`§5` rule 22).
- **D** — fence: every clause in both directions, anti-vacuity, stale-exemption, live ratchet.

## Escalation conditions

- A finding needs a platform-repo edit → route, never edit.
- A clause would have to be tuned until a known instance fires → **Trap A**; drop the clause instead.

## Acceptable close-no-lift outcomes

The class enumerates to zero true positives — the reach is the deliverable, and the denominator is
stated either way.
