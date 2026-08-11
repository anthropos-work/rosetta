---
iter: 143
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-08
active_strategy: TOK-08
route_closed: SURVEY-M257x-iter143-bare-orphan-bucket
---

# iter-143 — the `(bare)` orphan bucket, censused; head inference measured and REFUSED

**Active strategy reference:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop
sampling them*. `TOK-08`'s pre-registered class list has two classes. Class 1 (intra-corpus citation
resolution) is fenced GREEN. **Class 2 (platform-source citation resolution) is fenced at 59.1 %
reach** and the largest single unresolved head is `(bare)` at **384**. This iter closes
`SURVEY-M257x-iter143-bare-orphan-bucket`, the route iter-142 recorded at its close as the next
iter's starting point (`4edad03`).

**Step 0 — re-survey (mandatory).** Re-ran `anchor_construct_guard --repo-root .` at iter open:
**861/1456 = 59.1 %, `(bare)` x384**, identical to the figure `4edad03` recorded. Target live and
unabsorbed; no substitution needed.

## Cluster / target identified

`anchor_construct_guard` reports `(bare)` as **one opaque number**, with a source comment asserting
*"this is the bucket the ports are in, and resolving it would be the 134-findings first draft
returning."* `4edad03` already measured that comment as **only marginally true** (38 of 621 bare
matches carry a known port). What nobody has is the **decomposition** — and without it, *"head
resolution first"* (`FIX-M257x-iter138-anchor-rot-fence`'s re-specification) has no denominator to
work against.

## Hypothesis

The orphan bucket is dominated by anchors whose file is named as a **bare file mention** rather than
as a `path:NNN` citation — the corpus's *"all anchors in `handler.go`: … (`:1458`)"* construct. If so,
a **mention referent** would place a large share of them, and `_FILE_MENTION` — a regex the guard
**already has**, used only to BREAK a chain — supplies it without a new heuristic.

## Expected lift

Reach ≫ 59.1 %, with every newly-admitted anchor hand-read before publication.

## Phase plan (three planned lines — declared, per the scope-creep carve-out)

1. **Census** the 384 refusals by REASON, with the denominator stated (iter-114's rule).
2. **Derive** the head rule from the census, then **hand-read 100 %** of the population it admits
   *before* anything ships (iter-142's `D-M257x-142-1`, in its own words).
3. **Ship only what the audit supports** — plus the census decomposition itself, so the bucket stops
   being one opaque number.

## Escalation conditions

- Precision of the derived rule below fence quality → **do not ship it**; publish the refutation with
  its numbers and route the residual. iter-138 published an unaudited mechanical predicate over this
  exact population and iter-139 retracted it 0-for-12; that must not recur.
- A tuned numeric constant is the only thing separating true from false → **do not ship it**. Two such
  constants are already routed open (`-iter142-path-arm-window`, `-iter142-tier-b-underflag`); a third
  is a pattern, not a fix.

## Acceptable close-no-lift outcomes

**A measured refusal is a first-class outcome here.** If head inference over this bucket cannot reach
fence quality, saying so *with the precision number and the mechanism breakdown* answers the standing
`FIX-M257x-iter138-anchor-rot-fence` re-specification — which asked for head resolution *first*, and
is entitled to an answer of "not mechanically decidable at this quality, and here is why."
