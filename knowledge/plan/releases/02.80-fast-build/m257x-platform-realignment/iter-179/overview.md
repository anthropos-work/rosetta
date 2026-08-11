---
iter: 179
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-179 — the two fence registries are an EQUALITY, and one direction of it was fenced

> ⚠ **OUTCOME: this title's premise was REFUTED by the iter it opened.** The equality is a *conclusion*
> of two conjuncts, not a two-way contract with one side missing — the direction this overview expected
> to find unfenced is **entailed** by the pair, and what was actually broken was the check's **reach**,
> not its direction and not its cost. **The text below is left exactly as pre-registered**, because a
> hypothesis edited after its measurement is not a hypothesis. See [`progress.md`](progress.md) Phases
> B–C and `D-M257x-179-2` / `-3`.

**Type:** tik. **Active strategy: `TOK-08`** — *census the mechanical classes; stop sampling them.*
This iter works the class `TOK-08` names as sweepable: a **mechanically decidable** cross-registry
contract, enumerated by a check that keeps running, rather than sampled by a reading.

## Step 0 — re-survey before targeting (mandatory)

`TOK-08`'s standing direction is *work the mechanical classes in descending measured size*; it does not
name iter-179's target, so the target comes from the open-route queue. Re-surveyed at HEAD `8422706`,
trees clean (rosetta modulo the user's `.claude/settings.json`; rext on `main` at `f0ac4ed`):

| route | state at re-survey |
|---|---|
| `FIX-M257x-iter174-accept-registers-one-registry-of-two` | **class closed at iter-176** (the registry *population* is fenced, 7 sites). **Its named member is still open** and is now carried as the verdict text of one site in that very fence. |
| `FIX-M257x-iter177-derivation-registry-decline-rationale-is-false` | open; owner unassigned |
| `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites` | open by design (`F4` — widening silently is the conflation) |

**Target: the open member.** Not a substitution — iter-176 narrowed the route explicitly to *"the battery
seed list a `--accept` does not write"* and left it in the fence as an open obligation. Closing an
obligation the fence itself publishes is the cheapest legible next move under `TOK-08`.

## Cluster / target identified

Two registries have to agree or the mechanical-fences mutation battery cannot run:

1. `stack-core/repair_postcondition_baseline.json` — `fences` (written by `--accept`, a ratchet)
2. `stack-core/tests/test_m257x_mechanical_fences_mutation_battery.py` — `_COPY_FILES` (hand-maintained)

`test_000_the_copy_list_stages_every_fence_the_baseline_names` asserts **one** direction:
*baseline-named ⊆ staged*. That is the direction whose failure is *"the fence is broken"* dressed over
*"you forgot a file"* — the one iter-121 paid for.

## Hypothesis (stated so it can be wrong)

The routed repair — *"make `--accept` write the second registry too"* — is a **hypothesis, and this iter
expects to refute it**, per the standing rule that a routed item's proposed repair is not a plan.
`test_000` asserts *staged ⊇ the baseline's fence names*, so a `--accept` that wrote both from one source
would make that assertion compare a set with itself (iter-158's shape, already recorded in the battery's
own comment). What is expected instead:

- the contract is not *inclusion* but **equality** — the staged tree's `discover_fences()` derives
  participation from the `FENCE_KIND` declaration **on the staged disk**, and the staged run compares that
  against the staged copy of the baseline. A staged participating fence the baseline does **not** name
  fails the run in the *other* direction, with a different symptom (*unrecorded sites*), and **nothing
  asserts it**;
- the reporting lag the route complains about is **not a cost problem**. Predicted: the contract is
  sub-second, and its practical latency is set by *which runs reach the file it lives in*, not by how long
  it takes.

## Expected lift

No `P`/`N` reading is taken, so **no clause-5 movement is claimed** (`§9` — UNMEASURED is not unmoved).
The deliverable is a fence: the full equality, both directions, each direction's failure naming the
distinct symptom it prevents, reachable from a scoped run, with a mutation control per direction and an
anti-vacuity control that can fire.

## Phase plan (planned multi-step shape — declared, per the scope-creep carve-out)

1. **Measure** both directions on the real tree at HEAD, on **both** interpreters, naming units and runner.
2. **Measure the latency claim** — the check standalone, and the file it currently lives in.
3. **Land** the equality as one derivation with two call sites (no second derivation — iter-175's rule).
4. **Controls**: anti-vacuity + one mutation control per direction.
5. **Reconcile the registry classification entry** that carries this route's text, so the population fence
   does not keep publishing a closed obligation as open.
6. **Protocol doc** §8 rule.

## Escalation conditions

- If landing the fast check turns the mechanical-fences battery or the registry-population fence RED for a
  reason that is **not** this iter's change → route it, do not weaken the assertion (and re-derive the
  cause before believing the pre-registration's own guess — `D-M257x-174-3`).
- If the equality is **not** the true contract (i.e. the second direction is legitimately allowed) → say
  so with the measurement and land only the direction that is real. A wrong assertion is worse than none.

## Acceptable close-no-lift outcomes

Measuring that both directions already hold, that the second is a real requirement, and that the existing
fence covers one of two — with the instrument proven able to fire on both — is a complete iter even if the
repair turns out to be one assertion and one comment.
