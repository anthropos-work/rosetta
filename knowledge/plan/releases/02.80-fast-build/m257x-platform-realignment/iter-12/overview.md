---
milestone: M257x
iter: 12
iteration_type: tik
status: in-progress
opened: 2026-08-01
---

# iter-12 — gate clause 1: three consecutive cold cycles, 0 warnings

**Active strategy reference:** `TOK-01: instrument first, then follow` — step 5/5, *"prove it cold, three
times."* Steps 1–4 are done: the pin is clean, the migration set is derived, the fences are landed and
watched RED, and as of iter-11 the verifier no longer reports absence it cannot substantiate. What is left
under this strategy is the proof.

## Step 0 — Re-survey

- `demo-1` is UP (16 containers) and its verdict file, written by a run given **no path at all**, reads
  `{"warnings":0,"green":true}` (iter-11, 2026-08-01T06:28Z). That is the standing stack, **not** a cold
  cycle — clause 1 requires `demo-down --purge` then `demo-up`, three times.
- rext consumption clone re-pinned to `fast-build-m257x-iter-11`, verified on origin.
- Every autoverify defect iters 04–11 found was found on a bring-up, so a cold cycle is the measurement that
  has actually paid in this milestone.

## Cluster / target identified

Clause 1: *a cold `demo-down --purge` + `demo-up` reaches `autoverify green:true / 0 warnings` across three
consecutive cycles.* Cycles are ~18 min each here.

Opened first, because it fires on every one of those bring-ups:

- **`FIX-M257x-vmram-gib-unit`** — `up-injected.sh:258-262` floors `docker info` MemTotal bytes to integer
  GiB, so a VM documented as "12 GB" (decimal; measured 12528664576 B = 11.67 GiB) reads as **11** and trips
  the `< 12 GiB` note on every bring-up. A doc/code **unit mismatch** — this milestone's own subject matter,
  never re-measured. Non-fatal, so it does not itself block clause 1; it is noise in the exact log the gate
  is read from.

## Hypothesis

With the pin current and iter-11's derivation in place, a cold cycle reaches `green:true / 0 warnings`. Each
cycle is also a fresh negative control for every fence iters 02–11 landed — the seeders, the derived
migration set, the write-target fence, and the receipts asserts all run for real.

## Expected lift

Clause 1 goes 0 → 3 cycles, or the first cycle names a specific defect and the iter becomes that defect's
root-cause work (which is how iters 04, 05, 06 and 10 each actually went).

## Phase plan

1. `FIX-M257x-vmram-gib-unit` + its regression test + the `demo-up-defaults.md` citation repair that editing
   `up-injected.sh` forces (iter-05 hit this; the corpus guard's `--fix` is the handler).
2. Cycle 1: `demo-down 1 --purge` → `demo-up 1`. Read `autoverify.json` + the full transcript.
3. Cycles 2 and 3, same. Heartbeat throughout — they are long and mostly waiting.
4. Close on the measured cycle count, whatever it is.

## Escalation conditions

- A cycle failing for a reason that needs a platform-source change → `demopatch` per `demopatch-spec.md`,
  never a platform edit.
- Two cycles failing for **different** causes → stop stacking cycles and root-cause instead; three green
  cycles is the claim, and grinding cycles against a moving fault is not evidence.

## Acceptable close-no-lift outcomes

A cycle that fails with a named, measured mechanism is the iter's deliverable even with 0 green cycles —
that is exactly how the four highest-value iters of this milestone closed.
