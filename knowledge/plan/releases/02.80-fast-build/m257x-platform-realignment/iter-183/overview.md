---
iter: 183
milestone: M257x
iteration_type: tik
status: in-flight
opened: 2026-08-09
---

# iter-183 — the milestone's OWN backlog is a registry, and it is the only one with no fence

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* Route
disposition is mechanical: a route id either has a recorded closure or it does not, and a later bullet
either contradicts that closure or it does not. No sentence has to be interpreted.

## Step 0 — re-survey before targeting

`TOK-08` fixes the *order* (descending measured size), not the target. iter-182's close carried nine open
routes forward. Before picking one, the queue itself was measured — and the measurement is the target.

| derivation | value |
|---|---|
| distinct route ids under `iter-*/progress.md`, **line-scoped grep** | 207 |
| distinct route ids, **after joining hyphen-wrapped lines** | **204** |
| ids that exist ONLY as a line-wrap fragment (phantoms a naive grep manufactures) | **3** |
| routes appearing in at least one `**Routes carried forward:**` block | **188** |
| total (route × iter) dispositions in those blocks | **792** |
| routes carrying at least one `CLOSED` disposition | **47** |
| fences in `stack-core` whose subject is this registry | **0** |

**The queue is the substrate every sub-agent is briefed from**, and 792 dispositions across 188 routes
have been maintained entirely by eye. The orchestrator's own brief for this run flags one entry as
*"may already be closed — verify before working it"*, which is the consumer noticing the defect from
outside.

## Cluster / target identified

Two mechanically-decidable properties of the registry, both currently unasserted:

1. **Fragment-safety.** Markdown hard-wraps ids mid-token with a trailing hyphen, so `SURVEY-M257x-iter179-thirty-battery-` / `tests-unrun` reads as two things to any line-scoped reader. Measured: **3 of 207** ids a naive grep returns do not exist. Every prior statement about "the open routes" was made by a reader with this hazard live.
2. **Disposition consistency.** A route recorded `CLOSED` at iter K, then asserted `unchanged; open` at iter M > K with no re-open language, is a contradiction the registry publishes about itself. Raw detector: **6 sites**. Raw is not a finding — each is adjudicated individually before anything is published, per `D-M257x-122-3`.

## Hypothesis

The registry has a real, non-zero contradiction population; the raw 6 contains both true drift and
detector artifacts; and the class closes only by an **enumeration that keeps running** (iter-176's rule),
not by repairing the members found today.

## Expected lift

No `P`/`N` reading is taken this iter (`§9`: UNMEASURED, not unmoved). The deliverable is: the population
enumerated with its denominator stated, every raw candidate adjudicated with its verdict recorded, the
true defects repaired **in place by correction annotation** (never by rewriting a closed iter's record),
and a fence that enumerates the whole registry, ships green, and is registered in `guard_family`'s
`INVOCATIONS` in both directions.

## Phase plan

- **A — census + adjudicate.** Enumerate; adjudicate all 6 raw candidates one by one; publish the
  refuted ones as refuted.
- **B — repair.** Correction annotations at the contradicting sites, in the iter idiom already in use
  (`✅ CORRECTED — iter-NN`).
- **C — fence.** `route_disposition_guard.py` + mutation control + anti-vacuity control; `INVOCATIONS`
  entry; README index row.
- **D — close.** Re-run the scoped suite; record.

## Phase 0d — pre-flight tooling check (RUN — result recorded)

This iter adds a `*_guard.py`. `guard_family.py:303` makes a guard on disk with no `INVOCATIONS` entry
**exit 2**, checked in both directions. So the registration is a precondition of the iter, not a
follow-up — recorded here before any code is written so it cannot become the ninth vacuous-fence finding.

## Escalation conditions

- If the adjudication finds that a "contradiction" is actually the registry's **inability to express a
  half-closure**, that is a grammar gap (iter-180's class) and is fixed by giving it a grammar, not by
  rewriting the prose.
- If the fence cannot ship green without waiving a member, the waiver is named with its reason in the
  fence's own source — a silent skip is the failure mode `§5` rule 8 names.

## Acceptable close-no-lift outcomes

If all 6 raw candidates adjudicate as detector artifacts, the iter still closes complete: the population,
its denominator, the refutation, and a fence that keeps the property asserted are the deliverable. A
census that returns zero must prove its instrument (`§9`, iter-149) — the mutation control is that proof.
