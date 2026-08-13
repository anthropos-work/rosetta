---
iter: 152
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-08
---

# iter-152 — the probe registry has never been fenced against the platform

**Active strategy reference:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop sampling
them.* This iter takes the class **"a registry row's port literal versus the platform's own published
port"**, enumerates it exhaustively in **both** directions, and ships the fence.

## Step 0 — re-survey (mandatory)

`TOK-08` does not name a next target by row; run 3 flagged
`SURVEY-M257x-iter148-registry-is-hand-maintained` as the strongest open route and the prompt attached a
condition to it: **grade the cost first, because iter-145's `REGISTRY_BASES` is a deliberate anti-vacuity
control that a naive derivation would delete.** The re-survey confirms the route is still open and
sharpens what it actually is:

- `stack-verify/lib/services.sh` is a 13-row hand-maintained table, one **base host port literal** per row.
- iter-145 already fenced it against its **test-side twin** (`REGISTRY_BASES` +
  `test_the_test_side_registry_mirrors_services_sh`) — membership **and**, through the offset sweep,
  ports. That half is closed.
- **What is NOT fenced is the half that matters to this milestone.** Both copies are hand-maintained
  against a third thing — the platform's own compose — and the table says so **in its own header**:
  *"Source of truth: the platform's docker-compose.yml service set."* Nothing checks it. This is the
  milestone's founding class verbatim: a service leaves the platform, or moves a port, and **nothing on
  our side notices.**

So the target is **not** "derive the registry from compose" (that would delete the independent copy and
make iter-145's sweep assert nothing). The target is the **third edge**: assert the two independently
maintained artifacts agree, and fail by name when they do not.

## Cluster / target identified

The registry↔platform-compose edge, censused in both directions at platform **`0c91421`** (`stack-demo/platform`,
verified `0` behind `origin/main` at iter open — the same ref gate clauses 1/2 were proven at).

## Hypothesis

The forward direction (*does every row's literal match compose*) is clean — the ports have not drifted.
The **reverse** direction is where the finding is, because `§5` rules 66/69 say it plainly: **a token
census finds a WRONG value and can never find an ABSENT one.** The registry's denominator is
**services**; compose's denominator is **published ports**. If those denominators differ, the registry
has a blind class it cannot express, and no reading has ever looked.

## Expected lift

No `N` reading is planned, so **no `N` movement will be claimed** (`§9` guard-rail 1). The deliverable is
a fence that enumerates the class corpus-wide and holds it at zero, per `TOK-08`'s definition of working
a mechanical class.

## Phase plan

- **A** — census both directions at `0c91421`; state both denominators.
- **B** — make retirement **declared** rather than inferred from a prose comment (an absent-value fence
  whose "absent" arm reads a comment is vacuous — `D-M257x-151-1`'s partition lesson).
- **C** — ship `service_registry_guard.py`; run it; run a mutation control and an anti-vacuity control
  that actually fire.
- **D** — regression test; run the `stack-core` guard tests that cover it and `stack-verify`'s registry
  fence. **`stack-core` in full is NOT run unless this iter changes something it covers** (`§5` rule 60).

## Escalation conditions

If the census finds a **live** row whose literal disagrees with compose, that is a live defect in the one
section that grades every stack — land the repair in this iter, do not route it.

## Acceptable close-no-lift outcomes

A clean negative in both directions is a first-class result (iter-149 is the precedent) **provided** the
fence ships with controls that fire — otherwise the iter proved nothing.
