# iter-214 — the claim fences read the corpus; the TOOLING's prose is outside all of them

**Type:** tik — under [`TOK-08`](../decisions.md).

See [`overview.md`](overview.md) for T1–T6 + the stop condition, sealed at `3f81916` before any repair.

## What was measured

Every figure below is the arms' own live output, not a transcription:

- **T1/T2 — CONFIRMED.** 264 adjudicated claims · the guard's population **114** documents · outside it
  **186** rext Python modules and **211** rext markdown documents.
- **T3 — CONFIRMED.** fixtures **138 docs / 217 hits** · test modules **1 / 0** · **real prose 72 / 0**.
- **T4 — CONFIRMED.** The zero proves its instrument **in the same run**: 217 findings over the fixture
  bucket from the identical matcher and claim set.
- **T5 — CONFIRMED.** Python: **186 modules, 10 findings, 7 outside a retraction context.**
- **T6 — CONFIRMED and load-bearing.** iter-212's defect is not in the 264 at any scope.

## Close — 2026-08-09

**Outcome:** `claim_twin_guard`'s scope heading reads *"SCOPE IS TREE-WIDE FROM THE FIRST RUN,
DELIBERATELY"* and the tree it means is rosetta's **114** documents; `rosetta-extensions` — the repo
holding every fence and every rationale those fences act on — is in **no** claim fence's population
(**186** `.py` + **211** `.md`). Censused with the guard's own machinery: the tooling's **72 real-prose
documents state zero already-refuted claims**, and that zero **proves its own instrument in the same
run** (217 findings over the deliberately-RED fixture bucket). The Python half is **refused** for
widening with its cost re-derived every run.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (forty-sixth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — **counted, not felt: iters 212, 213, 214 = three tiks this run against a cap of
five** — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-214-1` … `D-M257x-214-3` (see [`decisions.md`](decisions.md))

**Audit:** `/usr/bin/python3 -m pytest` (**8.4.2** / CPython **3.9.6**), **Python**, `stack-core` only —
**40 passed / 0 failed** across `test_claim_twin_guard`, `test_claim_twin_guard_iter47_answer_key`,
`test_claim_twin_guard_iter48_answer_key` and `test_m257x_claim_twin_mutation_battery` (21 s).
**Stop condition did NOT fire:** 0 findings over the 72 real-prose documents.
**RED-proof battery, staged and removed:** copying ONE captured refuted-claim fixture into the tooling's
real prose took the census **RED at 73 documents / 1 finding**; the staged file was deleted and
`git status` verified back to the single intended edit.
*Scope, stated rather than implied (`§5` r60): `stack-core` only, Python only, changed-code reach. The
census READS all 11 rext sections' markdown and Python, but the arms and the runner are `stack-core`.
No whole-section run — the tree was edited during the iter. No Go, no TypeScript.*

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter212-a-retraction-does-not-reach-the-code-that-acts-on-it` — **still open, and this
  iter measured WHY it cannot be closed by the existing fence** (`D-M257x-214-3`): the claim was
  retracted in a route ledger, and `claim_ledger` derives only from blocker-ledgers.
- `SURVEY-M257x-iter214-route-retractions-are-not-in-the-claim-ledger` — **NEW.** The 264 claims come
  from audits' blocker-ledgers only. Route retractions — the ones iters write every close — are a
  **second, larger, unread** verdict source. Sizing it is the natural successor to this iter.
- `SURVEY-M257x-iter213-a-route-id-is-english` — unchanged.
- All routes from iters 207–210 and 212, unchanged, plus the standing queue.

**Lessons:**
- **A fence's scope heading can be true and still be read wrong.** *"Tree-wide, deliberately"* was
  accurate about the rosetta tree and silent about the repo the fence itself lives in.
- **A zero is worth what its instrument proves in the SAME run.** The fixture bucket makes this census's
  zero a measurement; without it, an identical arm would pass forever on a broken matcher.
- **Refusing to widen is a result, when the cost is measured and kept running.** 7 false REDs is the
  reason; an arm that fails when those 7 disappear is what stops the reason from rotting.
- **Censusing a neighbouring population does not close the route that pointed at it** — said explicitly,
  because the arm is sitting right where a future reader would assume otherwise.
