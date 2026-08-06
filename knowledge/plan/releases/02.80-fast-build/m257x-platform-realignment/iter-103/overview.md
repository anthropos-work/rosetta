---
iteration_type: tik
iter_shape: reading
status: closed
opened: 2026-08-06
closed: 2026-08-06
outcome: "N = 33 — the pre-sealed rule's >=23 branch: THE BURN-DOWN LEG DOES NOT REACH THE RESIDUAL. 22 predicates then, 22 now; anchors 24 -> 33. Bands 4 HELD of 10. Gate unchanged at 4 of 5; clause 5 open."
active_strategy: TOK-04
---

# iter-103 — the reading that asks whether the burn-down leg works

**Shape:** tik · **the MEASURING pass — no repair inside it** · `iter_shape: reading`

## Active strategy

`TOK-04` (P1/P2 discipline: re-derive ground truth at every open; measuring and repairing never share a
pass). No re-survey substitution: the TOK-directed next target — *re-read over the repaired tree* — is
exactly what the evidence still names, and iter-102's close routes to it explicitly.

## Why this iter exists

iter-101's band #3 was the governing result of the milestone. Blind overlap with iter-99's published 28,
matched on predicate, came out **6** against a pre-registered **[14, 22]** — a fail LOW, meaning the two
readings are **more independent than Chapman assumes**, so every prior estimate was biased low.
Cross-reading Chapman gives **N̂ ≈ 102.6, itself a floor**.

**The right reading of the series 16.7 → 29.4 → 45.2 → ~103 is four successive corrections to an
underestimate, not four measurements of a growing pool.** The pool was probably always ~100.

iter-102 then paid **both** outstanding unions in one pass — 52 anchors → 98 sites, 94 repaired — roughly
half the estimated residual, with machine-graded reach of 80.4 % / 80.6 % against the full booked sets and
effectively 100 % against the upheld sets.

**So this reading answers the question that decides everything after it:** now that ~94 sites have been
repaired against a residual estimated at ~100, **does the burn-down leg actually work?** A large drop says
the pool is draining and an ETA becomes derivable. **A flat or rising `N` says repair is not reaching the
residual — and that is a far more important finding than the number.**

The verdict rule is **pre-registered before any seat is dealt** (`pre-registration.md`): `N ≤ 16` works ·
`17–22` ambiguous · `≥ 23` does not reach.

## What makes this reading different from the four before it

**It is NOT a replicate, and the sheet says so first.** iter-102 grew the in-scope corpus **10,278 →
10,646 lines (+3.6 %)**, which is far more than the `+2` that let iter-101 be a replicate on a fixed
subject. The greedy LPT partition is therefore **recomputed** and comes out different. The partitioning
*algorithm* is proven unchanged — the same script reproduces iter-101's published partition exactly when
run over the file sizes at `8f04d3a`. **The instrument is the same; the subject moved under it.**

Two operational changes, both disclosed, neither touching the instrument:

- **14 seats dealt in two batches of 7.** iter-101 dealt all 14 at once and lost `r24-D` to a spend limit,
  which made that reading a 13-seat union and forced every cross-reading quantity to be restated on a
  6-seat common subject. Batching is an operations change, not an instrument change.
- **The ADDENDUM now names which `rosetta-extensions` tree settles a claim** (§5 rule 45), while line 37
  of the frozen instrument still names the wrong one. Band #6 measures whether an addendum can repair a
  defect in a frozen instrument without editing it.

## The known instrument defect, delivered unfixed for the third reading

`briefing-iter76-AS-RUN.md:37` names the rext **authoring** copy as "the tooling"; a claim about what a
stack runs is settled in the **pinned per-stack clone**. **It is not edited.** Editing it would break the
comparability the series exists to establish. `N` is post-adjudication and this class is caught by
adjudication, so `N` is uncontaminated; the **upheld rate is not**, so it is reported twice — raw, and with
the `wrong-tree` class separated. Routed as `DEF-M257x-iter101-briefing-rext-tree`.

## Exit condition for the iter

The reading returns a number. `N = 0` meets clause 5; anything else leaves the gate where it is and routes.
No other reading of clause 5 is available — four user rulings.

**Clauses 1 and 2 cannot move in this pass and nothing here tries to move them.** They are a close blocker
(`D-M257x-102-3`) owned by Lane B, which has cycle 1 green and the Playthrough suite at 30/0/0 on platform
`0c91421`. This iter brings up no stack, touches no `stack-demo/**` path, and — per §5 rule 41a — **fetches
nothing**.
