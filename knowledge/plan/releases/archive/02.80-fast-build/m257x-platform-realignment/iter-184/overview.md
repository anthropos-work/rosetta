---
iter: 184
milestone: M257x
iteration_type: tik
status: in-flight
opened: 2026-08-09
---

# iter-184 — a carry-forward that names a SET must enumerate it, and this one denotes the empty set

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* iter-183
built the first fence over the route registry and asserted **one** property. This iter works the same
population's next mechanical property: a carried route must be a **route**.

## Step 0 — re-survey before targeting

iter-183's own close routed `SURVEY-M257x-iter183-only-ONE-registry-property-is-asserted`. Probing that
residual immediately turned up a member of the fence's own population that is not a route:

| finding | value |
|---|---|
| routes the iter-183 fence enumerates | **185** |
| of those, members that are **not routes** | **1** — `SURVEY-M257x-iter181-` |
| its source | iter-182's carry-forward bullet, written `` `SURVEY-M257x-iter181-*` `` — a **glob** |
| routes iter-181 actually created | **0** |

**So the glob denotes the empty set**, and it has read as a live backlog item ever since — including in
run 17's orchestrator brief, which lists it. `§5` rule 73 already says **a glob is not a derivation**;
what iter-183 added is a population that can be checked against it.

## Cluster / target identified

The registry's second mechanical property: **every enumerated member is a well-formed route id**. Two
ways it currently is not — a glob remnant (`-*` truncated by the id grammar) and a line-wrap fragment
(measured at iter-183: a naive reader sees **207 ids where 204 exist**). Both are decidable.

Separately measured and deliberately **not** fixed here: the coarse wildcard `"The standing queue,
unchanged"` appears in **17 of 165** carry-forward blocks. It is a policy construct, not a malformed id.

## Hypothesis

The glob is the only non-route member today; the class is closed not by removing it but by an assertion
that keeps running, and the assertion is cheap because iter-183 already derives the population.

## Expected lift

No `P`/`N` reading (`§9`: UNMEASURED, not unmoved). Deliverable: the glob adjudicated and repaired in
place, an arm on `route_disposition_guard` that refuses a malformed member with its site, RED-proven
against the unrepaired tree first, and the standing-queue wildcard routed **with its measurement** rather
than as a mood.

## Phase plan

- **A — census** the malformed members and the wildcard blocks.
- **B — adjudicate + repair** iter-182's glob by in-place correction.
- **C — fence** the well-formedness property; RED-prove it against the pre-repair tree.
- **D — close.**

## Phase 0d — pre-flight tooling check (RUN)

iter-183's own lesson applies to itself: *a pre-flight that finds one precondition has not established
there is only one.* The new arm goes into the **existing** `test_route_disposition_guard.py` and the
**existing** guard, so no new `*_guard.py` appears on disk — which means no `INVOCATIONS` entry, no
provenance stamp, and **no movement of the README index triple**. That is the reason to extend rather
than add, and it is checked rather than assumed: the two registry guards that fired at iter-183 are
re-run at close.

## Escalation conditions

- If the glob turns out to denote a **non-empty** set, this is a dropped-routes finding, not a
  formatting one, and the repair is an enumeration — not a deletion.
- If refusing malformed members would make the fence RED on members it cannot repair, the members are
  named with reasons, never skipped silently (`§5` rule 8).

## Acceptable close-no-lift outcomes

If the glob adjudicates as a harmless shorthand with a non-empty, already-carried denotation, the iter
still closes complete: the property is enumerated, the measurement published, and the assertion shipped.
