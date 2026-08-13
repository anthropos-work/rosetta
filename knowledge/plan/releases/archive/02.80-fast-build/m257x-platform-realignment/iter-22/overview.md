---
milestone: M257x
iter: 22
iteration_type: tik
status: archived
opened: 2026-08-01
---

# iter-22 — the enumerated clause-5 residual, executed

**Active strategy reference:** `TOK-01: instrument first, then follow` (milestone-root `decisions.md`) — the
only TOK on the chain; no triggered tok has fired. This tik is squarely inside it: iter-21's full read *was*
the instrument, and this iter follows what it measured.

## Step 0 — re-survey before targeting

Re-ran the open-of-iteration checks before committing to the hand-off's target list:

| check | reading |
|---|---|
| platform origin HEAD | **`2adcf71`** — unchanged since iter-19. Re-scope trigger stays at **occurrence 1 of 2** |
| `git rev-list --count HEAD..main` | **0** — the `CHECK-M257x-iter21-branch-behind-main` check, now in the checklist, reads clean |
| rext pin coherence | `main` @ `dc79b9d`, tag `fast-build-m257x-iter-20` on origin; `.agentspace/rext.tag` **and** the `stack-demo` consumption clone both match |
| all 21 hand-off anchors | present, at the quoted `file:line`, quotes verbatim |

Target is **NOT** substituted: `DOC-M257x-iter21-full-read-residual` is current and still the shortest path
to clause 5.

## Cluster / target identified

The 21 enumerated clause-5 blockers from iter-21's full read. Each arrives with `file:line`, a verbatim quote,
a refuting citation and a one-line correction — deliberately authored to be mechanical.

**But mechanical is not the same as trusted.** Nine consecutive hand-offs in this milestone have been refuted
or materially corrected on re-measurement. So each correction is re-derived against platform source
(`docker-compose.yml` / `repos.yml` / `app/main.go` at origin HEAD) *before* it is applied — not just the
anchor, the **correction**.

## Hypothesis

Applying the 21 corrections + the re-audit closes clause 5 (KB-fidelity GREEN, or YELLOW with 0 blockers)
over `corpus/services/**` + `corpus/architecture/**`, taking the gate from **3 of 5** to **4 of 5** and
leaving clause 2 as the sole residual.

## Expected lift

+1 gate clause. Secondary: the correction set is itself audited, so a wrong correction is caught here rather
than shipped as a confident falsehood.

## Phase plan

1. Re-verify all 21 anchors at `file:line` (cheap; fails loudly if any moved).
2. **Re-derive each correction against platform source.** Any that does not survive is re-graded, not applied.
3. Apply as an enumerated sweep (`(file, old, new)` tuples, each asserting `old` occurs **exactly once** —
   0 fails loudly, 2+ fails loudly). Reuses the iter-20/21 harness shape.
4. Run both fences: `corpus_index_guard.py`, `platform_alignment_guard.py`.
5. **Re-audit by FULL READ of all 40 in-scope files**, fanned across sub-agents, with a positive-control
   line-count per file. Per `platform-alignment.md` §5, a term-scoped audit measures the terms, not the
   corpus — it is not admissible for closing this clause.
6. Fix whatever the re-audit returns; close.

## Escalation conditions

- A second platform commit lands → re-scope trigger FIRES (occurrence 2 of 2) → exit `re-scope-trigger`.
- The re-audit returns blockers that need a **platform-repo edit** to fix → user-blocker (v2.8's zero-edit
  constraint is binding).
- The re-audit returns a residual too large to close in-iter → land what is correct, route the rest with a
  named handler, close `closed-fixed-partial`.

## Acceptable close-no-lift outcomes

If the re-audit shows the 21 were substantially mis-derived — i.e. the corrections make the corpus *less*
true — that falsification is itself the deliverable, and the iter closes `closed-no-lift` with the residual
re-enumerated from evidence rather than inherited.
