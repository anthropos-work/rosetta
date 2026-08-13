---
iter: 66
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-04
---

# iter-66 — root `CLAUDE.md`, against what this session measured

**Active strategy reference:** `TOK-05`. Corpus-only; no rext change.

## Cluster / target identified

Root `CLAUDE.md` — a long-standing routed item, and the highest-traffic file in the corpus: it is what
every agent reads before doing anything. iter-62 repaired its profile table; iters 63–65 then measured
two facts it still states wrongly, and neither is reachable by any current fence (both are prose about
*which tier a service is in* and *which RPC edges are live*, not a profile token or an address value).

## Hypothesis

Two specific, artifact-backed corrections, both of which change what a reader would *do*.

## Expected lift

Zero fence findings (there were none to begin with — this is prose the fences cannot see), and two
false statements removed from the file agents read first.

## Phase plan

A. Re-derive both facts from the artifacts (already measured this session — cite, do not re-assert).
B. Edit; re-run every corpus guard + the §5 rule 34 re-point.

## Escalation conditions

- Anything requiring an rext change → route forward; this iter is corpus-only by design.

## Acceptable close-no-lift outcomes

- Both statements turn out already correct → record the falsification.
