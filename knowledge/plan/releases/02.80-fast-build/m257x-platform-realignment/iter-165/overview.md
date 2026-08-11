---
iteration_type: tik
status: closed-no-lift
---

# iter-165 — audit the ACCEPT side of the fence family

**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07),
against the axis iter-164 opened: `SURVEY-M257x-iter164-acceptance-clauses-are-unaudited-for-over-reach`.

## Step 0 — re-survey

iter-164 measured an acceptance clause that was hiding a real candidate (a prose-block rule spanning
lines 1–154 of a shell script). Eleven iters have audited guards for *can it fire*; since iter-161,
for *can it still demonstrate that*. **Nobody has audited the accept side.** The most enumerable accept
clause in the family is the **waiver**: four checked-in files, ~20 entries, every one a standing
"this is fine."

## Hypothesis

A waiver that no longer matches any live site is an accept clause with no subject — dead weight that
also means the retraction it protects may be gone. Enumerating them is mechanical.

## Expected lift

A count of dead waivers, at a stated denominator, plus a fence keeping it at zero.

**Falsifiable, and this is the clause that fired:** if the waivers cannot be adjudicated without
re-implementing each guard's own matching, the audit is not mechanical and the hypothesis is wrong as
framed.

## Phase plan

- **A** — enumerate the waiver files; adjudicate each entry against its target.
- **B** — repair/remove the dead ones; fence.

## Escalation conditions

- The adjudication needs a guard's internal predicate → **that is the falsification**, not a detour.

## Acceptable close-no-lift outcomes

Every waiver is live, or the audit turns out to require the guards to report their own waiver usage —
which they do not. Either is a complete iter ending in characterization.
