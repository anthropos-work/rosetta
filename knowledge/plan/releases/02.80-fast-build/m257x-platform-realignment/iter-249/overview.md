---
iter: 249
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
---

# iter-249 — the owed whole-suite reading, taken against a tree that cannot move

**Type:** tik · **Active strategy:** [`TOK-08`](../decisions.md) — census the mechanical classes; stop
sampling them.

## Step 0 — Re-survey before targeting

`ROUTE-M257x-248-whole-suite-was-never-measured-cleanly-this-run` is the route iter-248 named as **the
first thing the next session should do, before any new work**. Re-surveyed at open:

- rosetta `971cdc4`, tree clean modulo the user-owned `.claude/settings.json` (excluded by standing order).
- `.agentspace/rosetta-extensions` `dfb3fb6`, **0 dirty files**.

The route is live and its precondition — a tree nobody is editing — holds **right now** and will stop
holding the moment this iter does any work. That is exactly the trap iter-248 fell into: it started a
27-minute suite three minutes before its own commits and had to refuse to attribute the result.

## Cluster / target identified

`TOK-08` names the controlling strategy; the target is `ROUTE-M257x-248`. But the route as written is
**not executable as stated** — "run the suite before any new work" serialises a 27-minute measurement in
front of every session, and iter-248 already demonstrated that the discipline fails under budget
pressure. The measured cause is not "someone ran it at the wrong time"; it is that **the instrument reads
a tree that the operator owns**. `stack-core`'s guards resolve their subject via `ROSETTA_ROOT`, falling
back to `Path(__file__).resolve().parents[3]` — the **live** corpus
(`claim_twin_guard.py:335`, `anchor_construct_guard.py:1156`). So the suite has no way to describe
anything but the tree as it is at each moment of a 27-minute walk.

**The repair is to give the reading a subject that cannot move**: take the measurement against a
content-addressed *frozen export* of both repos at pinned shas, with `ROSETTA_ROOT` pointed at the export.
Then the run describes exactly one state, is reproducible from two shas, and no longer forbids working
while it runs.

## Hypothesis

Against a frozen export at rosetta `971cdc4` + rext `dfb3fb6`, the whole `stack-core` section returns a
**materially smaller** failure count than iter-248's `16 failed`, because most of that 16 was the moving
tree. Whatever survives the freeze is a **real** failure and is attributable to a commit.

## Pre-registered numeric claims — sealed in this iter's FIRST commit

Stated before the frozen export is built or run. Each is falsifiable and will be graded verbatim at close.

| # | claim | prediction |
|---|---|---|
| **PR-1** | whole-section `stack-core` failures against the frozen export | **≤ 6** (point estimate **2**) |
| **PR-2** | at least one iter-248 failure is **NOT** reproducible on the frozen export (i.e. was a race) | **true** |
| **PR-3** | collected test count on the frozen export **exceeds** iter-248's `2,064 + 16 = 2,080` | **true** — iter-248 repaired 3 files whose classes sat after the `__main__` guard |
| **PR-4** | every surviving failure is attributable, by `git log` on its own test file, to a commit **inside this milestone** (not inherited from before M257x) | **false** — I expect ≥1 inherited |
| **PR-5** | the frozen export changes the *verdict*, not just the noise: at least one test that PASSED in iter-248's moving run **FAILS** frozen | **false** |

**Honest direction of the guess.** Three consecutive pessimistic pre-registrations about this tooling were
wrong in the same direction — the tooling was better than predicted. PR-1's point estimate is set
accordingly (2, not 10), and PR-5 predicts the boring outcome rather than the interesting one.

## Phase plan

1. **Seal** this overview as `probe(M257x/249)` — before the export exists.
2. **Build the frozen export**: `git archive HEAD` from both repos into the scratchpad, laid out so that
   `<export>/.agentspace/rosetta-extensions` mirrors the live layout, and run with
   `ROSETTA_ROOT=<export>`.
3. **Run** the whole `stack-core` section against it, in the background, capturing the full log.
4. **Attribute** every failure: reproduce, read it, `git log` its test file and its subject.
5. **Repair** what this milestone caused; route what it did not.
6. **Grade** PR-1…PR-5 verbatim and close.

## Escalation conditions

- If the frozen export cannot reproduce the suite at all (import errors from the archive layout), the
  export design is refuted — record it, fall back to an in-place frozen window, and say so.
- If a surviving failure is a real defect in shipped tooling with no cheap repair, route it forward with a
  named handler rather than growing this iter.

## Acceptable close-no-lift outcomes

- The frozen export runs and returns **0 failures**: PR-1 holds, PR-2 holds, and the deliverable is the
  reading plus the reusable frozen-export instrument. That is a complete iter even with nothing to repair.
- A surviving failure turns out to be inherited from before M257x: PR-4 is graded true, the failure is
  routed, and the iter closes on the reading.
