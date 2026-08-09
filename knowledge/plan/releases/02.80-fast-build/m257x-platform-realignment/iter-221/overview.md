---
iter: 221
milestone: M257x
iteration_type: tik
status: in-progress
created: 2026-08-09
---

# iter-221 — the CORPUS cites 2,117 files and nothing checks that any of them exist

**Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey (mandatory)

iter-220 landed direction B (*a cited file exists*) over `rosetta-extensions`' 32 READMEs. Re-surveyed
immediately: **the same construct, on the surface clause 5 is actually about, is unfenced.**
`corpus_citation_guard` checks markdown **links** (C1 resolution, C2 anchors) — a **backticked
filename is not a link**, and 2,117 of them live in the corpus sources.

## Cluster / target identified

Backticked file citations in `fence_provenance.corpus_sources` — **2,117** across **114** documents:
876 `.md`, 428 `.go`, 173 `.sh`, 149 `.py`, and the rest. The `.md` half is decidable **today**,
because every pool it could resolve against is on this box.

## Hypothesis

The residual after a correctly-scoped pool is small, and every member of it is legitimate for a
**nameable** reason — which makes the class a taxonomy to declare rather than a defect list to repair.

## Expected lift

A running enumeration over the corpus's own file citations, with each non-resolving citation
**declared by class and size**, reconciled both ways.

## Phase plan

1. **Seal** the pool-scope readings, including this iter's own two wrong ones, before landing.
2. Adjudicate every non-resolving citation by hand.
3. Ship the census with the declared classes; prove it fires.

## Escalation conditions

- If a citation resolves nowhere and fits **no** declared class, it is a corpus defect: report it, do
  not widen a class to absorb it.

## Acceptable close-no-lift outcomes

A clean adjudicated zero closes this **`closed-fixed`** — the deliverable is the enumeration that keeps
running, provided it is proven to fire (`§9`).
