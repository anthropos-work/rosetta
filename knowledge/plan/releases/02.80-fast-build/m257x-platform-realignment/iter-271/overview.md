---
iter: 271
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
route: ROUTE-M257x-270-prove-the-spent-pin-cold
---

# iter-271 — prove the spent pin COLD: three cycles at tooling nobody has run

**Type:** tik, under `TOK-08` (*census the mechanical classes; stop sampling them*).

## Step 0 — re-survey (mandatory, before targeting)

`TOK-08`'s next-tik direction is the mechanical-class sweep for gate clause 5. **It is stale for this
iter, and the substitution is forced by our own last iter**, not by a change of strategy.

Re-surveyed at open (2026-08-10T19:59Z):

| fact | measured |
|---|---|
| `.agentspace/rext.tag` | `fast-build-m257x-iter-270` |
| that tag on **origin** (rung zero) | **present** — `refs/tags/fast-build-m257x-iter-270` → `4e5fb25`, peeled `2833a64` |
| `stack-demo/rosetta-extensions` checked out at | **`fast-build-m257x-iter-101`** — the OLD pin |
| clause-1 proof of record (iter-260, 3 consecutive green cycles) | ran at **`fast-build-m257x-iter-101`** |
| free disk / docker VM / host load | 193 GiB · 11.67 GiB, overlayfs, 8 CPU · load 2.66 |

**Clause 1's proof is stale by our own hand.** iter-270 spent the frozen-pin control deliberately and
bumped the pin **206 commits**. The three cycles that establish clause 1 exercised tooling that a cold
bring-up no longer consumes, and this milestone's own rule (r60/66) says a scoped green is evidence about
its scope alone. iter-270 named the consequence itself and routed it as the highest-value item:
`ROUTE-M257x-270-prove-the-spent-pin-cold`.

The substitution is therefore: **TOK-08 named the clause-5 mechanical sweep; iter-270's own close
invalidated the clause-1 evidence, so clause 1 is re-proven first.** The strategy chain is untouched —
only the named next-target is superseded, by evidence generated after `TOK-08` was written.

## Cluster / target identified

**The bring-up paths iter-270 changed are exactly the paths a cold cycle exercises and a warm stack does
not.** iter-270 rewrote service injection (`derive_inject_svcs` now fails CLOSED; `INJECT_CANDIDATES`,
`INJECTED` and `REUSE_DEV` pruned of decommissioned names) and the studio acquisition (demo's was anchored
on the decommissioned `cms` clone; dev's did not exist). None of that has been run cold.

A warm `demo-2` has been up 5 hours on the old tooling. It proves nothing about the new.

## Hypothesis

The repairs hold under a cold cycle, and clause 1 is re-established at the pin that will ship. The
alternative is the more valuable finding: **the class of defect iter-270 repaired is the class that only a
cold cycle can surface**, so if the repairs are wrong, this is where it shows.

## Expected lift

- Gate **clause 1 re-established at the shipping pin** — 3 consecutive `down --purge` + `up` cycles green
  (`autoverify green:true / 0 warnings`, `EXIT_CODE=0`), byte-identical invocations, no retries, no
  per-cycle intervention. Anything less is reported as what it is.
- Every pre-registration below graded against measurement, refutations carried forward.

## Phase plan (declared multi-step — the tripwire counts UNPLANNED lines only)

1. Seal these pre-registrations (first commit).
2. Re-point `stack-demo/rosetta-extensions` to the SoT pin, so **both halves** of every cycle run the new
   tooling (the M217 FATAL guard requires the checkout; the operator does it, not the script).
3. Cycle 1: `rosetta-demo down 2 --purge` + bare `up-injected.sh 2`. Grade on `autoverify`.
4. Cycles 2 and 3: byte-identical. Clause 1 wants **three consecutive**.
5. Grade the pre-registrations; record what the cold run proved about iter-270's repairs.

## Out of this iter's planned scope (declared, so the tripwire is clean)

- `FIX-M257x-269-force-append-grows-the-demo-env-without-bound` — rides a later tag (iter-270's ruling).
- `ROUTE-M257x-270-directus-consumer-cms-key` — no tag, no cold cycle needed; a later iter.
- Gate clause 2 (`FIX-M257x-267-capture-the-succession-RESPONSE`) and clause 5 (the reading) — both want a
  **green stack to run against**, which is what this iter produces. They are the natural next iters, not
  this one.

## Escalation conditions

- **`demo-1` is not ours.** No `--purge 1`, no stop or restart of any `demo-1-*` container. If a cycle
  touches `demo-1`, that is a defect in the tooling and it escalates.
- **The host is permanently contended.** A boolean survives contention; a timing does not. Every duration
  is labelled CONTENDED and none is published as a baseline.
- If a cycle goes red, the iter does **not** retry to manufacture a green. It reports how far it got and
  the failure is the deliverable.
- **No force-push of any kind.** No new tag is required by this iter.

## Acceptable close-no-lift outcomes

A red cycle with the failure characterised — which path, which change of iter-270's, and the evidence — is
a first-class result. It is the finding the route was written to obtain, and it is worth more than a green.

## Pre-registrations (sealed in this iter's FIRST commit, before any measurement)

Stated falsifiably, before looking.

- **PR-1 — the cold cycle is green at the new pin, first attempt.** Cycle 1 reaches `autoverify
  green:true` with **0 warnings** and `EXIT_CODE=0`, with no retry and no intervention.
  *Refuted by:* any non-green first attempt, or any hand-intervention needed to reach green.
- **PR-2 — the new studio acquisition runs and no longer sources from the decommissioned clone.** After a
  cold cycle, `stack-demo/app/studio/requirements.txt` is present, and `stack-demo/cms/studio/` is **not**
  re-populated by this cycle (its mtime does not advance into the cycle window).
  *Refuted by:* `app/studio` unpopulated, or `cms/studio` written during the cycle.
- **PR-3 — the fail-CLOSED arm never fires, because the derivation succeeds.** The bring-up derives its
  inject set from the platform compose without hitting iter-270's `die`, and the derived set contains
  **no** `cms` and **no** `jobsimulation`.
  *Refuted by:* a die on the derivation, or either name in the injected set.
- **PR-4 — no decommissioned service name reaches the generated compose override.** The override the
  cycle generates yields **0** compose service keys in
  {`cms`, `jobsimulation`, `roadrunner`, `storage`, `skillpath`, `messenger`, `customerio-sync`}.
  *Refuted by:* any one of them present as a service key.
- **PR-5 — the three cycles are not independent draws.** If cycle 1 is green, cycles 2 and 3 are green
  too; the failure mode this tooling has shown historically is deterministic (a missing path, a wrong
  hardcode), not flaky. *Refuted by:* a mixed result — any green/red split across the three.
