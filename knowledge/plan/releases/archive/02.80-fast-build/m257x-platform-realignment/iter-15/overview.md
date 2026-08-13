---
milestone: M257x
iter: 15
iteration_type: tik
status: closed-fixed-partial
opened: 2026-08-01
---

# iter-15 — gate clause 2: the full Playthrough suite on the clause-1 stack

**Active strategy reference:** `TOK-01` — *instrument first, then follow*. Step 5 of its ordering is
*"prove it cold"*; clause 1 did that for the bring-up, and clause 2 is the same proof carried one layer up —
a green bring-up must not be allowed to mean an empty world.

## Step 0 — re-survey before targeting (mandatory)

Run before committing to the hand-off's target. Three claims re-measured:

1. **The denominator is still 30.** `grep -hoE 'playthrough: pt-…' manifest/*.yaml | sort -u` = **30**;
   `grep -hoE '@pt:pt-…' e2e/tests/*.spec.ts | sort -u` = **30**; `diff` of the two sorted sets is
   **identical**. (The first cut of this count returned **0** — I grepped `^\s+- playthrough:` with a
   leading-dash anchor that the field does not carry. That is §5 rule 3, the exact incident the protocol
   already records for this very number. Corrected by opening a manifest.)
2. **The hand-off's router warning is a COMMENT, not a break.** `run-playthroughs.sh:77` does name
   `https://<magicdns>:<15050+offset>/graphql` — inside the `--public-host` explainer block, not in any
   executed line. A whole-section grep for `5050|graphql` across `playthroughs/` returns **5 hits, all
   comments or doc-strings** (2 manifest provenance notes, 2 spec/page-object doc-strings, this one). The
   Playthroughs drive the **browser** against next-web, whose GraphQL origin is baked at build time — the
   surface iter-13 re-pointed and iter-14 proved live. So the routed-forward warning is real but is a
   **doc** item, not a plumbing one. Re-measured rather than trusted: five inherited hand-offs in a row
   have been refuted, and this is the sixth to be checked.
3. **demo-1 is up and serving.** app `:13000` 307 · hiring `:13001` 307 · studio `:19000` 302 · academy
   `:13077` 200 · cockpit `:17700` 200 · backend GraphQL `POST :18082/graphql/query` **200**. The last one
   matters: it is iter-13's re-pointed address AND path answering live.

**Target confirmed, unchanged:** clause 2.

## Cluster / target identified

Gate clause 2 — *"the full Playthrough suite passes on that stack (30 live / 0 failing / 0 error) — presence
AND function, so a green bring-up cannot mean an empty world."* It is the only remaining clause measurable on
the live stack, and the stack is up **now**; clauses 3 and 5 are corpus work that needs no stack at all.
Running it first is forced by that ordering, not preferred.

## Hypothesis

The suite has not been run since the router drop. Two outcomes are informative and neither is a surprise:
either it passes (clause 2 falls, and the bring-up's green is corroborated by function), or it fails and the
failures name the surfaces the bring-up does not exercise — which is what seven consecutive iters have found
in the wider run.

## Expected lift

Gate clauses 2/5 → 3/5, or a measured failing set with named handlers. A partial pass is **not** a clause;
the clause is 30/0/0.

## Phase plan

- **A — pre-flight (done, Phase 0d).** Denominator, router-reference sweep, live surfaces, harness deps.
- **B — reset-to-seed.** `run-playthroughs.sh 1 --reset` — the real `stackseed --reset` + `pt-world.seed.yaml`,
  never an additive re-seed. This replaces demo-1's demo world with the decoupled Playthrough world; clause-1
  evidence is already checked in at `evidence/av-cycle{1,2,3}.json`, so nothing is lost.
- **C — run the full suite unscoped.** Unscoped on purpose: the ptreport gate is **binding** on a full run and
  **advisory** on a scoped one (`run-playthroughs.sh:300-307`), and the residual documented at `:292` says a
  broad `--grep '@pt'` would be graded advisory. A scoped run cannot satisfy this clause.
- **D — triage.** Route each failure by surface, per `coverage-protocol.md`'s fix-surface routing table.
- **E — close.** Record the measured four-state map as the clause-2 evidence.

## Escalation conditions

- A second platform commit invalidating an alignment attempt → **re-scope trigger fires** (occurrence 1 of 2);
  stop and escalate, do not re-point a third time.
- A failure whose only fix is a platform-repo edit → escalate; `demopatch` should make it unnecessary.
- Suite red in a way that needs a design decision → route forward with a named handler, close on what landed.

## Acceptable close-no-lift outcomes

A measured, triaged failing set with named handlers **is** a complete iter under the protocol: it converts
"unknown" into "known and routed", which is exactly what the milestone exists to do. What is not acceptable is
reporting a number the run did not produce, or grading a scoped run as a full one.
