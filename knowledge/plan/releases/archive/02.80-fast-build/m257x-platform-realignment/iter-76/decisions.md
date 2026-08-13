# iter-76 — decisions

## `D-M257x-76-1` — clause 5 is NOT met, and the distance was never 4-of-5-minus-a-little

Two blind readings of the identical partition returned **77 and 75 blockers**. The gate reads *"a
reading that returns zero"*. It returned neither zero nor anything near it.

The milestone has stood at **4 of 5** for eighteen iterations with clause 5 described as a *residual*.
This reading measures that residual for the first time since iter-53, against a corpus the platform
moved under twice (the router deletion, then the storage/messenger fold). **It is not a residual; it
is the largest open class in the milestone.** Recording that is the iteration's deliverable.

## `D-M257x-76-2` — five GREEN guards and ~150 findings are not in contradiction, and the guard said so all along

`platform_predicate_guard` prints on every run:

```
G2 3 repo-count claim(s); G5 24 migration claim(s) = 1 enumerated + 21 free prose UNREACHED
```

**G5 reaches one claim in twenty-four.** The corpus's false predicates are in the twenty-one it
cannot see. Every guard was honest; every guard was green; the corpus was wrong the whole time.

> **Every GREEN verdict is a statement about reach.** This iteration is that sentence with a
> measurement attached: 5 guards green, ~150 blockers, no contradiction.

The consequence for how this milestone reports progress: **a fence's GREEN may never again be quoted
without its reach line beside it.** "Five corpus guards OK" has appeared in every iter close since
iter-59 and, read alone, it implied something it never claimed.

## `D-M257x-76-3` — the routed count is a hypothesis, and this one already has a proven false-positive class

`r13-F` B13/B14 (and sibling findings in other seats) book `docker-compose.yml:311` / `:362` /
`:352` / `:337-341` as past the end of a 271-line file. Traced: **both sit in blocks pinning
`2adcf71`, where the file is 387 lines**, and one names its pin in words (*"since platform
`2adcf71`"*). §5 rule 33 settles them as **correct**.

So ~150 is an **upper bound**, not a work item. iter-77 adjudicates before repairing — the fifth
time this milestone has had to write that sentence, and the first time it has been written *about a
reading* rather than about a routed backlog.

**The instrument lesson, which is mine and not the seats':** the briefing *told* the seats that a
claim is settled at the ref it names, in a section of its own, and seats still graded dated claims
against the checkout. **A rule stated in a briefing is not a rule enforced by an instrument.** The
same predicate is already mechanical inside `anchor_construct_guard.block_ref` — the read had no
access to it. Routed as `CHECK-M257x-iter76-seat-ref-discipline`: give the seats the resolved ref
per citation, or expect this class in every future reading.

## `D-M257x-76-4` — repair by predicate, and not in the iteration that discovered the class

Pre-registered escalation (*"more than ~15 union blockers → measure and route"*) fired at ~150.
§5 rule 19: **half-repairing a uniformly-wrong corpus is worse than leaving it.** The repair is
routed whole, as `FIX-M257x-iter76-read-union`, under `D-M257x-59-1`'s predicate unit — six
predicates cover most of the ~150, and every one of the six has a **legal set already derived and
printed by `platform_predicate_guard` on every run**:

| predicate | legal set, derived |
|---|---|
| which containers the default profile starts | `core` → 5: backend · gotenberg · postgresql · redis · sentinel |
| how many services compose declares | 8, or 10 with the `include` |
| which profile tokens are legal | 8, and `graphql` is not among them |
| how many repos `repos.yml` clones | 6 |
| whether `STORAGE_RPC_ADDR` is read | 0 read sites at the ref claimed |
| whether a husk container exists | no compose service for cms / jobsimulation / roadrunner |

**The repair is not complete until the reach hole is closed.** G5 reaching 1 of 24 is *why* these
claims survived seventy-five iterations of a milestone whose entire subject is this fold. Repairing
the sites and leaving the fence at 1-of-24 buys a corpus that is correct today and equally blind
tomorrow.

## `D-M257x-76-5` — one number in this iteration disagreed with the tested instrument, and I am recording it as unsettled

A quick `grep -cE '^  [A-Za-z0-9_-]+:$'` over `docker-compose.yml` returned **9** service-shaped
lines, while `platform_predicate_guard`'s tested parser derives **8 (+2 in `common.yml`)**. Multiple
seats independently booked *"nine services"* as **false**, saying 8 or 10.

Per §5 rule 32, **the instrument that agrees with nothing is usually the new one** — mine, written
in one line during a close. But *"8 vs 9 vs 10"* is exactly the sort of denominator the corpus keeps
getting wrong, and settling it by assertion here would be the error this decision is warning about.
**Explicitly left unsettled and handed to iter-77** as the first thing to derive, before any site
naming a service count is touched.
