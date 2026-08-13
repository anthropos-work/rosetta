---
iteration_type: tik
status: in-progress
milestone: M258
iter: 06
created: 2026-08-12
---

# M258 iter-06 — wire the batch gate (and the restore leg it cannot ship without)

**Active strategy reference:** `TOK-01` — *measure the composition before engineering it* — **steps 2 + 3**.
Steps 1 (measure both halves) is discharged: bring-up **247.79 s** (iter-05, gateable, single-box),
batch **129 s** (iter-04, red set empty). The composition arithmetic is known (**≈ 376.8 s** vs a 480 s
ceiling), so the strategy's own precondition for treating this as a wiring job is satisfied.

## Step 0 — re-survey (mandatory, and it moved one item)

| target | TOK/route said | re-surveyed now | verdict |
|---|---|---|---|
| batch gate at `up-injected.sh:2810` | unwired | `grep run-playthroughs demo-stack/up-injected.sh` → **1 hit, a comment**. Genuinely unwired | **stands** |
| `RESTORE-M258-world-contract` | *"now owed in FACT — demo-1 is a pt-world stack"* (iter-04) | **REFUTED as a present-tense fact.** demo-1 holds **4 story orgs** (Cervato Systems, Meridian Talent, Northwind Aviation, Solvantis), **591 users**, and `cockpit-manifest.json` advertises all four hero trios | **re-framed** |
| red-set source | — | `report.State` ∈ {passing, failing, unimplemented, unimplementable-…}; `NoRegressions()` reads **`failing == 0`** only | **confirmed** |

The restore re-frame is the substantive one. iter-05's three bring-ups each re-seeded the presenter
world, so the demo world was restored **incidentally**. It is therefore owed as a **mechanism the wiring
must carry**, not as a repair owed today — and that is the stronger reason to build it in the same iter
as the gate, not a reason to drop it.

## Cluster / target identified

`TOK-01` step 2 names the gate; step 3 names the restore leg. **They are one deliverable and are planned
as a two-step shape**, because a batch wired *without* a restore leg is not a partial delivery of the
gate — it is a **regression against the gate's own text**. The gate requires the stack be left *"in a
presenter-usable world"*, and `overview.md` § *The world contract* shows the naive composition ending in
"a cockpit full of dead CTAs" that still satisfies *"the stack is left UP regardless"*. Shipping step 2
alone would make **every** bring-up do that. Declared here so the scope-creep tripwire counts against a
**planned** 2-step shape.

## Hypothesis

The batch can be driven from `up-injected.sh`'s tail under `D-v28-3` semantics, and the world restored
after it, entirely inside `rosetta-extensions` — 0 platform-repo edits. The composed measurement then
becomes takeable for the first time.

## Expected lift

Not a metric move: this iter makes the metric **measurable**. The primary metric (composed p50 over 3
cold cycles) cannot be sampled at all until the batch runs *inside* the bring-up. Success = one cold
cycle in which a single command brings the stack up, runs the batch to completion, emits one
consolidated red set, restores the presenter world, and exits with the right code — **proven live, and
proven in both directions** (green → 0; injected red → non-zero + loud).

## Phase plan

- **A — design + wire.** `batch-gate.sh` (D-v28-3 semantics + consolidated red set + verdict artifact),
  `restore-presenter-world.sh` (resolution (b)), and the hook beside the `autoverify.sh` invocation.
- **B — the fences.** New `DEMO_NO_BATCH` knob ⇒ a doc row in `demo-up-defaults.md` (the knob fence is
  bidirectional: a parser knob with no doc row is *undiscoverable*). Corpus delivery to
  `verification.md`.
- **C — controls.** Unit/behaviour tests, and a **live control proving the non-zero exit and the loud
  message** — the gate's red direction must be demonstrated, not assumed.
- **D — live proof.** One cold single-box cycle end-to-end with the batch wired.
- **E — close** + tag + `git push --tags` (a tag that exists only locally is unreachable to every stack).

## Escalation conditions

- A non-empty red set on the live proof that is **not** injected by the control → that is the D-v28-3
  escalation itself; report it, do not fix it silently.
- The restore leg failing to reconstitute the presenter world → user-blocker (it defeats a gate clause).
- Any need to touch a platform repo → hard stop (hard constraint).

## Acceptable close-no-lift outcomes

If the wiring proves to need a design decision the milestone has not taken (e.g. the batch cannot be made
default-on without breaking the presenter path), closing with that falsification **documented and
measured** satisfies the protocol — the composition arithmetic already exists and is not at risk.
