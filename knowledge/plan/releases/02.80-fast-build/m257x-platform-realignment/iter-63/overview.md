---
iter: 63
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-04
---

# iter-63 — the citations, and the hole a ref-pin leaves

**Active strategy reference:** `TOK-05` — *stop repairing claims; fence the predicates under them*.
Step 1 (fence the profile predicate) landed across iters 60–62. **This is step 2: citations**, plus the
structural hole iter-62 promoted (`D-M257x-62-2`).

## Cluster / target identified

Two routed items, in the order the briefing gives them:

1. **`FIX-M257x-iter58-mainline-shift`** — the corpus citations the iter-58 pin advance moved. Two prior
   sizings disagree in **both** numerator and denominator: iter-58 read *"21 of 22 moved"* (protocol §7
   rule 4 records *"22 of 23"*), iter-61 read *"5 of 16 distinct `app/main.go:N` citations still landing"*.
   **Re-measure first**; name which figure was wrong and why.
2. **`CHECK-M257x-iter60-stale-pin-exemption`** — promoted and structural. A ref-pin exempts a claim from
   the fence, so `messenger.md:107-110` and `service_taxonomy.md:55-67` sit behind `@ 2adcf71` pins the
   guard correctly cannot reach. Decide what a pinned claim owes the reader; then fence it.

## Hypothesis

The citation class and the pin-exemption class are **the same defect wearing two costumes**: both are
claims whose truth is indexed to a ref, held in a corpus that reads as present-tense. The citation half
is caught by nothing (§7 rule 4's measured **4.5%** fence catch-rate); the pin half is *deliberately*
skipped by the fence that would otherwise catch it. Repairing one without the other leaves the corpus
half-dated.

## Expected lift

- The mainline citation class re-measured against `app` @ **`b948604` v1.366.0** and repaired to zero
  moved-and-unpinned sites, with the real denominator stated.
- A decision on ref-pinned claims, recorded, and whatever it implies fenced.

## Phase plan

A. Re-measure the citation class from a derived enumeration (both constructs: `path:line` **and** the
   bare `:N` continuation form the corpus uses). Grade against the app clone at HEAD.
B. Repair, adjudicating every claim against **platform artifacts** (`docker-compose.yml`, `repos.yml`,
   `app` source) — never against another corpus document (§5 rule, iter-21).
C. Apply §7 rule 4's own citation-safety half to **this iter's corpus edits** — a corpus repair moves the
   corpus's own line numbers exactly as a pin advance moves the platform's.
D. The ref-pin decision + fence + repair of whatever it un-exempts.
E. Re-measure, run the guard suites, close.

## Escalation conditions

- A platform artifact that contradicts the map → adjudicate against the artifact, record, continue.
- Un-exempting ref-pins produces a blast radius large enough to make the fence cry wolf → do NOT tune a
  threshold to the answer key (§4 Trap A); narrow by **structure** or route the residual whole.

## Acceptable close-no-lift outcomes

- The prior figures turn out to be right and the class is already clean — recorded as a falsification.
- The ref-pin rule turns out to be correct as written, with the two holes explained by a different
  mechanism — recorded, with the real mechanism named.
