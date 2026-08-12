---
iter: 5
milestone: M257
iteration_type: tok
tok_flavor: triggered
status: in-progress
opened: 2026-08-11
---

# iter-05 — the gate re-pointed at a host that exists (triggered tok)

**Type:** tok (triggered — the 3-no-prog-tik streak) · **Authors:** `TOK-02`

## Why a triggered tok fires here

Phase 0 rule 2 walks back over the last **3 tiks** — iter-02, iter-03, iter-04 — and every one of them
closed with the same line: **"metric delta 0, by design."** That is the streak, and it is real rather than
an artefact of measurement noise: the protocol's primary metric is the **cold `demo-down --purge` +
`demo-up` p50 on the gate's named host**, and that number **has never been measured, on any host this
project may still use**. Three tiks in a row could not move a metric that does not exist yet.

**Step 0 re-survey (mandatory, and it confirms rather than falsifies the trigger).** The skill requires
re-running the primary measurement before authoring a revised strategy, because a stale trigger must not
produce a strategy revision. Re-surveyed:

- `hostprofiles/` holds `billion.json`, `laptop.json`, `macmini.json`. Only `billion.json` carries a
  `gated_baseline`. **No p50 exists for any host this milestone may measure on** → the delta across the
  three tiks is genuinely 0, not "moved but unrecorded".
- The gate's named host, `odysseus`, is **retired** (`D-v28-15`). There is no profile for it and there
  never will be.

The trigger stands. What the three tiks actually delivered was **everything that had to be true before a
number could be taken** — the instrument proven falsifiable (iter-02), the two blockers that made READY
unsatisfiable on *every* host (iter-03), and the host that exists, measured (iter-04). That is not a
stall in the work; it is a stall in the **metric**, and a metric that cannot move for three iters is
exactly what the tok mechanism is for.

## What stopped working: `TOK-01`'s step 3 named a machine that no longer exists

`TOK-01` sequenced: **(1)** the host can run a cycle → **(2)** the instrument is proven able to fail →
**(3)** that host's own `n ≥ 3` p50 baseline is checked in → *then* levers, largest-measured-second first.

Steps 1 and 2 landed. **Step 3 could not, and not for want of trying**: every one of its three nouns —
the host, its profile filename, the baseline — named `odysseus`, and `D-v28-15` retired it one day after
`TOK-01` was written. The strategy was correct in shape and dead in its references.

## Scope of this iter (a tok, so the deliverable is doctrine + the repair that makes it usable)

1. **Re-cut the `exit_gate`'s DEAD HOST REFERENCE** — `odysseus` → this Mac mini (`macmini.json`).
   **A stale-reference repair, not a re-scope, and explicitly not a relaxation** (see § below).
2. **Land `DOC-M257-hostclass-retraction`** — `state.md`, `roadmap.md` `D-v28-15` and this milestone's own
   § HOST CLASS PROBLEM all assert *"the Mac pays no unpack leg."* **Measured false on this machine.**
3. **Author `TOK-02`** in the milestone-root `decisions.md`.

Deliberately **not** in this iter (routed to iter-06, named handlers): `FIX-M257-load1-units-vm` and the
`n ≥ 3` contended baseline campaign. Both are tik work — code and measurement — and a tok that took them
would blur the distinction the close-status grading depends on.

## Phase 0b — SKIPPED, and why

Phase 0b re-runs on a triggered tok only when the revised strategy **redirects into a subsystem the
milestone's standing audit did not cover**. `TOK-02` redirects the *host*, not the subsystem: the same
`buildbench` harness, the same `demo-stack` bring-up, the same `autoverify` verdict, all covered by the
iter-01 audit whose YELLOW verdict is recorded in `spec-notes.md`. No new blind area is entered, so the
standing verdict is inherited.

## Phase 0d — SKIPPED, and why

No wire-through pipeline is authored here. The iter's changes are plan-state edits (frontmatter, prose)
plus a doctrine entry; the gate tool this milestone consumes (`buildbench`) is not being extended in this
iter — it is extended in iter-06, which is where its pre-flight belongs.

## Escalation conditions

- If re-cutting the host reference cannot be done **without** moving a target or dropping a clause → stop
  and escalate, because that would be planning rather than repair.
- If the retraction turns out to rest on `docker info` alone → stop; a config string is not a hardware
  measurement, and that is the precise mistake being retracted.
