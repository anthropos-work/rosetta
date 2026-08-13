---
iter: 70
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-04
closed: 2026-08-04
---

# iter-70 — the class I routed one iteration ago was twelve ports and four citations

**Active strategy reference:** `TOK-05`, step 2 (citations) — clearing the last named prerequisite
before step 4, the graded read.

## Step 0 — re-survey before targeting (and it re-scoped the iter immediately)

iter-69 routed **`FIX-M257x-iter69-pathless-antecedent` — "23 citations across 14 lines"** whose
antecedent is a backticked path with no line number, and called it *"a prerequisite for the graded
read."* Re-derived at this open, six of the 23 were `shared_libraries.md:70`'s, repaired in iter-69
itself, leaving **17 across 13 lines**.

Then the values were read rather than counted: `:5050`, `:8082`, `:3000`, `:8000`, `:9000`, `:7070`.
**They are PORT NUMBERS.** *"GraphQL is served by `backend` itself at `:8082/graphql/query`"* is not
a citation into `docker-compose.yml`; it is a URL.

Discriminated mechanically — **does the antecedent file even HAVE that many lines?**

| | n |
|---|---|
| impossible as a line (`repos.yml` has 31 lines; `:5050` is a port) | **12** |
| line-plausible | **5**, of which one (`labsapi/client.go` *"default `:7070`"*) is a port my own `max()`-over-basenames heuristic mis-sized |

**The routed class is 4 citations across 3 lines**, not 23. `platform-alignment.md` §5 rule 32 —
*re-derive the hand-off's numbers* — applied to my own hand-off, one iteration later.

## Cluster / target identified

Grade the 4; repair what is broken; and settle the **fence** iter-69 routed beside them
(`FENCE-M257x-iter69-citation-antecedent`), because the derivation above is exactly the evidence that
decides whether it should be built.

## Hypothesis

The routed FIX closes as largely falsified, and the routed FENCE **must not be built as designed** —
`path` + bare `:N` is ambiguous between a line and a port, and the corpus uses the port sense 3× more
often. A fence over it would be a false-positive generator, which §4 Trap A says is not a fence.

## Expected lift

Clause 5's last *named* blocking prerequisite is removed or shown never to have been one — so the
graded read is not held behind a phantom.

## Phase plan

- **A** — re-derive the routed class (**done at open**).
- **B** — discriminate line-from-port mechanically, and state the discriminator's own defect.
- **C** — read and adjudicate the survivors against their sources.
- **D** — repair; settle the routed fence with a recorded falsification.
- **E** — gates.

## Escalation conditions

A citation that cannot be adjudicated at a **ref** (the source repo is not cloned) is recorded
UNMEASURED with a named handler — **never patched on speculation**, per iter-68's precedent of
recording two unreachable boundary defects rather than guessing at them.

## Acceptable close-no-lift outcomes

All four holding would be a complete iter: the deliverable is then the **falsification of a routed
prerequisite**, which is worth more to the read than four edits would have been.
