---
iter: 17
milestone: M258
iteration_type: tik
status: in-progress
created: 2026-08-12
---

# iter-17 — prove it live: the policy fix, on a demo stack and on dev

**Active strategy reference:** `TOK-01` (measure the composition before engineering it) — this iter is
the *measure* half of iter-16's fix. `TOK-02` does not control it.

## Step 0 — re-survey (mandatory)

`PROVE-M258-iter17-policy-fix-on-demo-and-dev` re-verified **OPEN**: iter-16 proved the *mechanism*
(publish → `policy changed elsewhere, reloading` on `demo-3`) but **nobody has re-run the batch and
watched the 15 reds clear**. A mechanism proof is not an end-to-end proof, and this milestone's whole
premise is that "it should work" is not a result.

Host at iter open, `macmini`, 11:32Z: `load1` **3.41** (it was 26–33 during the batch that produced
the red set), **182 GiB** free, **one** stack resident (`demo-3`, 10 containers). These are the best
conditions of the milestone — the orchestrator's ruling notes it, and it is why this iter is
affordable now and was not before.

## Cluster / target identified

Two rulings converge on one iter:

1. **The 15 reds must be RESOLVED, not carried.** iter-16 attributed and fixed; what remains is the
   proof that the fix clears them.
2. **Every piece of new functionality this release built must be proven on a demo stack AND a dev
   stack.** Named: the batch gate, the world-contract restore leg, the studio-desk multi-stage image,
   the `down -v` volume fix, L1's multi-stage next-web/hiring, and the space-reclaim classes.

## Hypothesis

Re-running the full bring-up + batch on a **fresh slot**, with `rosetta-extensions` pinned to
`fast-build-m258-iter-16`, clears the org-scoped red cluster. **Falsifiable and predicted per-item,
before the run** — this is the sharp edge of the iter, so the prediction is written down first:

- The **15 org-scoped reds go green** (11 refused reads + 4 refused writes).
- The **2 negative controls go green** — they assert a manager sees her *own* tenant and not another's;
  with everything refused, the own-tenant half could not pass.
- The **15 already-passing user-scoped Playthroughs stay green** (nothing in the fix touches them).
- `pt-hiring-recruiter-compare` is the **one I do not predict green**: it asserts *exactly 5* shared
  positions, and `FIX-M258-iter15-hiring-under-set-dressed` is independently open and live
  (`autoverify.json` `green:false, warnings:1`, 38 sessions vs ≥40). If it stays red **for that
  reason**, that is a separate defect surfacing, not this fix failing — and the distinction must be
  read from its assertion text, not assumed.

**If the org-scoped cluster does NOT clear, the iter-16 attribution is wrong** and must be reopened
rather than explained. That is the falsifier.

## The dev half — three questions per piece, not one verdict for the set

Much of this release landed under `demo-stack/`. Some of it may be **demo-only by construction**. For
each named piece, answer all three:

| | question |
|---|---|
| a | does it **apply** to the dev path at all? |
| b | is it **wired** there? |
| c | if not — **is that correct**, or a gap? |

A piece that does not exist on the dev path is a **finding to report**, not a proof failure. The
answers are per-piece; a single verdict for the set would hide exactly the interesting case.

## Phase plan

- **A** — re-pin the `stack-demo` consumption clone to `fast-build-m258-iter-16`. ⚠️ This clone owns
  `demo-3` (`D23`); `demo-stack/.gitignore:8` ignores `stacks/`, so the live stack's state is not
  disturbed. iter-15 hit the M217 FATAL pin guard on a **half-completed** re-pin — complete it, verify.
- **B** — heartbeat, then bring up a **fresh slot** (`demo-4`) cold. The bring-up drives the batch
  gate itself, so one command produces build + set-dress + batch + restore.
- **C** — read the verdict against the per-item prediction above. Attribute anything left red.
- **D** — the dev half: the three questions per piece, answered from the dev path's own code.
- **E** — converge back to **exactly one stack** (`END-M258-one-stack`), tearing down `demo-4` and
  never `demo-3`. Heartbeat before the teardown. Record, route, close.

## Escalation conditions

- The org-scoped cluster does **not** clear → the iter-16 attribution is refuted; reopen it in this
  iter's record, do not paper over it, and escalate.
- The bring-up needs `demo-3` torn down to fit → **refuse**, record as a HEADROOM-class result, and
  report what a second slot would have cost.
- Any dev-side proof would require editing a platform repo → refuse; that line is absolute.

## Acceptable close-no-lift outcomes

- The bring-up fails for a reason unrelated to the fix (build break, cold-cache stall) → record the
  failure with its environment and `load1`, and state plainly that the fix is **still unproven
  end-to-end**. An honest "not proven" beats a green read from a run that did not test it.
