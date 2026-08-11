---
iter: 42
milestone: M257x
---

# iter-42 — progress

**Type:** tok (triggered semantics; **user-directed** trigger — the 3-no-prog streak did not fire)

Shape declared per `build-mstone-iters` Phase 0 rule 2. Phase 0b **skipped**: the revision redirects into no
subsystem the milestone's standing KB-fidelity verdict does not already cover — it stays inside the corpus
(which *is* clause 5) and inside `stack-core`'s guard layer, both audited continuously since iter-01. Phase
0d **skipped**: a tok authors strategy and wires no artifact through a pipeline.

## Phase A — re-survey (measured at open, not inherited)

- **In-scope corpus byte-identical to what iter-41 measured.** `git diff 103ad31..HEAD -- corpus/services/
  corpus/architecture/` empty, and `103ad31` *is* iter-41's own closing commit. **The 18 stand.**
- **Spot-verified four blockers live** rather than trusting the ledger: #18 (`architecture_overview.md:243`
  still publishes *"EU-first routing — Azure OpenAI EU → Azure OpenAI US → direct OpenAI"* against
  `external_services.md:537`'s *"There is **no** ordered EU-first fallback chain"*), #15
  (`roadrunner.md:23-25` asserts jobsimulation was removed from `repos.yml` + `docker-compose.yml`; it is at
  **`repos.yml:17`** and **`docker-compose.yml:83`** in the platform clone), #9 (`service_taxonomy.md:136`
  *"not in main docker-compose"* against `:75` in the same file), #8 (`:145` lists **React** against
  `studio-desk.md:20`'s *"vanilla TS frontend, no framework"*). All four present, all four contradicted by
  the twin the ledger names.
- **Gate 4 of 5.** Platform origin `2adcf71` unchanged → re-scope trigger stays at occurrence **1 of 2**.
- **rext `main` @ `069c238`**, clean, pin `fast-build-m257x-iter-37` on origin, both pins match.

## Phase B — the classification

The 18 split by the cheapest instrument that could have caught each (row-by-row out of iter-41's
`blocker-ledger.md`): **13 corpus-contradicts-itself · 3 anchor-resolves-but-names-the-wrong-construct ·
2 derived-scalar-vs-source.** And of the **9 repair-induced**, **8 are self-contradiction** — one mechanical
class, detectable with no platform read at all. At least **4 of the 18 are RETURNING claims** with a verdict
already recorded in a prior blocker-ledger, and nothing checks for that.

Full derivation, weaknesses stated, in [`overview.md`](overview.md); the reasoning in
[`decisions.md`](decisions.md) `D-M257x-42-1` / `D-M257x-42-2`.

## Phase C — TOK-02 authored

Appended to the milestone-root [`decisions.md`](../decisions.md) as
**`TOK-02: fence the prose the way the anchors are fenced`**.

## Phase D — close

## Close — 2026-08-02

**Outcome:** the milestone's first strategy revision in 41 iterations. TOK-01 is **extended, not replaced**:
its derive-and-fence principle — which holds clauses 1–4 — is applied to clause 5, the one surface where the
milestone reverted to hand-maintenance and the one clause still open. The revision rests on a classification
of iter-41's 18 blockers by *cheapest reaching instrument*: **13 of 18 are the corpus contradicting itself**,
and **8 of the 9 repair-induced blockers are that same single mechanical class** — a claim repaired at one
site and left standing at another. So the ~50% induction rate that iter-41 escalated on is a property of the
**repair method**, not a law about the corpus, and it fails in a way that needs no platform knowledge to
detect. The audit instrument is untouched; clause 5 is untouched; the 18 are repaired, not deferred.
**Sequencing is load-bearing: the fence is built and reddened BEFORE any repair, because today's 18-defect
corpus is the only test fixture with a known answer key and it is perishable.**
**Type:** tok
**Status:** closed-fixed *(the revised strategy plus its justifying measurement is the planned deliverable;
`overview.md` committed to authoring TOK-02 and forbade repairing corpus text — no corpus text was written.)*
**Gate:** N/A for tok — stands at **4 of 5**; clause 5 unchanged at 18 by construction.
**Phase 5 grading:** (1) gate-met: n — (2) **triggered-tok: y** — (3) re-scope: n (platform origin `2adcf71`
re-fetched at open, unchanged; occurrence stays 1 of 2) — (4) user-blocker: n — (5) cap-reached: n (0 tiks) —
(6) protocol-stop: n — Outcome: **exit-2**
**Decisions:** `D-M257x-42-1` … `D-M257x-42-5`, plus `TOK-02` in the milestone-root `decisions.md`.
**Side-deliverables:** none. **Zero words of in-scope corpus text written**, and no fence built — a tok
authors strategy; building the fence inside the tok would have consumed the fixture `D-M257x-42-3` exists to
protect.
**Routes carried forward:**
- **`FENCE-M257x-iter42-claim-twin`** — **the next tik.** Derive the claim ledger from the five existing
  blocker-ledgers (never by hand — §5 rule 19's list-derivation clause), build the tree-wide claim-twin
  fence, **watch it go RED** on today's corpus, then GREEN. Repair nothing in that tik.
- `FIX-M257x-iter41-blocker-set` — the 18, anchored and ready, now sequenced **after** the fence (step 4 of
  TOK-02) rather than before it.
- `CHECK-M257x-iter33-derived-fact-fence` — **subsumed by `FENCE-M257x-iter42-claim-twin`**, which is the
  correct cut of it: a claim ledger with verdicts, not a general fact-checker.
- Unchanged and still open: `CHECK-M257x-iter41-tenancy-fence-fifth-failure`,
  `DOC-M257x-iter41-ops-collateral`, `DOC-M257x-iter41-minors` (~85),
  `CHECK-M257x-iter38-ai-act-classification` (still needs an owner outside this milestone),
  `FIX-M257x-iter37-dev-twin-has-no-fallback`, and the standing `RF-2`/`RF-3`/`RF-7`…`RF-12` harden queue.
**Lessons:**
1. **An undifferentiated count invites an undifferentiated remedy.** Six passes reported a single number and
   every strategy discussion — including the escalation — reasoned about that number as one quantity. The
   same 18, split by *which instrument could have caught them*, offered a move that the aggregate hid
   completely. Classify before concluding a residual is irreducible.
2. **Measure what the repair induces — then measure what KIND it induces.** §5 rule 20 (iter-41) got the
   milestone to the induction rate. The decision-relevant quantity was one step further on: the induced
   defects are **homogeneous**, and homogeneity is what makes mechanization available. A rate says *stop*;
   a class says *what to build*.
3. **A strategy that works everywhere except one surface is not a failed strategy.** TOK-01 held four gate
   clauses on derived, RED-watched fences. Clause 5 is the one place it was never applied — and reading the
   milestone as *"40 tiks, one strategy, stuck"* obscured that the strategy had simply never been pointed at
   the open clause. Revise by asking where the working principle is *absent*, not by replacing it.
4. **The fixture is perishable.** A corpus with a known, anchored answer key exists exactly once, and
   repairing it destroys the only thing that can falsify the fence built to protect it. Sequence the fence
   first — the same pre-commitment logic iter-41 used to protect its reading.
